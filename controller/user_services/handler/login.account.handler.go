package userhandler

import (
	"net/http"

	usermodel "github.com/PhuPhuoc/koi-story-api/controller/user_services/model"
	userrepository "github.com/PhuPhuoc/koi-story-api/controller/user_services/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

//	@BasePath	/api/v1

// Login godoc
//
//	@Summary		User login
//	@Description	User login
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		usermodel.UserLoginInfo	true	"User log in info"
//	@Success		200		{object}	map[string]interface{}	"user data"
//	@Failure		400		{object}	error					"Bad request error"
//	@Router			/users/login [post]
func loginHandler(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req usermodel.UserLoginInfo

		// Bind JSON data to usermodel.Register struct
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// Create a new instance of the repository
		repo := userrepository.NewUserStore(db)

		// Pass the request data to the repository to create a new account
		user, err := repo.Login(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "login successfully", "data": user})
	}
}
