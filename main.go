package main

import (
	db2 "back-end1-mini-project/db"
	"back-end1-mini-project/modules/customer"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	configdb, err := db2.LoadConfig()
	if err != nil {
		log.Printf("Config load error : %s", err)
	}

	db, err := db2.ConnectDB(configdb)

	if err != nil {
		fmt.Println("Failed to connect database, error : ", err)
		return
	}

	customerRouter := customer.NewCustomerRouter(db)
	customerRouter.Handle(router)

	err = router.Run(":8081")
	if err != nil {
		fmt.Println("Error running server:", err)
		return
	}
}
