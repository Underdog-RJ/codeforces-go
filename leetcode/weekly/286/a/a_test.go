// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_a(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithFile(t, findDifference, "a.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-286/problems/find-the-difference-of-two-arrays/
