package main


func main(){
	r :=gin.Default()
	r.User(middlerware())

	handler.RegisterUserHandler(r)

}


func middlerware() gin.HandlerFunc{
	return func (c *gin.Context){
		path := c.Request.URL.Path
		fmt.Println("the path is " + path)
		c.Set("key","value")
		// c.Set可以赋予对象，c.Get 需要进行强制转换

		c.Next()
		//next 之后就开始执行真正的路由函数，写在next后面的代码是路由函数返回后的
		//注意中间件的执行顺序类似于python 装饰器：
		// A1 B1 B2 A2
	}
}