package handler

import (
	_ "AudioShare/backend/internal/entity"
	"AudioShare/backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	AdminRoleId = 1
)

type UserHandler struct {
	service service.UserServiceInterface
}

func NewUserHandler(srv service.UserServiceInterface) *UserHandler {
	return &UserHandler{service: srv}
}

// @Summary Get user by ID
// @Description Get user profile information by user ID
// @Tags users
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 200 {object} entity.User "User profile"
// @Failure 400 {object} map[string]string "Invalid user ID"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "User not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /users/{id} [get]
func (h *UserHandler) ObtainProfileById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	user, err := h.service.GetOneById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Summary Get all users
// @Description Get a list of all users (admin only)
// @Tags users
// @Produce json
// @Security BearerAuth
// @Success 200 {array} entity.User "List of users"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 403 {object} map[string]string "Forbidden"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /users/ [get]
func (h *UserHandler) ObtainAllUsers(c *gin.Context) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authorized"})
		return
	}
	userID, ok := userIDVal.(uint64)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	user, err := h.service.GetOneById(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}
	if user.RoleId != AdminRoleId {
		c.JSON(http.StatusForbidden, gin.H{"error": "admin access required"})
		return
	}
	users, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// @Summary Delete user by ID
// @Description Delete a user by ID (admin only)
// @Tags users
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 204 "User deleted successfully"
// @Failure 400 {object} map[string]string "Invalid user ID"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 403 {object} map[string]string "Forbidden"
// @Failure 404 {object} map[string]string "User not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /users/{id} [delete]
func (h *UserHandler) RemoveUserById(c *gin.Context) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authorized"})
		return
	}
	userID, ok := userIDVal.(uint64)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	user, err := h.service.GetOneById(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}
	if user.RoleId != AdminRoleId {
		c.JSON(http.StatusForbidden, gin.H{"error": "admin access required"})
		return
	}
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	err = h.service.DeleteOneById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "user deleted"})
}
