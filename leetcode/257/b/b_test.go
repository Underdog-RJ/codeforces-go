// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	examples := [][]string{
		{
			`[[5,5],[6,3],[3,6]]`, 
			`0`,
		},
		{
			`[[2,2],[3,3]]`, 
			`1`,
		},
		{
			`[[1,5],[10,4],[4,3]]`, 
			`1`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, numberOfWeakCharacters, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-257/problems/the-number-of-weak-characters-in-the-game/
