package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Привет! Это простой калькулятор на Go!")
	fmt.Println("В нём есть 4 стандартные операции(+-/*), которые можно проводить над типом int.")
	fmt.Println("Если захотите выйти, введите 'exit'.")
	fmt.Printf("Приятного пользования:)\n\n")

	for {
		fmt.Println("Введите выражение в формате: 2+2, 2-2, 2*2 или 2/2")
		st := bufio.NewScanner(os.Stdin)
		st.Scan()
		if st.Text() == "exit" {
			break
		}
		var result int64 = 0
		var slice []string
		var isCool bool = false
		if strings.Contains(st.Text(), "+") {
			slice = strings.Split(st.Text(), "+")
			n1, err1 := strconv.ParseInt(slice[0], 10, 64)
			n2, err2 := strconv.ParseInt(slice[1], 10, 64)
			if err1 != nil || err2 != nil {
				fmt.Println("Некорректный ввод")
				continue
			}
			result = n1 + n2
			isCool = true
		} else if strings.Contains(st.Text(), "-") {
			slice = strings.Split(st.Text(), "-")
			n1, err1 := strconv.ParseInt(slice[0], 10, 64)
			n2, err2 := strconv.ParseInt(slice[1], 10, 64)
			if err1 != nil || err2 != nil {
				fmt.Println("Некорректный ввод")
				continue
			}
			result = n1 - n2
			isCool = true
		} else if strings.Contains(st.Text(), "*") {
			slice = strings.Split(st.Text(), "*")
			n1, err1 := strconv.ParseInt(slice[0], 10, 64)
			n2, err2 := strconv.ParseInt(slice[1], 10, 64)
			if err1 != nil || err2 != nil {
				fmt.Println("Некорректный ввод")
				continue
			}
			result = n1 * n2
			isCool = true
		} else if strings.Contains(st.Text(), "/") {
			slice = strings.Split(st.Text(), "/")
			n1, err1 := strconv.ParseInt(slice[0], 10, 64)
			n2, err2 := strconv.ParseInt(slice[1], 10, 64)
			if err1 != nil || err2 != nil || n2 == 0 {
				fmt.Println("Некорректный ввод")
				continue
			}
			result = n1 / n2
			isCool = true
		}
		if !isCool {
			r, err := strconv.ParseInt(st.Text(), 10, 64)
			if err != nil {
				fmt.Println("Некорректный ввод")
				continue
			}
			result = r
		}
		fmt.Printf("= %d\n", result)
	}

}
