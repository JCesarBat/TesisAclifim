package users

import (
	database "TesisAclifim/database/sqlc"
	"TesisAclifim/internal/server/common_data"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListUserRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}
type ListUserResponse struct {
	Users []GetUserResponse `form:"users" binding:"required"`
}

// @BasePath /login
// PingExample godoc
// @Summary 	get many users
// @Schemes
// @Description  obtienes varios usuarios en un rango dado por el pageID and pageSize
// @Tags users
// @Accept json
// @Produce json
// @Param listUser body ListUserRequest true "list user  request"
// @Router /users [get]
// Return a list of users
func (s *Server) ListUser(ctx *gin.Context) {
	var req ListUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, common_data.ErrorResponse(err))
		return
	}

	arg := database.ListUsersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	users, err := s.GetStore().ListUsers(ctx, arg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, common_data.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, common_data.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, users)
}
