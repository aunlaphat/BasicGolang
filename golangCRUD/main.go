package main

import (
	// "fmt"
	// "log"
	// "net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var book []Book

//w http.ResponseWriter = ทำให้ตอบกลับส่งค่าไปในฝั่งของ client
//r *http.Request = ข้อมูลที่มีการส่งมา

//เป็นโค้ดทำมือที่ใช้เวลามากเกินไป ใช้ตัวช่วยจาก service อื่นเพื่อย่อยโค้ดให้สั้นลงแต่มีการทำงานเท่าเดิม
/* func helloHandler( w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET"{
		http.Error(w, "Method is not support.", http.StatusNotFound)
		return
	}

	//ทำงานเมื่อมันไม่หยุดหา เงื่อนไขก่อนหน้า
	fmt.Fprintf(w, "Hello world!")
} */

// func main() {
// 	http.HandleFunc("/hello", helloHandler)

// 	fmt.Println("Start server at port 8080\n")
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatal(err)
// 	}

// }

func main() {

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	//จาก func helloHandler to fiber.New()
	//app := fiber.New()

	book = append(book, Book{ID: 1, Title: "Harry Potter", Author: "J.K.Rolling"})
	book = append(book, Book{ID: 2, Title: "The Conjuring", Author: "P.James"})

	app.Get("/book", getBook)

	//เรียกเป็นแต่ละ id
	app.Get("/book/:id", getBookID)
	app.Get("/test-html", testHtml)
	app.Get("/api/config", getEnv)

	//create
	app.Post("/createBook", createBook)
	app.Post("/upload", uploadBook)
	//update
	app.Put("/updateBook/:id", updateBook)

	//delete
	app.Delete("/deleteBook/:id", deleteBook)
	// print(book)

	// app.Get("/book", func(c *fiber.Ctx) error {
	// 	return c.JSON(book)
	// })
	app.Get("/hello123", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.Listen(":8080")

}
