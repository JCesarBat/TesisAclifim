package database

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDinamicList(t *testing.T) {

	param := DinamicListParam{Nombre: "Juan", Ocupacion: "asdfw"}
	res, err := testStore.DinamicList(context.Background(), param)
	require.NoError(t, err)
	fmt.Println(res)
}
