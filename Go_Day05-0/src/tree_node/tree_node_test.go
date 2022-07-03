package main

import (
	"reflect"
	"testing"
)

var trees = []TreeNode{
	{true, nil, nil},
	{false, nil, nil},
	{false, &TreeNode{true, nil, nil}, &TreeNode{true, nil, nil}},
	{false, &TreeNode{false, nil, nil}, &TreeNode{false, nil, nil}},
	{false, &TreeNode{true, nil, nil}, &TreeNode{false, nil, nil}},
	{true, &TreeNode{false, nil, nil}, &TreeNode{true, nil, nil}},
	{
		true,
		&TreeNode{
			false,
			&TreeNode{true, nil, nil},
			&TreeNode{false, nil, nil},
		},
		&TreeNode{
			true,
			&TreeNode{false, nil, nil},
			nil,
		},
	},
}

var toysBalancedExpected = []bool{
	true,
	true,
	true,
	true,
	false,
	false,
	true,
}

var unrollGarlandExpected = [][]bool{
	{true},
	{false},
	{false, true, true},
	{false, false, false},
	{false, true, false},
	{true, false, true},
	{true, false, true, false, false, true},
}

func TestAreToysBalanced(t *testing.T) {
	for i, tree := range trees {
		result := tree.areToysBalanced()
		expected := toysBalancedExpected[i]
		if result != expected {
			t.Errorf("#%d Expected:\n%t\n Got:\n%t", i+1, expected, result)
		}
	}
}

func TestUnrollGarland(t *testing.T) {
	for i, tree := range trees {
		result := tree.unrollGarland()
		expected := unrollGarlandExpected[i]
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("#%d Expected:\n%t\n Got:\n%t", i+1, expected, result)
		}
	}
}

func BenchmarkTestAreToysBalanced(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tree := range trees {
			tree.areToysBalanced()
		}
	}
}

func BenchmarkTestUnrollGarland(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tree := range trees {
			tree.unrollGarland()
		}
	}
}
