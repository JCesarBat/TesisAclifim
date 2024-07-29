package auth

import (
	database "Tesis/database/sqlc"
	"Tesis/pkg/util"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterRequest struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Password2 string `json:"password2" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Provincia int    `json:"provincia" binding:"required"`
	Municipio int    `json:"municipio" binding:"required"`
}
type RegisterResponse struct {
	Id        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Provincia string `json:"provincia"`
	Municipio string `json:"municip"`
}

// @BasePath /register
// PingExample godoc
// @Summary 	register example
// @Schemes
// @Description  this is the documentation abaut register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param register body RegisterRequest true "register request"
// @Success 200 {object} RegisterResponse
// @Router /register [post]
func (s *Server) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Password != req.Password2 {
		c.JSON(http.StatusConflict, gin.H{"error": "the password dont match's"})
	}
	hashPassword, err := util.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	params := database.InsertUserParams{
		Username:    req.Username,
		Password:    hashPassword,
		Email:       req.Email,
		IDMunicipio: int64(req.Municipio),
	}
	user, err := s.store.InsertUser(c, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error "})
		return
	}
	mun, err := s.store.GetMunicipio(c, int64(req.Municipio))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "the municipio selected is not in the database"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error "})
		return
	}
	prov, err := s.store.GetProvincia(c, int64(req.Municipio))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "the provincia selected is not in the database"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error "})
		return
	}

	response := RegisterResponse{
		Id:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Provincia: prov.Name,
		Municipio: mun.Name,
	}
	c.JSON(http.StatusOK, response)
}
