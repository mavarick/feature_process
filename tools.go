package main

import (
    "sort"
)

// find the splitting points in iterdata, make each part have equal points number
func split_equal_num_opt(iterdata []float64, num int) []float64{
    // sort the data
    sorted_data := iterdata[0:len(iterdata)]
    sort.Float64s(sorted_data)
    total_length := len(sorted_data)
    d_length := len(sorted_data)
    d_num := num
    n := 0
    split_points := []float64{}
    d_index := -1
    
    // when the unique values count is smaller than num, it will get to end earlier
    is_end := 0
    for{
        if n == num {
            break
        }
        index := d_length / d_num
        d_index += index
        // find the end index of the value
        split_value := sorted_data[d_index]
        for i:=1;;i++{
            _d_index := d_index + i
            if _d_index >= total_length {
                is_end = 1
                break 
            }
            
            if sorted_data[_d_index] > split_value{
                d_index = _d_index - 1
                break
            }
        }
        if is_end == 1{
            if len(split_points) < num - 1{
                split_points = append(split_points, split_value)
            }
            break
        }
            
        split_points = append(split_points, split_value)

        d_length = total_length - 1 - d_index
        d_num -= 1

        n += 1
    }
    return split_points
}

func get_index(value float64, sorted_array []float64, start_index int) int{
    index := 0
    for i:=0;i<len(sorted_array);i++{
        if value <= sorted_array[i]{
            index = i
            break
        }
        if i == len(sorted_array) - 1 {
            index = len(sorted_array)
        }
    }
    return index + start_index
}

func round(v float64) int64{
    compensation := 0.5
    if v < 0{
        compensation *= -1
    }
    return int64(v + compensation)
}
