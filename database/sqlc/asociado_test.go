package database

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateAsociado(t *testing.T) {
	param := InsertAsoiciadoParams{
		Nombre:      "Pdro",
		Apellido1:   "Buitista",
		Apellido2:   "carlos",
		Activo:      true,
		Sexo:        true,
		NumeroT:     sql.NullInt64{Int64: 1231432531, Valid: true},
		Direccion:   "chae",
		IDMunicipio: 4,
	}
	_, err := testStore.InsertAsoiciado(context.Background(), param)
	require.NoError(t, err)
}
