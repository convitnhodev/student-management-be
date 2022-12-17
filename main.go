package main

import (
	"fmt"
	"managerstudent/common/setupDatabase"
)

func main() {
	db := setupDatabase.InitMongoDB()
	fmt.Println(db)

}
