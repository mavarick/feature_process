package main
import (
    "os"
    "bufio"
    "fmt"
)

type DataSet struct{
    Samples []*Sample
    FeasIndexes map[string][]int
}

func (self *DataSet) load_data(filename string){
    fp, err := os.Open(filename)
    if err != nil{
        panic(fmt.Sprintf("%s %s", 
            "Can not open file: ", filename))
    }
    defer fp.Close()
    reader := bufio.NewReader(fp)

    // init
    self.Samples = []*Sample{}  // or make([]Sample, 0)
    self.FeasIndexes = make(map[string][]int)
    n := 0
    for {
        line, err := reader.ReadString('\n')
        if err != nil{
            fmt.Println(err)
            break
        }
        samp  := new(Sample)
        samp.parse(line)
        self.Samples = append(self.Samples, samp)
        for k, _ := range samp.Fea_dict{
            self.FeasIndexes[k] = append(self.FeasIndexes[k], n)
        }
        n += 1
    }
    fmt.Println("Total Number: ", n)
}

func (self *DataSet) get_fea_indexes(fea string) []int{
    value, ok := self.FeasIndexes[fea]
    if !ok {
        panic(fmt.Sprintf("%s %s",
            "DataSet Does not Contains Feature", fea))
    }
    return value
}

func (self *DataSet) get_fea_values(fea string) []float64{
    indexes := self.get_fea_indexes(fea)
    values := []float64{}
    for _, index := range indexes{
        value, missing_flag := self.Samples[index].get(fea)
        if missing_flag == 1 {
            continue
        }
        values = append(values, value)
    }
    return values
}

func (self *DataSet) get_labels(indexes []int) []int64 {
    labels := []int64{}
    for _, v := range indexes{
        label := self.Samples[v].Label
        labels = append(labels, label)
    }
    return labels
}



