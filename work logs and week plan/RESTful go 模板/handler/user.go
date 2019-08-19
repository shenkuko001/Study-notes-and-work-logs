var userHandler *gin.RouterGroup

import ...

func RegisterUserHandler(router *gin.Engine){
	userHandler=router.Group("/user")
	userHandler.POST("",addUser)
}

fucn addUser(c *gin.Context){
	user := new(model.User)
	if err:=c.BindJSON(user);err!=nil{
		logger.Error(err)
		handler.ErrorHandler(c,err)
		return 
	}
	if err := service.InserUser(user);err!=nil{
		logger.Error(err)
		handler.ErrorHandler(c,err)
		return
	}
	c.JSON(200,gin.H{
		"data":"ok"
	})

}