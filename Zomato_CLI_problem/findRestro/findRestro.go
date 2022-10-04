package findRestro

import (
	"errors"
	"fmt"
	"main/methods"
	"main/structures"

	"github.com/gofiber/fiber/v2"
)

var Records = methods.GetRecords()

func AddR(c *fiber.Ctx) error {
	var Userdata structures.Data
	var mainStr [][]string
	err := c.BodyParser(&Userdata)
	if err != nil {
		fmt.Println(err)
		return errors.New("error occured File: findRestro.go, Method: FindR")
	}
	mainStr = methods.ConvertedData(Userdata)
	result := methods.AddData(mainStr)
	return c.JSON(result)
}

func FindR(c *fiber.Ctx) error {

	var filteredData []structures.TwoFieldData
	Cuisine := c.Query("cuisine")
	City := c.Query("city")
	HasTableBooking := c.Params("hasTableBooking")
	dataStructs := methods.GeneratingStructs(Records)
	filteredData = methods.StudyData(dataStructs, Cuisine, City, HasTableBooking)

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
