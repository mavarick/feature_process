package main
import (
    "fmt"
    "strings"
    "os"
)

// generate splitting points of feature and print it on console
func gen_split_points(filename string, split_num int, out_file string){
    ds := new(DataSet)
    ds.load_data(filename)

    fp, err := os.OpenFile(out_file, os.O_RDWR|os.O_CREATE,0644)
    if err != nil{
        panic(fmt.Sprintf("Error: Can not open or create file: %s", out_file))
    }
    defer fp.Close()
    for k, _ := range ds.FeasIndexes{
        fea_values := ds.get_fea_values(k)
        split_points := split_equal_num_opt(fea_values, split_num)
        print_lst := []string{}
        for _, v := range split_points{
            print_lst = append(print_lst, fmt.Sprintf("%f", v))
        }
        
        fp.WriteString(fmt.Sprintf("%s\t%s", k, strings.Join(print_lst, ",")))
        fp.WriteString("\n")
    }
}

