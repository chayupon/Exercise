package handles

import (
	"database/sql"
	
	"github.com/chayupon/Exercise/Sale/service/v1/register"
	"github.com/gin-gonic/gin"
)
// Router app
func Router(db *sql.DB) {
	regist :=register.Input{
		Db : db,
	}
	r := gin.Default()
	r.POST("/register/email", regist.Email)
	r.GET("/member", regist.ListAccounts)
	r.Run()
}
