package database

import (
	"context"
	"log"
	"testing"
)

func CreateProvincias(t *testing.T) {
	Prov := []string{"Pinar de el Río", "La Habana", "Artemisa",
		"Mayabeque", "Matanzas", "Villa Clara", "Cienfuegos",
		"Santi Spiritus", "Ciego de Avila", "Camagüey", "Las Tunas",
		"Holguín", "Granma", "Santiago de Cuba", "Guantánamo", "Isla de la Juventud"}

	for _, v := range Prov {
		_, err := testStore.InertarProv(context.Background(), v)
		if err != nil {
			log.Printf("error ocurred possibly the provincia :%s exists ,%v", v, err)
		}
	}
}

func TestCreateProvincias(t *testing.T) {
	CreateProvincias(t)
}
