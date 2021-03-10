package main

import "fmt"

func main() {
    var num0 = 0
    var num1 = 1
    var num2 = 0 

    for i := 0; i < 100; i++ {
       num0 = num1
       num1 = num2
       num2 = num1 + num0;
       fmt.Printf("%d\n", num2)
    }
}
