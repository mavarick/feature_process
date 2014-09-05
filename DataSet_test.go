package main

import (
    "fmt"
    "testing"
)

func _TestDataSet(t *testing.T){
    filename := "/Users/apple/Documents/Kplan/src/51card-analyze/data/kplan.fea"
    ds := new(DataSet)
    ds.load_data(filename)
    fmt.Println(len(ds.Samples))
    for k,v := range ds.FeasIndexes{
        //fmt.Println(k, len(v))
        fmt.Println(k, len(v))
    }
    //for index, samp := range ds.Samples{
    //    fmt.Println(index, len(samp.Fea_dict))
    //}
    fea := "billdatepostponed"
    indexes := ds.get_fea_indexes(fea)
    labels := ds.get_labels(indexes)
    fmt.Println(labels) 
}
