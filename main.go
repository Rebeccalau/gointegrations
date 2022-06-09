package main

import (
	"databases/awsintegration"
	"databases/cmdexecutions"
	"databases/nosql"
	sql2 "databases/sql"
	"fmt"
)

func main() {
	//SqlDatabaseCalls()
	//NoSqlDatabaseCalls()
	//CommandExecutions()
	AWSCalls()

}
func AWSCalls() {
	awsintegration.Upload()
}

func CommandExecutions() {
	cmdexecutions.ExecuteCLICommand()
	cmdexecutions.ExecuteShellScript()
}

func NoSqlDatabaseCalls() {
	database := nosql.Connect()
	database.NewCollection()
	defer database.Close()

	database.InsertDoc()
	result := database.FindDoc()
	fmt.Println(result)

	result1 := database.FindAll()
	fmt.Println(result1)

	database.DeleteDocument()
}

func SqlDatabaseCalls() {
	database := sql2.NewSqlDatabase()

	database.CreateTable()

	database.InsertRow()
	result, err := database.QueryTable()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	database.DeleteRow()
	result1, err1 := database.QueryTable()
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(result1)
	}

	//database.AddTableColumn()
	//database.DropTableColumn()
}
