package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// GetAll GetAllUsers godoc
// @Summary Get all users
// @Description Fetch all users
// @Tags Users List
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} Response
// @Router /api/users [get]
func (h *Handler) GetAll(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

// Create CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags Users Create
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "Create User Request"
// @Success 201 {object} Response
// @Router /api/users/create [post]
func (h *Handler) Create(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.service.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Login Login godoc
// @Summary User login
// @Description Authenticate user and return JWT token
// @Tags Users Auth
// @Accept json
// @Produce json
// @Param credentials body LoginRequest true "Login Request"
// @Success 200 {object} TokenResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/auth/login [post]
func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := h.service.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	c.JSON(http.StatusOK, TokenResponse{Token: token})
}
