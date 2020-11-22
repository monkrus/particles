package main

import (
	"bufio"
	"fmt"
	"os"
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

	//read from  input
	for   _, s := range numbers {
		convert = ints32(s)
		fmt.Println(convert)
	}



		



	

   
// 4. Преобразовать элементы массива в числа, первый элемент - L, второй - N

// 5. Создать строковый срез на N элементов

 
// 6. В цикле считать N строк
 
// а) если больше строк нет - вернуть ошибку
 
// б) если строка имеет длину не равную L - вернуть ошибку
 
// в) добавить строку в срез

  

	
