package go_model
import (
	"database/sql"

	_ "github.com/lib/pq"
	"log"
	"fmt"
)

func PgConnction() {
	connStr := "user=chenjian dbname=cable_0316 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM users")

	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("===== runing ===")
	fmt.Println(rows)
}