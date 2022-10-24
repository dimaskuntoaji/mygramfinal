package server

import (
	"finall/controller/controllercomment"
	"finall/controller/controllerphoto"
	"finall/controller/controllersocialmedia"
	"finall/controller/controlleruser"
	"finall/middleware"
	"finall/repository/repositorycomment"
	"finall/repository/repositoryphoto"
	"finall/repository/repositorysocialmedia"
	"finall/repository/repositoryuser"
	"finall/service/servicecomment"
	"finall/service/servicephoto"
	"finall/service/servicesocialmedia"
	"finall/service/serviceuser"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(r *gin.Engine, db *gorm.DB) {

	repoUser := repositoryuser.New(db)
	srvUser := serviceuser.New(repoUser)
	ctrlUser := controlleruser.New(srvUser)

	routeUser := r.Group("/users")

	// route user
	routeUser.POST("/register", ctrlUser.Create)
	routeUser.POST("/login", ctrlUser.Login)
	routeUser.PUT("", middleware.Authorization, ctrlUser.Update)
	routeUser.DELETE("", middleware.Authorization, ctrlUser.DeleteByID)

	// route photos
	repoPhoto := repositoryphoto.New(db)
	srvPhoto := servicephoto.New(repoPhoto)
	ctrlPhoto := controllerphoto.New(srvPhoto)

	r.GET("photos", middleware.Authorization, ctrlPhoto.GetPhotos)
	r.POST("photos", middleware.Authorization, ctrlPhoto.Create)
	r.PUT("photos/:photoID", middleware.Authorization, ctrlPhoto.Update)
	r.DELETE("photos/:photoID", middleware.Authorization, ctrlPhoto.Delete)

	// route comment
	repoComment := repositorycomment.New(db)
	srvComment := servicecomment.New(repoComment)
	ctrlComment := controllercomment.New(srvComment)
	r.GET("comments", middleware.Authorization, ctrlComment.Get)
	r.POST("comments", middleware.Authorization, ctrlComment.Create)
	r.PUT("comments/:commentID", middleware.Authorization, ctrlComment.Update)
	r.DELETE("comments/:commentID", middleware.Authorization, ctrlComment.Delete)

	// route social media
	repoSocialmedia := repositorysocialmedia.New(db)
	srvSocialmedia := servicesocialmedia.New(repoSocialmedia, repoPhoto)
	ctrlSocialmedia := controllersocialmedia.New(srvSocialmedia)
	routerSocialmedia := r.Group("/socialmedias")
	routerSocialmedia.POST("", middleware.Authorization, ctrlSocialmedia.Create)
	routerSocialmedia.GET("", middleware.Authorization, ctrlSocialmedia.GetList)
	routerSocialmedia.PUT("/:socialmediaid", middleware.Authorization, ctrlSocialmedia.UpdateByID)
	routerSocialmedia.DELETE("/:socialmediaid", middleware.Authorization, ctrlSocialmedia.DeleteByID)
}
