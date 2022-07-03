package main

import (
	"reflect"
	"testing"
)

var nCoolestTests = []struct {
	presents []Present
	n        int
	expected []Present
}{
	{[]Present{{3, 1}}, 0, []Present{}},
	{[]Present{{3, 1}}, 2, nil},
	{[]Present{{3, 1}}, -1, nil},
	{
		[]Present{{5, 1}, {4, 5}, {3, 1}, {5, 2}},
		2,
		[]Present{{5, 1}, {5, 2}},
	},
	{
		[]Present{{3, 1}, {5, 2}, {5, 1}, {4, 5}},
		4,
		[]Present{{5, 1}, {5, 2}, {4, 5}, {3, 1}},
	},
}

func TestGetNCoolest(t *testing.T) {
	for i, test := range nCoolestTests {
		result := getNCoolestPresents(test.presents, test.n)
		expected := test.expected
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("#%d Expected:\n%v\n Got:\n%v", i+1, expected, result)
		}
	}
}

func BenchmarkTestGetNCoolest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range nCoolestTests {
			getNCoolestPresents(test.presents, test.n)
		}
	}
}

var grabPresentsTests = []struct {
	presents []Present
	cap      int
	expected []Present
}{
	{
		[]Present{{5, 1}, {4, 5}, {3, 1}, {5, 2}},
		3,
		[]Present{{5, 1}, {5, 2}},
	},
	{
		[]Present{{3, 5}, {5, 10}, {4, 6}, {2, 5}},
		14,
		[]Present{{3, 5}, {4, 6}},
	},
	{
		[]Present{{5, 4}, {4, 3}, {3, 2}, {2, 1}},
		6,
		[]Present{{4, 3}, {3, 2}, {2, 1}},
	},
	{
		[]Present{
			{505, 23}, {352, 26}, {458, 20}, {220, 18}, {354, 32},
			{414, 27}, {498, 29}, {545, 26}, {473, 30}, {543, 27},
		},
		67,
		[]Present{{505, 23}, {220, 18}, {545, 26}},
	},
}

func TestGrabPresents(t *testing.T) {
	for i, test := range grabPresentsTests {
		result := grabPresents(test.presents, test.cap)
		expected := test.expected
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("#%d Expected:\n%v\n Got:\n%v", i+1, expected, result)
		}
	}
}

func BenchmarkTestGrabPresents(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range grabPresentsTests {
			grabPresents(test.presents, test.cap)
		}
	}
}
