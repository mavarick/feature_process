package main
import (
    "fmt"
    "sort"
    "testing"
)

func _TestSplitEqualNumOpt(t * testing.T){
    iterdata := []float64{1,2,3,4,5,6,7,8,9,19,11,10,12,13,14,15,16,17,18}
    //iterdata := []float64{1,2,3, 1,2,3,1,2,3}
    num := 5
    sp := split_equal_num_opt(iterdata, num)
    sorted_data := iterdata[0:len(iterdata)]
    sort.Float64s(sorted_data)
    fmt.Println(sorted_data)
    fmt.Println(sp)
}

func _TestGet_index(t *testing.T){
    iterdata := []float64{1,2,3,4}
    index := get_index(3, iterdata, 1)
    if index != 3{
        t.Error("Get Wrong result")
    }
    index = get_index(1, iterdata, 1)
    if index != 1{
        t.Error("Get Wrong result")
    }
    index = get_index(4, iterdata, 1)
    if index != 4{
        t.Error("Get Wrong result")
    }
    index = get_index(10, iterdata, 1)
    fmt.Println(index)
    if index != 5{
        t.Error("Get Wrong result")
    }
}

