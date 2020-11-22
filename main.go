package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	// make an empty slice
	input := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numbers := strings.Split(scanner.Text(), " ")
	if len(numbers) != 2 {
		log.Fatal("First line should contain only 2 numbers\n")
	}

	//read  and   save from  input
	L, err := strconv.Atoi(numbers[0])
    checkError(err)
	N, err := strconv.Atoi(numbers[1])
	checkError(err)


   // create slice with N elements
	var s []int	
	
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	if len(s) = 0 {
		log.Fatal("The slice contains no strings")
	}
	else len(s) != L(numbers) {
		log.Fatal("????")
	}
    









// 5. Создать строковый срез на N элементов

 
// 6. В цикле считать N строк
 
// а) если больше строк нет - вернуть ошибку
 
// б) если строка имеет длину не равную L - вернуть ошибку
 
// в) добавить строку в срез

  

	
