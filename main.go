package main

import (
  "GinApi/routes"
  "GinApi/database"
)

func main() {
  database.ConnectDB()
  routes.HanndlerFunc()
}
