package main
import(
    "fmt"
    "strings"
    "strconv"
)

type Sample struct{
    Label int64
    //Fea_dict map[string]float64 
    Fea_dict map[string]float64
}

func (this *Sample) parse(line string){
    items := strings.Split(strings.Trim(line, "\n"), "\t")
    label, ok := strconv.ParseInt(items[0], 10, 4)
    if ok != nil{
        panic(fmt.Sprintf("%s %d", 
            "Error happend when transform label: ", items[0]))
    }
    this.Label = label
    n := 1
    this.Fea_dict = make(map[string]float64)
    for n < len(items){
        fv := strings.Split(items[n], ":")
        value, ok := strconv.ParseFloat(fv[1], 64)
        if ok != nil{
            panic(fmt.Sprintf("%s %s",
                "Error happend when transform feature value: ", fv[1]))
        }
        this.Fea_dict[fv[0]] = value
        n += 1
    }
}

func (this *Sample) output() string{
    print_lst := []string{}
    print_lst = append(print_lst, strconv.FormatInt(this.Label, 10))
    
    for k, v := range this.Fea_dict{
        new_fv := fmt.Sprintf("%s:%d", k, int(v))
        print_lst = append(print_lst, new_fv)
    }
    return strings.Join(print_lst, "\t")
}

func (this *Sample) get(fea string)(float64,int){
    missing_flag := 0
    value, ok := this.Fea_dict[fea]
    if ok == false{
        missing_flag = 1
    }
    return value, missing_flag
}




