package main

import (
    "fmt"
    "testing"
)

func _TestSample(t *testing.T){
    samp := new(Sample)
    line := "1\ta:1\tb:2"
    samp.parse(line)
    fmt.Println(samp.output())
    fmt.Println(samp.get("aaaa"))
}

