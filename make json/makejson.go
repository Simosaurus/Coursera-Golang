package main

import (
  "bufio"
  "os"
  "fmt"
  "encoding/json"

)

func main() {
    fmt.Printf("Enter a name:")
    in := bufio.NewScanner(os.Stdin)
    in.Scan()
    Name := in.Text()
    fmt.Printf("Enter an address:")
    in2 := bufio.NewScanner(os.Stdin)
    in2.Scan()
    Address := in2.Text()
    Person := map[string]string {Name:Address}
    barr, err := json.Marshal(Person)
    if err != nil {
        panic(err)
    }
    fmt.Println("Here's the JSON created: \n",string(barr))
    }
