package findRestro

import (
	"errors"
	"main/methods"
	"main/structures"

	"github.com/gofiber/fiber/v2"
)

func FindR(c *fiber.Ctx) error {
	Userdata := new(structures.Choices)
	var filteredData []structures.TwoFieldData
	err := c.BodyParser(Userdata)
	if err != nil {
		return errors.New("error occured File: findRestro.go, Method: FindR")
	}

	Records := methods.GetRecords()
	dataStructs := methods.GeneratingStructs(Records)

	filteredData = methods.StudyData(dataStructs, Userdata.Cuisine, Userdata.City, Userdata.HasTableBooking)

	return c.JSON(filteredData)

}

func FindForm(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}
