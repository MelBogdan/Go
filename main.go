package main

import "fmt"
import "math"
import "os"
	
// + - * / radical square
func main(){
	var a, b, res float64
	var oper string

	fmt.Println("Введите первое число: ")
	fmt.Scanln(&a)

	fmt.Println("Введите оператор или функцию: ")
	fmt.Scanln(&oper)

	if !(oper == "radical") && !(oper == "square"){
		fmt.Println("Введите второе число: ")
		fmt.Scanln(&b)
	}

	switch oper {
		case "+":
			res = a + b
		case "-":
			res = a - b
		case "*":
			res = a * b
		case "/":
			if b == 0{
				fmt.Println("Деление на ноль невозможно")
				os.Exit(1)
			}else{
				res = a/b
			}
		case "radical":
			res = math.Sqrt(a)
		case "square":
			fmt.Println("Введите степень числа: ")
			fmt.Scanln(&b)
			res = math.Pow(a, b)
		default:
			fmt.Println("Операция не найдена")
			os.Exit(1)
	} 

	fmt.Printf("Результат выполнения операции: %f\n",res)
}
