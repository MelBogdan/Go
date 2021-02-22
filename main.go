package main

import "fmt"
import "math"
	

func main(){
	var firstSide,secondSide,thirdSide float64;
	fmt.Println("Введите значение первой стороны: ")
	fmt.Scan(&firstSide)
	fmt.Println("Введите значение второй стороны: ")
	fmt.Scan(&secondSide)
	fmt.Println("Введите значение третьей стороны: ")
	fmt.Scan(&thirdSide)
	triangle(firstSide,secondSide,thirdSide)


	var space float64
	fmt.Println("Введите площадь окружности: ")
	fmt.Scan("&space")
	circle(space)

	var number int
	fmt.Println("Введите трехзначное число")
	fmt.Scan(&number)
	num(number)

}

const(
	pi = math.Pi
)

func triangle(a,b,c float64) {
	var p float64 = (a + b + c) / 2
	var total = math.Sqrt(p * (p - a) * (p - b) * (p - c))
	fmt.Println("Площадь треугольника равна: ", total)
}

func circle(a float64){
	var dia = 2 * math.Sqrt(a / pi)
	var lenght = 2 * pi * (dia / 2) 
	fmt.Println("Диаметр окружности равен: " ,dia , " Длина окружности равна: " , lenght)
}

func num(a int){
	var array = [3]string{"Единицы", "Десятки", "Сотни"}
	for i := 0; i < 3; i++{
		fmt.Println(array[i],": ",a%10);
		a /= 10;
	}
}
