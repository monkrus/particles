package main

import (
    "bufio"
    "fmt"
    "os"
)

reader := bufio.NewReader(os.Stdin)
fmt.Print("Enter text: ")
text, _ := reader.ReadString('\n')
fmt.Println(text)