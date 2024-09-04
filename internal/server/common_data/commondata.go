package common_data

import (
	database "Tesis/database/sqlc"
	"context"
	"github.com/gin-gonic/gin"
)

func GetMunAndProv(id int64, store database.Store, ctx context.Context) (database.Municipio, database.Provincium, error) {
	mun, err := store.GetMunicipio(ctx, id)
	if err != nil {
		return database.Municipio{}, database.Provincium{}, err
	}
	prov, err := store.GetProvincia(ctx, mun.IDProvincia)
	if err != nil {
		return database.Municipio{}, database.Provincium{}, err
	}
	return mun, prov, nil
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
