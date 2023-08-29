package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/godror/godror"
)

type Item struct {
	itemNumber string
	itemPrice  string
}

func main() {
	connoteNumbers := ReadFile("./file.txt")

	//Connect to DB
	db, err := sql.Open("godror", `user="username" password="password" connectString="x.x.x.x/x" noTimezoneCheck=true`)
	if err != nil {
		fmt.Println("Failed to connect DB:", err)
		return
	}

	defer db.Close()

	//Interact with DB
	for i := 0; i < len(connoteNumbers); i++ {
		query := fmt.Sprintf("UPDATE VT_PPS_ACCEPTANCEDTL SET CODGOODSPRC = '%s' where itemno = '%s'\n",
			connoteNumbers[i].itemPrice, connoteNumbers[i].itemNumber)
		//Querying data
		rows, err := db.Query(query)
		if err != nil {
			fmt.Println("Failed to query data:", err)
			return
		}
		defer rows.Close()
	}

}

func ReadFile(fileLocation string) []Item {
	file, err := os.Open(fileLocation)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return nil
	}

	defer file.Close() // Close the file at the end

	items := []Item{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		connote := strings.Split(line, ",")
		items = append(items, Item{itemNumber: strings.ToUpper(connote[0]), itemPrice: connote[1]})
	}
	return items
}
