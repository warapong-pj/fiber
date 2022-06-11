package main

import (
	"fmt"
	"log"
	"os"
	"tidy/domain"
	"tidy/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("HOST"),
		os.Getenv("USERNAME"),
		os.Getenv("PASSWORD"),
		os.Getenv("DATABASE"),
		os.Getenv("PORT"),
		os.Getenv("SSL"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	user := domain.User{}
	db.AutoMigrate(&user)

	return db
}

func main() {
	db := initDB()

	app := fiber.New()
	app.Use(logger.New(logger.Config{Format: "${time} ${locals:requestid} ${ip} ${status} - ${method} ${path} ${latency}\n"}))
	app.Use(requestid.New())
	app.Use(pprof.New())

	app.Get("/metrics", monitor.New(monitor.Config{
		APIOnly: true,
	}))

	domain := domain.NewUserRepositoryDB(db)
	service := service.NewUserService(domain)
	handle := handle.NewUserHandle(service)

	if handle != nil {
		fmt.Println("test")
	}

	// app.Get("/:id", func(c *fiber.Ctx) error {
	// 	user := new(User)
	// 	id := c.Params("id")

	// 	db.First(&user, id)

	// 	return c.Status(200).JSON(user)
	// })

	// app.Post("/", func(c *fiber.Ctx) error {
	// 	user := new(User)
	// 	err := c.BodyParser(user)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	db.Create(&User{
	// 		Name:  user.Name,
	// 		Email: user.Email,
	// 	})

	// 	return c.SendStatus(fiber.StatusCreated)
	// })

	// app.Patch("/:id", func(c *fiber.Ctx) error {
	// 	user := new(User)
	// 	updateUser := new(User)
	// 	id := c.Params("id")
	// 	err := c.BodyParser(updateUser)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	db.First(&user, id)
	// 	user.Name = updateUser.Name
	// 	user.Email = updateUser.Email
	// 	db.Save(&user)

	// 	return c.Status(200).JSON(user)
	// })

	// app.Delete("/:id", func(c *fiber.Ctx) error {
	// 	users := new([]User)
	// 	id := c.Params("id")

	// 	db.Find(&users)
	// 	db.Where("id = ?", id).Delete(&users)

	// 	return c.SendStatus(fiber.StatusOK)
	// })

	log.Fatal(app.Listen(":3000"))
}
