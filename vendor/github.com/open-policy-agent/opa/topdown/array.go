// Copyright 2018 The OPA Authors.  All rights reserved.
// Use of this source code is governed by an Apache2
// license that can be found in the LICENSE file.

package topdown

import (
	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/topdown/builtins"
)

func builtinArrayConcat(a, b ast.Value) (ast.Value, error) {
	arrA, err := builtins.ArrayOperand(a, 1)
	if err != nil {
		return nil, err
	}

	arrB, err := builtins.ArrayOperand(b, 2)
	if err != nil {
		return nil, err
	}

	arrC := make(ast.Array, 0, len(arrA)+len(arrB))

	for _, elemA := range arrA {
		arrC = append(arrC, elemA)
	}

	for _, elemB := range arrB {
		arrC = append(arrC, elemB)
	}

	return arrC, nil
}

func builtinArraySlice(a, i, j ast.Value) (ast.Value, error) {
	arr, err := builtins.ArrayOperand(a, 1)
	if err != nil {
		return nil, err
	}

	startIndex, err := builtins.IntOperand(i, 2)
	if err != nil {
		return nil, err
	}

	stopIndex, err := builtins.IntOperand(j, 3)
	if err != nil {
		return nil, err
	}

	// Clamp stopIndex to avoid out-of-range errors. If negative, clamp to zero.
	// Otherwise, clamp to length of array.
	if stopIndex < 0 {
		stopIndex = 0
	} else if stopIndex > len(arr) {
		stopIndex = len(arr)
	}

	// Clamp startIndex to avoid out-of-range errors. If negative, clamp to zero.
	// Otherwise, clamp to stopIndex to avoid to avoid cases like arr[1:0].
	if startIndex < 0 {
		startIndex = 0
	} else if startIndex > stopIndex {
		startIndex = stopIndex
	}

	return arr[startIndex:stopIndex], nil
}

func init() {
	RegisterFunctionalBuiltin2(ast.ArrayConcat.Name, builtinArrayConcat)
	RegisterFunctionalBuiltin3(ast.ArraySlice.Name, builtinArraySlice)
}