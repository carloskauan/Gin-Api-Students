package database

import(
  "log"
  "GinApi/models"
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
)

var(
  DB *gorm.DB
  err error
)

func ConnectDB(){
  connector := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
  DB, err = gorm.Open(postgres.Open(connector))
  if err != nil{
    log.Panic("Erro ao connectar ao banco de dados. ",err)
  }
  DB.AutoMigrate(&models.Student{})
}
