package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

/*
"StartApplication function of user app"
*/
func StartApplicatin() {
	mapUrls()
	router.Run(":8080")
}
