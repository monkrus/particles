package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Cell int

const (
	OneRock Cell = iota
	TwoRocks
	Table
	EmptySpace
)

//карта ключ - значение
var Keys = map[rune]Cell{
	'.': OneRock,
	':': TwoRocks,
	'T': Table,
	' ': EmptySpace,
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numbers := strings.Split(scanner.Text(), " ")
	if len(numbers) != 2 {
		log.Fatal("First line should contains only 2 numbers\n")
	}
	L, err := strconv.Atoi(numbers[0])
	checkError(err)
	N, err := strconv.Atoi(numbers[1])
	checkError(err)
	input := make([]string, N)
	for i := 0; i < N; i++ {
		if !scanner.Scan() {
			log.Fatal("Unexpected EOF\n")
		}
		text := scanner.Text()
		if len(text) != L {
			log.Fatalf("Each line should contains exactly %d charactes\n", L)
		}
		for idx, v := range text {
			//если символ не найден в карте, значит это лишний символ
			if _, ok := Keys[v]; !ok {
				log.Fatalf("%d row contains incorrect symbol %c at %d position\n", i, v, idx)
			}
		}
		input[i] = text
	}
	//строки в Go иммутабельны, поэтому выходные данные храним в двумерном срезе символов
	output := make([][]byte, N)
	//выделяем память
	for i := 0; i < N; i++ {
		output[i] = make([]byte, L)
	}
	//меняем строки местами, чтобы внизу оказались верхние строки
	reverse(input)

	for i := 0; i < L; i++ {
		//позиции плит
		//в Go нет типа для структуры данных множество (set)
		//поэтому используют карту, хранящую пустую структуру
		tables := map[int]struct{}{}
		//ключ - начальная позиция, значение - количество камней
		caps := map[int]int{}
		cur := 0 //начальная позиция "стопки" камней, либо 0 - земля, либо числовое значение обозначающее верх плиты
		for j := 0; j < N; j++ {
			if Keys[rune(input[j][i])] == Table {
				tables[j] = struct{}{} //добавили к плитам
				cur = j + 1            //перенесли указатель на начало "стопки"
			}

			if Keys[rune(input[j][i])] == OneRock {
				caps[cur] += 1 //одиночный камень
			}

			if Keys[rune(input[j][i])] == TwoRocks {
				caps[cur] += 2 //двойной камень
			}
		}
		//конструируем строку по значениям весов и плит
		temp := construct(caps, tables, N)
		for j := 0; j < N; j++ {
			output[N-j-1][i] = temp[j] //переворачиваем в обратную сторону
		}
	}
	fmt.Println("======================================================")
	for _, str := range output {
		fmt.Println(string(str)) //выводим в виде строки
	}
}

func construct(caps map[int]int, tables map[int]struct{}, n int) string {
	arr := make([]rune, n)
	//по-умолчанию все пробелы
	for i, _ := range arr {
		arr[i] = ' '
	}
	//переносим плиты
	for k, _ := range tables {
		arr[k] = 'T'
	}

	//считаем веса камней
	//ключ обозначает начальную позицию
	for k, v := range caps {
		//камни кладем по двое
		for i := 0; i < v; i += 2 {
			arr[k+i/2] = ':'
		}
		//если камней нечетное число
		if v%2 != 0 {
			arr[k+v/2] = '.'
		}
	}
	//возвращаем строку
	return string(arr)
}

//copied from https://stackoverflow.com/a/34816623
func reverse(ss []string) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
