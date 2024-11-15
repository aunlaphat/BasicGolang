package main

import (
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func getBook(c *fiber.Ctx) error {
	return c.JSON(book)
}

func getBookID(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	//_ เรียก blank แทนค่าของ index เมื่อไม่ได้ใส่ค่าลงไป , range ใช้วน loop
	for _, b := range book {
		if b.ID == bookId {
			return c.JSON(b)
		}
	}

	//หากไม่เข้าเงื่อนไขใดเลย แสดงไม่พบค่า ไม่เพิ่มข้อความอันแรกใช้			เพิ่มข้อความเองในการ error ของประเภทนี้
	return c.Status(fiber.StatusNotFound).SendString("Not found this ID")
}

// choose post and select body > JSON
func createBook(c *fiber.Ctx) error {
	newbook := new(Book)
	if err := c.BodyParser(newbook); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	book = append(book, *newbook)
	return c.JSON(newbook)
}

func updateBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	bookUpdate := new(Book)
	if err := c.BodyParser(bookUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, b := range book {
		if b.ID == bookId {
			book[i].Title = bookUpdate.Title
			book[i].Author = bookUpdate.Author
			return c.JSON(b)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)

}

func deleteBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, b := range book {
		if b.ID == bookId {
			book = append(book[:i], book[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}

// run in body > form-data > put text image on column Key and select choice file
func uploadBook(c *fiber.Ctx) error {
	file, err := c.FormFile("image")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	//สร้างโฟลเดอที่จะให้เข้าไป path นั้นด้วย เช่น ./uploads/ ก็ต้องมีโฟลเดอ uploads
	err = c.SaveFile(file, "./uploads/"+file.Filename)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString("Upload file success")
}

func testHtml(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		//ทำการ map key ของฟิลด์ที่มีฟิลด์ตรงกับ index pages เพื่อส่งค่า key ออกไปแสดง index
		"Title": "Test func",
		"Name":  "AUnlaphat",
	})
}

// stamp value on dev use command >"set SECRET=1234" check value >"set"
// stamp value again >"set SECRET=1234" go on postman
func getEnv(c *fiber.Ctx) error {
	if value, exists := os.LookupEnv("SECRET"); exists {
		return c.JSON(fiber.Map{
			"SECRET": value,
		})
	}

	return c.JSON(fiber.Map{
		"SECRET": "defaultsecret",
	})
}
