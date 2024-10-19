package userhandler

import (
	"net/http"

	usermodel "github.com/PhuPhuoc/koi-story-api/controller/user_services/model" // Import usermodel để sử dụng struct Register
	userrepository "github.com/PhuPhuoc/koi-story-api/controller/user_services/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath	/api/v1

// RegisterNewAccount godoc
//
//	@Summary		Register new user account
//	@Description	Register new user account
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		usermodel.Register		true	"User info for registering a new account"
//	@Success		201 	{object}	map[string]interface{}	"message success"
//	@Failure		400		{object}	error					"Bad request error"
//	@Router			/users/register [post]
func registerNewAccountHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req usermodel.Register

		// Bind JSON data to usermodel.Register struct
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// Create a new instance of the repository
		repo := userrepository.NewUserStore(db)

		// Pass the request data to the repository to create a new account
		err := repo.RegisterNewAccount(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Account created successfully"})
	}
}
