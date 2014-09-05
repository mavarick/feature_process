package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

// split the feas with splitting points dict
func split(fea_dict map[string]float64, fea_sp_dict map[string][]float64) map[string]float64{
    new_fea_dict := map[string]float64{}
    for f, v := range fea_dict{
        split_points, is_exist := fea_sp_dict[f]
        if is_exist == false{
            continue
        }
        split_index := float64(get_index(v, split_points, 1))
        new_fea_dict[f] = split_index
    }
    return new_fea_dict
}

func add_missing(fea_dict map[string]float64, missing_feas []string, missing_value float64) map[string]float64{
    for _, fea := range missing_feas{
        fea_dict[fea] = missing_value
    }
    return fea_dict
}

func extend_feature(fea_dict map[string]float64) map[string]float64{
    new_fea_dict := map[string]float64{}
    for k, v := range fea_dict{
        new_f := fmt.Sprintf("%s_%d", k, round(v))
        new_v := float64(1)
        new_fea_dict[new_f] = new_v
    }
    return new_fea_dict
}

func add_bias(fea_dict map[string]float64, bias_name string, bias_value float64) map[string]float64{
    fea_dict[bias_name] = bias_value
    return fea_dict
}

func transform_feature(fea_dict map[string]float64, fea_sp_dict map[string][]float64, total_fs_set []string) map[string]float64{
    new_fea_dict := split(fea_dict, fea_sp_dict)

    missing_feas := []string{}
    for _, fea := range total_fs_set{
        _, is_exist := new_fea_dict[fea]
        if is_exist == false{
            missing_feas = append(missing_feas, fea)
        }
    }

    var missing_value float64 = -1
    new_fea_dict = add_missing(new_fea_dict, missing_feas, missing_value)
    new_fea_dict = extend_feature(new_fea_dict)

    var bias_name string = "bias"
    var bias_value float64 = 1
    new_fea_dict = add_bias(new_fea_dict, bias_name, bias_value)
    return new_fea_dict
}

func process_feature(feature_file string, split_points_file string, out_file string){
    fea_sp_dict := read_split_points(split_points_file)
    total_fs_set := []string{}
    for k,_ := range fea_sp_dict{
        total_fs_set = append(total_fs_set, k)
    }
    // read the feature data
    fp, err := os.Open(feature_file)
    if err != nil{
        panic(fmt.Sprintf("Error When Open file: %s", feature_file))
    }
    defer fp.Close()
    reader := bufio.NewReader(fp)
    // create the writing file
    fp_out, err := os.OpenFile(out_file, os.O_RDWR|os.O_CREATE,0644)
    defer fp_out.Close()
    n := 0
    for{
        line, err := reader.ReadString('\n')
        if err != nil{
            fmt.Println(err)
            break
        }
        samp := new(Sample)
        samp.parse(line)
        new_fea_dict := transform_feature(samp.Fea_dict, fea_sp_dict, total_fs_set)
        samp2 := new(Sample)
        samp2.Label = samp.Label
        samp2.Fea_dict = new_fea_dict
        //fmt.Println(samp2.output())
        fp_out.WriteString(samp2.output())
        fp_out.WriteString("\n")
        n += 1
    }
    fmt.Println("Total Line Count: ", n)
}

func read_split_points(filename string) map[string][]float64{
    fp, err := os.Open(filename)
    if err != nil{
        panic(fmt.Sprintf("Error When openning file: %s", filename))
    }
    defer fp.Close()

    reader := bufio.NewReader(fp)
    fea_sp_dict := map[string][]float64{}
    for{
        line, err := reader.ReadString('\n')
        if err != nil{
            fmt.Println(err)
            break
        }
        fea, sps := parse_line(line)
        fea_sp_dict[fea] = sps
    }
    return fea_sp_dict

}

func parse_line(line string)(string, []float64){
    items := strings.Split(strings.Trim(line, "\n"), "\t")
    fea_name := strings.Trim(items[0], " ")
    sps_str := strings.Split(items[1], ",")
    sps := []float64{}
    for _, v := range sps_str{
        sp, err := strconv.ParseFloat(strings.Trim(v, " "), 8)
        if err != nil{
            panic(fmt.Sprintf("Error when Converse value `%s` to float64", v))
        }
        sps = append(sps, sp)
    }
    return fea_name, sps
}


//func output(label int, fea_dict map[string]float64) string




