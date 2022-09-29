package auth

import (
	"context"
	"net/http"
	"time"
	"todo/db"
	"todo/ent/user"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	ctx := context.Background()
	var login LoginDTO
	c.Bind(&login)

	// check if the user is valid
	u, err := db.DB.User.Query().Where(user.UsernameEQ(login.Username)).First(ctx)

	if err != nil {
		return c.JSON(http.StatusBadRequest, Respose{Errors: map[string]string{"username": "User not found"}})
	}

	if u.Password != login.Password {
		return c.JSON(http.StatusBadRequest, Respose{Errors: map[string]string{"password": "Password don't match"}})

	}

	claims := Claims{
		UserID:   u.ID,
		Username: u.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Respose{Errors: "Token cannot be signed"})
	}

	return c.JSON(http.StatusOK, Respose{Data: map[string]interface{}{
		"token": t,
		"user":  u,
	}})
}
