package main

import "fmt"

type Product struct {

	//กำหนดว่าจะเก็บข้อมูลอะไรบ้าง
	name     string
	price    float64
	category string
	discount int
}

func main() {

	product1 := Product{name: "เก้าอี้", price: 500.25, category: "เครื่องใช้", discount: 10}
	product1.price = 100
	fmt.Println(product1)

	//fmt.Println("Test123")

	// var name string = "test" //ค่าที่มีการตั้งตัวแปรแบบนี้จะสามารถเปลี่ยนแปลงค่าได้ แบบไม่คงที่
	// name = "jojo"
	// age := 25
	// var score float32 = 95.8
	// fmt.Println("this is a", name)
	// fmt.Printf(" age = %T\n", age)
	// fmt.Println("Score is ", score)

	// var score2 int
	// fmt.Print("score = ")
	// fmt.Scanf("%d", &score2)

	// if score2 >= 100 {
	// 	fmt.Println("pass")
	// } else {
	// 	fmt.Println("not pass")
	// }

	// switch score2 {
	// case 50:
	// 	fmt.Println("half")
	// case 100:
	// 	fmt.Println("full")
	// default:
	// 	fmt.Println("final new")
	// }

}