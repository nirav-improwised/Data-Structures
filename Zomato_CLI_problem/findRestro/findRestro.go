package findRestro

import (
	"errors"
	"fmt"
	"main/methods"
	"main/structures"

	"github.com/gofiber/fiber/v2"
)

var Records = methods.GetRecords()
var dataStructs = methods.GeneratingStructs(Records)

func AddR(c *fiber.Ctx) error {
	fmt.Println("From AddR")
	var Userdata structures.Data
	var mainStr [][]string
	var existFlag bool
	err := c.BodyParser(&Userdata)
	if err != nil {
		fmt.Println("Error while body parsing ", err)
		return errors.New("error occured File: findRestro.go, Method: AddR")
	}
	existFlag = methods.CheckExistingData(dataStructs, Userdata)
	mainStr = methods.ConvertedData(Userdata)
	fmt.Println(Userdata)
	result := methods.AddData(mainStr, existFlag)
	return c.JSON(result)
}

func FindR(c *fiber.Ctx) error {
	fmt.Println("From FindR")

	var filteredData []structures.Data
	Cuisine := c.Query("cuisine")
	City := c.Query("city")
	HasTableBooking := c.Query("hasTableBooking")
	HasOnlineDelivery := c.Query("hasOnlineDelivery")
	// dataStructs := methods.GeneratingStructs(Records)
	filteredData = methods.StudyData(dataStructs, Cuisine, City, HasTableBooking, HasOnlineDelivery)

	return c.JSON(filteredData)
}

func DeleteR(c *fiber.Ctx) error {
	id := c.Query("id")
	result := methods.DeleteData(id, Records)
	return c.JSON(result)
}

func UpdateR(c *fiber.Ctx) error {
	var QueryData structures.Data
	err := c.BodyParser(&QueryData)
	if err != nil {
		fmt.Println(err)
	}
	message := methods.UpdateData(QueryData, Records)
	return c.JSON(message)
}
