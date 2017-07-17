package SDRouter

import (
	"github.com/buaazp/fasthttprouter"
	FCAdminHandler "github.com/snowheat/falconism/handlers/admin"
)

//Set ...
func Set(router *fasthttprouter.Router) {
	router.GET("/admin", FCAdminHandler.New)
	router.POST("/admin/post", FCAdminHandler.Post)
}
