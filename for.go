package main

import "fmt"

func main()  {
    // 1st for
    fmt.Println("1st Loop for")
    for idx := 0; idx < 9; idx++ {
        if idx == 2 {
            fmt.Println("ini Dua")
            continue
        }
        fmt.Println(idx)
    }

    // 2nd for
    fmt.Println("2nd Loop for")
    idx := 0
    for  idx < 5 {
        fmt.Println(idx)
        idx = idx + 1
    }
}
