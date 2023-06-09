// package routes

// import (
// 	"github.com/gin-gonic/gin"
// 	"github.com/itsmaheshkariya/gin-gorm-rest/controller"
// )

// func UserRoute(router *gin.Engine){
// 	router.GET("/",controller.GetUsers)
// 	router.POST("/",controller.CreateUsers)
// 	router.DELETE("/:id",controller.DeleteUsers)
// 	router.PUT("/:id", controller.UpdateUsers)

// }
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/itsmaheshkariya/gin-gorm-rest/controller"
)

func UserRoute(router *gin.Engine) {
	// User routes
	router.GET("/users", controller.GetUsers)
	router.POST("/users", controller.CreateUsers)
	router.DELETE("/users/:id", controller.DeleteUsers)
	router.PUT("/users/:id", controller.UpdateUsers)

	// Laptop routes
	router.GET("/laptops", controller.GetLapUsers)
	router.GET("/laptops/:id", controller.GetLapUserByID)
	router.POST("/laptops", controller.CreateLapUsers)
	router.DELETE("/laptops/:id", controller.DeleteLapUsers)
	router.PUT("/laptops/:id", controller.UpdateLapUsers)
}
