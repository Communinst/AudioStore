package handler

import (
	"AudioShare/backend/internal/entity"
	"AudioShare/backend/internal/service"
	"AudioShare/backend/internal/validation"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const (
	default_user_role       = "user"
	default_expiration_time = 24
)

type AuthHandler struct {
	auth service.AuthServiceInterface
}

func NewAuthHandler(srvc service.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{
		auth: srvc,
	}
}

// @Summary User registration
// @Description Register a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param request body entity.User true "User registration data"
// @Success 200 {object} map[string]interface{} "User created successfully"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /auth/signup [post]
func (this *AuthHandler) SignUp(c *gin.Context) {
	slog.Info("auth handler: sign up: initiated")
	var userData entity.User

	if err := c.ShouldBindBodyWithJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := validation.IsPasswordValid(userData.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid password"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	userData.Password = string(hashedPassword)
	userData.Registered = time.Now()

	defaultRoleString := os.Getenv("DEFAULT_ROLE") //Add obtaining default via db/redis by default_user_role
	defaultRole, err := strconv.ParseUint(defaultRoleString, 10, 8)
	if err != nil {
		slog.Info("Failed to convert .env data. Role id value is 1")
		defaultRole = 1
	}
	userData.RoleId = uint8(defaultRole)

	resultId, err := this.auth.PostOne(c.Request.Context(), &userData)
	if err == nil { //PostOne guarantees valid id since there's no error occured
		slog.Info("auth handler: sign up: succeeded")
		c.JSON(http.StatusOK, gin.H{
			"addedId": resultId,
			"message": "User created successfully"})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	slog.Info("auth handler: sign up: finished")
}

// @Summary User login
// @Description Authenticate user and return tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param request body object{email=string,password=string} true "User login credentials"
// @Success 200 {object} map[string]interface{} "Login successful"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 401 {object} map[string]string "Invalid credentials"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /auth/signin [post]
func (this *AuthHandler) SignIn(c *gin.Context) {
	slog.Info("auth handler: sign in: initiated")
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Retrieve the user by email
	user, err := this.auth.GetOneByEmailFull(c.Request.Context(), creds.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Verify the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate an authentication token
	token := os.Getenv("AUTHORIZATION_TOKEN_SECRET")
	if token == "" {
		log.Fatal("ACCESS_TOKEN_SECRET environment variable not set")
	}
	expiry, err := strconv.ParseInt(os.Getenv("AUTHORIZATION_EXPIRE_TIME"), 10, 64) //TODO: move from here

	if err != nil {
		log.Fatal("AUTHORIZATION_EXPIRE_TIME environment variable parse failed")
	}
	if expiry == 0 {
		log.Fatal("AUTHORIZATION_EXPIRE_TIME environment variable not set")
	}
	userToken, err := this.auth.GenerateAuthToken(user, token, int(expiry))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Replace env dependecy with request to cache
	// in order to refine feedback on roles
	defaultAdmin, _ := strconv.ParseInt(os.Getenv("DEFAULT_ADMIN_ROLE_ID"), 10, 64)
	if user.RoleId == uint8(defaultAdmin) {
		//c.Redirect(http.StatusOK, "/admin")
		c.JSON(http.StatusOK, gin.H{"token": userToken, "role": "admin", "userId": user.Id})
	} else {
		//c.Redirect(http.StatusOK, "/api")
		c.JSON(http.StatusOK, gin.H{"token": userToken, "role": "user", "userId": user.Id})
	}
}
