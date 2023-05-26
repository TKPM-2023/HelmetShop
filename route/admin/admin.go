package admin

import (
	"LearnGo/component/appctx"
	"LearnGo/middleware"
	"LearnGo/module/category/categorytransport/gincategory"
	"LearnGo/module/upload/uploadtransport/ginupload"
	"LearnGo/module/user/usertransport/ginuser"
	"github.com/gin-gonic/gin"
)

func AdminRoute(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	admin := v1.Group("/admin",
		middleware.RequireAuth(appContext),
		middleware.RoleChecker(appContext,
			"admin"))

	admin.GET("/profile", ginuser.GetProfile(appContext))
	admin.DELETE("/remove/:id", ginupload.Remove(appContext))

	// category
	category := admin.Group("/categories")
	category.POST("/", gincategory.CreateCategory(appContext))
	category.GET("/:id", gincategory.GetCategory(appContext))
}
