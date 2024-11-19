package datos_sociales

import (
	"TesisAclifim/internal/server/common_data"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetDatosRequest struct {
	Asociado_id int64 `uri:"asociado_id" binding:"required"`
}

func (s *Server) GetDatosS(ctx *gin.Context) {
	var req GetDatosRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, common_data.ErrorResponse(err))
		return
	}
	datos, err := s.GetStore().GetDatosSociales(ctx, req.Asociado_id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, common_data.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, common_data.ErrorResponse(err))
		return
	}
	ctx.JSON(200, datos)
}
