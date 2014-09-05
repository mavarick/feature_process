package main

import (
    "testing"
)

func _Test_gen_split_points(t *testing.T){
    filename := "/Users/apple/Documents/Kplan/src/51card-analyze/data/kplan.fea"
    sp_num := 5
    gen_split_points(filename, sp_num)
}
