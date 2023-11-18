package config

import "fmt"

type Equipment struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Host        string `json:"host"`
	Name        string `json:"name"`
}

type Equipments struct {
	Equipments []Equipment `json:"equipment"`
}

func (e *Equipments) Print() {
	for _, equipment := range e.Equipments {
		fmt.Printf("ID: %s\n", equipment.ID)
		fmt.Printf("Description: %s\n", equipment.Description)
		fmt.Printf("Host: %s\n", equipment.Host)
		fmt.Printf("Name: %s\n", equipment.Name)
		fmt.Println()
	}
}
