package main

import (
  "bufio"
  "os"
  "fmt"
  "strings"
)

func main() {
    fmt.Printf("Enter a file name: ")
    in := bufio.NewScanner(os.Stdin)
    in.Scan()
    FileName := in.Text()
    FileData, e := os.Open(FileName)
    if e != nil {
        fmt.Printf("Error - Check file placement or file name and try program again.")
    }
     type name struct  {
        fname string
        lname string
    }
    nameslice := make([]name,0)
    scanner := bufio.NewScanner(FileData)
    for scanner.Scan() {
        FullNameText := scanner.Text()
        var names = strings.Fields(FullNameText)
        var name1 string = names[0]
        var name2 string = names[1];
        person := new(name)
        person.fname = name1
        person.lname = name2
        nameslice = append(nameslice,*person)
    }
    for _, i := range nameslice {
    fmt.Println(i.fname,i.lname)
}
    }
