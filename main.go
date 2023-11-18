package main

import (
	"fmt"

	"github.com/EcoPowerHub/EMS/config"
	"github.com/EcoPowerHub/EMS/utils"
)

func main() {
	fmt.Println("Starting app")
	var equipmentData config.Equipments
	err := utils.ReadJsonFile("tests/equipment.json", &equipmentData)
	if err != nil {
		return
	}
	equipmentData.Print()
}
