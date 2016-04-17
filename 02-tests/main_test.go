package main

import "testing"

// func Test(t *testing.T) {
// 	t.Log("Hello World")
// 	t.Fail()
// }

func BenchmarkPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		multStuff(1, 1)
	}
}
