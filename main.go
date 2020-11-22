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
		log.Fatal("First line should contains only 2 numbers\n")
	}

	//read from  input
	for {
		fmt.Print("Numbers: ")
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
        fmt.Println(text)
		input = append(input, text)
		
		} else {
			break
		
		}
	}
	fmt.Println(input)
}

   // split string by spaces
   s:= strings.Split (input)

   // check slice has N elements
   func Find(input []string, val string) (int, bool) {
    for i, item := range input {
        if item == val {
            return i, true
        }
    }
    return -1, false
}

   
// 4. Преобразовать элементы массива в числа, первый элемент - L, второй - N

// 5. Создать строковый срез на N элементов

 
// 6. В цикле считать N строк
 
// а) если больше строк нет - вернуть ошибку
 
// б) если строка имеет длину не равную L - вернуть ошибку
 
// в) добавить строку в срез

  

	
