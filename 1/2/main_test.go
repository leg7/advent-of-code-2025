package main

import (
	"bufio"
	"os"
	"testing"
)

var file string = "../input.txt"

func BenchmarkBranchless(b *testing.B) {
	for b.Loop() {
		t, _ := os.Open(file)
		r := bufio.NewReader(t)
		Branchless(r)
	}
}

func BenchmarkBranchy(b *testing.B) {
	for b.Loop() {
		t, _ := os.Open(file)
		r := bufio.NewReader(t)
		Branchy(r)
	}
}

