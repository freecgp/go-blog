package models

import(
	"fmt"
	"github.com/jinzhu/gorm"
	 _ "github.com/jinzhu/gorm/dialects/postgres"
)

var(
	db *gorm.DB
)
func init() {
	var err error
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=cgp dbname=cgp password=cgp sslmode=disable")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
}
