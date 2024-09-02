package database

import (
	"context"
	"fmt"
	"testing"
)

func CreateMunicipio(t *testing.T) error {
	PinarDelRio := []string{"Pinar del rio", "Consolación del sur",
		"San Juan y Martínez", "Los Palacios", "Sandino", "Guane", "La Palma",
		"San Luis", "Minas de Matahambre", "Viñales", "Viñales"}
	LaHabana := []string{"Diez de Octubre", "Arroyo Naranjo", "Boyeros", "Playa", "Habana de el este",
		"San Migel del Padron", "Plaza de la Revolucion", "Centro Habana", "Marianao", "La Lisa", "Cerro",
		"Guanabacoa", "Habana Vieja", "Cotorro", "Regla"}
	Artemisa := []string{"Artemisa", "San Cristobal", "San Antonio de los Baños", "Bauta", "Bahía Honda",
		"Mariel", "Caimito", "Güira de Melena", "Alquízar", "Guanajay", "Candelaria"}
	Mayabeque := []string{"San José de las lajas ", "Güines", "Santa Cruz del Norte", "Madruga",
		"Quivicán", "Bejucal", "Batabanó", "Jaruco", "Nueva Paz", "San Nicolás de Bari", "Melena de el Sur"}

	Matanzas := []string{"Matanzas", "Cárdenas", "Colón", "Jagüey Grande", "Jovellanos", "Unión de los Reyes", "Pedro Betancourt",
		"Perico", "Calimete", "Limonar", "los Arabos", "Marti", "Ciénaga de Zapata"}
	VillaClara := []string{"Santa Clara", "Manicaragua", "Placetas", "Camajuani", "Ranchuelo", "Sagua la Grande", "Santo Domingo", "Remedios", "Caibarién",
		"Encrucijada", "Cifuentes", "Corralillo", "Quemado de Güines"}
	Cienfuegos := []string{"Cienfuegos", "Cumanayagua", "Rodas", "Palmira", "Aguada de Pasajeros", "Cruces", "Abreus", "Lajas"}
	SanctiSpiritus := []string{"Sancti Spíritus", "Trinidad", "Cabaiguán", "Yaguajay", "Jatibonico", "Taguasco", "Formento", "La Sierpe"}
	CiegoDeAvila := []string{"Ciego de Ávila", "Morón", "chambas", "Baraguá", "Ciro Redondo", "Primero de Enero", "Venezuela", "Majagua",
		"Florencia", "Bolivia"}
	Camaguey := []string{"Camagüey", "Florida", "Guáimaro", "Vertientes", "Santa Cruz del Sur", "Nuevitas", "Minas",
		"Sibanicú", "Esmeralda", "Carlos Manuel de Céspedes", "Jimaguayú", "Sierre de Cubitas", "Najasa"}

	LasTunas := []string{"Las Tunas", "Puerto Padre", "Jesús Menéndez", "Jobabo", "Majibacoa", "Amancio", "Colombia",
		"Manatí"}
	Holguin := []string{"Holguín", "Mayarí", "Banes", "Moa", "Gibara", "Calixto García", "Báguanos", "Rafael Freyre",
		"Sagua de Tánamo", "Urbano Noris", "Cacocúm", "Cueto", "Frank País", "Antilla"}
	Granma := []string{"Bayamo", "Manzanillo", "Jiguaní", "Yara", "Bartolomé Masó", "Guisa",
		"Río Cauto", "Campechuela", "Niquero", "Media Luna", "Buey Arriba", "Pilón", "Cauto Cristo"}
	Santiago := []string{
		"Santiago de Cuba", "Palma Soriano", "Contramaestre", "Songo - La Maya",
		"San Luis", "Segundo Frente", "Mella", "Guamá", "Tercer Frente",
	}
	Guantanamo := []string{
		"Guantánamo", "Baracoa", "El Salvador", "Maisí", "San Antonio del Sur",
		"Imaiás", "Yateras", "Niceto Pérez", "Manuel Tames", "Caimaneras",
	}
	Isla := []string{"Isla de la Juventud"}
	for _, m := range PinarDelRio {
		param := InsertMunicipioParams{Name: m, IDProvincia: 1}
		_, err := testStore.InsertMunicipio(context.Background(), param)
		if err != nil {
			fmt.Printf("error inserting the municipaly: %v", err)
		}
	}
	for _, m := range LaHabana {
		param := InsertMunicipioParams{Name: m, IDProvincia: 2}
		_, err := testStore.InsertMunicipio(context.Background(), param)
		if err != nil {
			fmt.Printf("error inserting the municipaly: %v", err)
		}
	}
	for _, m := range Artemisa {
		param := InsertMunicipioParams{Name: m, IDProvincia: 3}
		_, err := testStore.InsertMunicipio(context.Background(), param)
		if err != nil {
			fmt.Printf("error inserting the municipaly: %v", err)
		}
	}
	for _, m := range Mayabeque {
		param := InsertMunicipioParams{Name: m, IDProvincia: 4}
		_, err := testStore.InsertMunicipio(context.Background(), param)
		if err != nil {
			fmt.Printf("error inserting the municipaly: %v", err)
		}
	}
	for _, m := range Matanzas {
		param := InsertMunicipioParams{Name: m, IDProvincia: 5}
		_, err := testStore.InsertMunicipio(context.Background(), param)
		if err != nil {
			fmt.Printf("error inserting the municipaly: %v", err)
		}
	}
	for _, m := range VillaClara {
		param := InsertMunicipioParams{Name: m, IDProvincia: 6}
		_, err := testStore.InsertMunicipio(context.Background(), param)
		if err != nil {
			fmt.Printf("error inserting the municipaly: %v", err)
		}
	}
	for _, m := range Cienfuegos {
		param := InsertMunicipioParams{Name: m, IDProvincia: 7}
		_, err := testStore.InsertMunicipio(context.Background(), param)
		if err != nil {
			fmt.Printf("error inserting the municipaly: %v", err)
		}
	}
	for _, m := range SanctiSpiritus {
		param := InsertMunicipioParams{Name: m, IDProvincia: 8}
		_, err := testStore.InsertMunicipio(context.Background(), param)
		if err != nil {
			fmt.Printf("error inserting the municipaly: %v", err)
		}
	}
	for _, m := range CiegoDeAvila {
		param := InsertMunicipioParams{Name: m, IDProvincia: 9}
		_, err := testStore.InsertMunicipio(context.Background(), param)
		if err != nil {
			fmt.Printf("error inserting the municipaly: %v", err)
		}
	}
	for _, m := range Camaguey {
		param := InsertMunicipioParams{Name: m, IDProvincia: 10}
		_, err := testStore.InsertMunicipio(context.Background(), param)
		if err != nil {
			fmt.Printf("error inserting the municipaly: %v", err)
		}
	}
	for _, m := range LasTunas {
		param := InsertMunicipioParams{Name: m, IDProvincia: 11}
		_, err := testStore.InsertMunicipio(context.Background(), param)
		if err != nil {
			fmt.Printf("error inserting the municipaly: %v", err)
		}
	}
	for _, m := range Holguin {
		param := InsertMunicipioParams{Name: m, IDProvincia: 12}
		_, err := testStore.InsertMunicipio(context.Background(), param)
		if err != nil {
			fmt.Printf("error inserting the municipaly: %v", err)
		}
	}
	for _, m := range Granma {
		param := InsertMunicipioParams{Name: m, IDProvincia: 13}
		_, err := testStore.InsertMunicipio(context.Background(), param)
		if err != nil {
			fmt.Printf("error inserting the municipaly: %v", err)
		}
	}
	for _, m := range Santiago {
		param := InsertMunicipioParams{Name: m, IDProvincia: 14}
		_, err := testStore.InsertMunicipio(context.Background(), param)
		if err != nil {
			fmt.Printf("error inserting the municipaly: %v", err)
		}
	}
	for _, m := range Guantanamo {
		param := InsertMunicipioParams{Name: m, IDProvincia: 15}
		_, err := testStore.InsertMunicipio(context.Background(), param)
		if err != nil {
			fmt.Printf("error inserting the municipaly: %v", err)
		}
	}
	for _, m := range Isla {
		param := InsertMunicipioParams{Name: m, IDProvincia: 16}
		_, err := testStore.InsertMunicipio(context.Background(), param)
		if err != nil {
			fmt.Printf("error inserting the municipaly: %v", err)
		}
	}

	return nil
}

func TestCreateMun(T *testing.T) {
	_ = CreateMunicipio(T)

}
