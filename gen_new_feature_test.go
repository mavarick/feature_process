package main

import (
    "testing"
)

func Test_gen_new_feature(t *testing.T){
    filename := "/Users/apple/Documents/Kplan/src/51card-analyze/data/kplan.fea"
    fea_sp_file := "/Users/apple/Documents/Kplan/src/feature_selection/split_points_r.txt"
    out_file := "kplan_r.fea"
    process_feature(filename, fea_sp_file, out_file)
}
