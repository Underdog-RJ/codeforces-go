// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	examples := [][]string{
		{
			`[10,6,8,5,11,9]`, 
			`[3,1,2,1,1,0]`,
		},
		{
			`[5,1,2,3,10]`, 
			`[4,1,1,1,0]`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, canSeePersonsCount, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-57/problems/number-of-visible-people-in-a-queue/
