package main

import (
    "fmt"
    "flag"
    "strconv"
    "time"
    "syscall"
)



func main(){
    beg_t := time.Now()
    params := prepare_params()

    process_type := params["type"]
    fmt.Println(params)
    if process_type == "split" {
        // generate spltting points
        handle_split_points(params)
        fmt.Println("Generating split points is Done")
    }else if process_type == "fea_process"{
        // preprocessing feature
        preprocess_data(params)
        fmt.Println("Preprocessing feature is Done")
    }
    end_t := time.Now()
    fmt.Printf("Time Used: %v \n", end_t.Sub(beg_t))
}

func preprocess_data(params map[string]string){
    feature_file := params["fea"] //"/Users/apple/Documents/Kplan/src/51card-analyze/data/kplan.fea"
    fea_sp_file := params["spfile"] //"split_points.txt"
    out_file := params["out"]   //

    process_feature(feature_file, fea_sp_file, out_file)
}

func handle_split_points(params map[string]string){
    feature_file := params["fea"] //"/Users/apple/Documents/Kplan/src/51card-analyze/data/kplan.fea"
    fea_sp_file := params["spfile"] //"split_points.txt"
    sp_num, err := strconv.ParseInt(params["spnum"], 10, 4)
    if err != nil{
        panic(fmt.Sprintf("Error happened when parsing param 'spnum' to int, the value is %s", params["spnum"]))
    }
    gen_split_points(feature_file, int(sp_num), fea_sp_file)
}

func prepare_params() map[string] string{
    const (
        default_process_type  =  "fea_process"
        usage_process_type    =  `1) fea_process : given fea_file/sp_file, to generate new feature file to out_file
                                  2) split : given fea_file and split number, to generate splitting points to sp_file`

        default_fea_file      =  "/Users/apple/Documents/Kplan/src/51card-analyze/data/kplan.fea"
        usage_fea_file        =  "the feature file path"

        default_spnumber      =  "5"
        usage_spnumber        =  "Specify the split number"

        default_sp_file       =  "split_points.txt"
        usage_sp_file         =  "the output file of spliting points"

        default_new_fea_file  =  "fea_new.fea"
        usage_new_fea_file    =  "the output of feature processing result"

        usage =
`USAGE examples: 
    To generate the gen_split_points.txt: run below
    ./feature_selection -type=split -fea={fea.file} -spnum={5} -spfile={gen_split_points.txt}"
    To generate new feature file: type below: 
    ./feature_selection -type=fea_process -fea={fea.file} -spfile={gen_split_points.txt} -out={new_fea.fea}
`
    )
    params := make(map[string]string)
    process_type := flag.String("type", default_process_type, usage_process_type)
    feature_file_path := flag.String("fea", default_fea_file, usage_fea_file)
    split_number := flag.String("spnum", default_spnumber, usage_spnumber)
    sp_file := flag.String("spfile", default_sp_file, usage_sp_file)
    new_fea_file := flag.String("out", default_new_fea_file, usage_new_fea_file)

    var usage_flag bool
    flag.BoolVar(&usage_flag, "usage", false, "usage message")

    flag.Parse()
    
    if usage_flag{
        fmt.Println(usage)
        syscall.Exit(0)
    }
    
    params["type"] = *process_type
    params["fea"] = *feature_file_path
    params["spnum"] = *split_number
    params["spfile"] = *sp_file
    params["out"] = *new_fea_file

    return params
}


