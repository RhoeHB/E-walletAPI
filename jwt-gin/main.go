package main

import (
  	"github.com/gin-gonic/gin"
	"ottoTest/controllers"
	"ottoTest/models"
	"ottoTest/middlewares"
)

func main() {

	models.ConnectDataBase()
	
	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login",controllers.Login)

	protected := r.Group("/api/wallet")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user",controllers.CurrentUser)
	protected.GET("/balance",controllers.GetUserBalance)
	protected.GET("/transactions",controllers.GetUserTransactions)
	protected.GET("/biller-transactions",controllers.GetBillerTransactions)
	protected.GET("/confirm-biller-transaction",controllers.ConfirmBillerTransaction)

	protected.POST("/top-up",controllers.TopUp)
	protected.POST("/delete-user",controllers.DeleteUser)
	protected.POST("/reset-password",controllers.ChangeUserPassword)


	r.Run(":8080")

}
// pls hire me
// herohartanto@gmail.com