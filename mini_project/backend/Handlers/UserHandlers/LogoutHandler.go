package UserHandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// fungsi untuk menangani permintaan logout pengguna.
func LogoutHandler(c *gin.Context) {
	c.SetCookie("jwtToken", "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Logout successfuly"})
}
