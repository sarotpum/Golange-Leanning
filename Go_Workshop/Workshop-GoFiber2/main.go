package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"

	Database "github.com/workshop-gofiber2/database"
	Model "github.com/workshop-gofiber2/model"
)

func main() {
	app := fiber.New()
	app.Use(middleware.Logger("${blue} ${time} ${red} ${cyan} ${method} ${path} ${yellow} ${status} ${green} ${ip} ${ua}${resetColor}\n"))

	app.Post("api/test01", func(c *fiber.Ctx) {
		c.Status(200).JSON(&fiber.Map{
			"message": "Hello world.",
		})
	})

	app.Post("api/create_table", func(c *fiber.Ctx) {
		db, err := Database.OpenPgSql()
		if err != nil {
			fmt.Println("Error :", err)
		}
		qr := fmt.Sprintf(`CREATE TABLE USERS(
			UserId varchar(50),
			FirstName varchar(50),
			LastName varchar(50),
			Age integer,
			Status integer,
			CreateAt varchar(50)
		)`)

		create, err := db.Exec(qr)
		_ = create

		c.Status(200).JSON(&fiber.Map{
			"message": "Create done.",
		})
	})

	app.Post("api/insert", func(c *fiber.Ctx) {
		req := new(Model.UserModelReq)

		fmt.Println("start :", req)
		if err := c.BodyParser(req); err != nil {
			fmt.Println(err)
			c.Status(400).JSON(&fiber.Map{
				"status":  400,
				"message": err,
			})
			return
		}

		fmt.Println("start2 :", req)
		db, err := Database.OpenPgSql()
		if err != nil {
			fmt.Println(err)
		}

		qr := fmt.Sprintf(`INSERT INTO USERS(UserId, FirstName, LastName, Age, Status, CreateAt)  `+
			`VALUES (uuid_generate_v4(), '%v', '%v', '%d', '%d', current_timestamp) `, req.FirstName, req.LastName, req.Age, req.Status)
		
		
		fmt.Println("qr : ", qr)
		
		insert, err := db.Exec(qr)
		_ = insert

		c.Status(200).JSON(&fiber.Map{
			"message": "Insert done.",
		})
	})

	app.Post("api/update", func (c *fiber.Ctx) {
		req := new(Model.UserModelUpdate)
		if err := c.BodyParser(req); err != nil {
			fmt.Println(err)
			c.Status(400).JSON(&fiber.Map{
				"status": 400,
				"message": err,
			})
			return
		}

		db, err := Database.OpenPgSql()
		if err != nil {
			fmt.Println(err)
		}
	
		qr := fmt.Sprintf(`UPDATE users ` + 
			`SET FirstName='%v', LastName='%v', Age=%d, Status=%d, CreateAt=current_timestamp ` + 
			`WHERE UserId='%v' `, req.FirstName, req.LastName, req.Age, req.Status, req.UserId)

		update, err := db.Exec(qr)
		_ = update

		c.Status(200).JSON(&fiber.Map{
			"message": "Update done",
		})
	})

	app.Listen(":3000")
}
