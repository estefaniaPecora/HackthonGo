package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/estefaniaPecora/HackthonGo/internal/customers"
	"github.com/estefaniaPecora/HackthonGo/internal/invoices"
	"github.com/gin-gonic/gin"
)

var repo = customers.NewRepository()
var service = customers.NewService(repo)

var invoice_repo = invoices.NewRepository()
var invoice_service = invoices.NewService(invoice_repo)

func SaveOne(c *gin.Context) {

	new, err := service.SaveCustomer("new_Lastname", "newFirst_name", "active")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Success %v", new)
	c.JSON(200, gin.H{"status": "OK"})
}

func Load(c *gin.Context) {

	itemName := c.Param("filename")
	currentFile := "../datos/" + itemName + ".txt"
	file, err := os.Open(currentFile)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer file.Close()

	readItems := bufio.NewScanner(file)

	for readItems.Scan() {
		line := readItems.Text()
		items := strings.Split(line, "#$%#")

		switch itemName {
		case "customers":
			_, err := service.SaveCustomer(items[1], items[2], items[3])
			if err != nil {
				log.Fatal(err)
			}
		case "invoices":
			date_time := items[1]
			total := 0.0
			id_customer, _ := strconv.Atoi(items[2])

			_, err := invoice_service.SaveInvoices(date_time, total, id_customer)
			if err != nil {
				log.Fatal(err)
			}
		}

	}

	c.JSON(200, gin.H{"status": "All Items Loaded"})
}

func main() {

	//endpoint loadData: abrir archivo de texto, recorrer el archivo y crear un customer, por cada customer llamo al save y lo guardo en la DB

	router := gin.Default()

	router.POST("/loaditems/:filename", Load)
	router.POST("/customers/saveone", SaveOne)

	router.Run()

}
