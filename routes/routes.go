package routes

import(
  "GinApi/controllers"
  "github.com/gin-gonic/gin"
)

func HanndlerFunc(){
  app := gin.Default()
  app.GET("/:name", controllers.Salut)
  app.GET("/students", controllers.GetStudents)
  app.POST("/students", controllers.CreateStudent)
  app.GET("/students/:id", controllers.GetStudentById)
  app.DELETE("/students/:id", controllers.DeleteStudent)
  app.PUT("/students/:id", controllers.EditStudent)
  app.GET("/students/cpf/:cpf", controllers.SearchStudentByCpf)
  app.Run()
}
