package controllers

import (
  "net/http"
  "GinApi/database"
  "GinApi/models"
  "github.com/gin-gonic/gin"
)

func Salut(c *gin.Context){
  name := c.Params.ByName("name")
  c.JSON(200, gin.H{
    "Api diz":"E ai "+name+" Tudo beleza?",
  })
}

func GetStudents(c *gin.Context){
  var students []models.Student
  database.DB.Find(&students)
  c.JSON(http.StatusOK, students)
}

func GetStudentById(c *gin.Context){
  var student models.Student
  id := c.Params.ByName("id")
  database.DB.First(&student, id)
  if student.ID == 0{
    c.JSON(http.StatusNotFound, gin.H{
      "Not found": "Aluno não encontrado",
    })
    return
  }
  c.JSON(200, student)
}

func CreateStudent(c *gin.Context){
  var aluno models.Student
  if err := c.ShouldBindJSON(&aluno); err != nil{
    c.JSON(http.StatusBadRequest, gin.H{
      "error":err.Error(),
    })
    return
  }
  database.DB.Create(&aluno)
  c.JSON(http.StatusOK, aluno)
}

func DeleteStudent(c *gin.Context){
  var student models.Student
  id := c.Params.ByName("id")
  database.DB.Delete(&student, id)
  c.JSON(http.StatusOK, gin.H{
    "data":"Aluno deletado com sucesso",
  })
}

func EditStudent(c *gin.Context){
  var student models.Student
  id := c.Params.ByName("id")
  database.DB.First(&student, id)
  if err := c.ShouldBindJSON(&student); err != nil{
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }
  database.DB.Model(&student).UpdateColumns(student)
  c.JSON(http.StatusOK, student)
}

func SearchStudentByCpf(c *gin.Context){
  var student models.Student
  cpf := c.Param("cpf")
  database.DB.Where(&models.Student{CPF: cpf}).First(&student)
  if student.ID == 0{
    c.JSON(http.StatusBadRequest,gin.H{
      "Not Found":"Aluno não econtardo",
    })
    return
  }
  c.JSON(http.StatusOK, student)
}
