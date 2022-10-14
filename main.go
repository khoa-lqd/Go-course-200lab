package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Note struct {
	Id      int    `json:"id,omitempty" gorm:"column:id;"`
	Title   string `json:"title" gorm:"column:title;"`
	Content string `json:"content" gorm:"column:content;"`
}

// set Note is the name of the table (notes)
func (Note) TableName() string {
	return "notes"
}

func main() {
	os.Setenv("DBConnectionStr", "food_delivery:19e5a718a54a9fe0559dfbce6908@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local")
	dsn := os.Getenv("DBConnectionStr")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// fmt.Println(db, err)
	if err != nil {
		log.Fatalln(err)
	}
	//insert new note
	// newNote := Note{Title: "Demo Note 3", Content: "This is a demo note 3"}
	// db.Create(&newNote)

	// fmt.Println(newNote)

	// get first note id from db
	var note Note
	db.First(&note)
	fmt.Println(note.Id)

	//find note where status = 1
	var result []Note
	db.Where("status = ?", 1).Find(&result)
	fmt.Println(result)

	//get note where primarykey(id = 3) from db
	db.First(&result, 3)
	fmt.Println(result)

	// get note where title = planning
	var note2 Note
	db.Where("title = ?", "planning").First(&note2)
	fmt.Println(note2)

	// get note where title = "demo note" and content ="this is a demo note"
	var note3 Note
	db.Where("title = ? AND content= ?", "demo note", "this is a demo note").Find(&note3)
	fmt.Println(note3)

	//-----------------------------------------------------------------------------------------//

	// select
	var note4 Note
	db.Select("Id", "Title").Find(&note4)
	fmt.Println(note4)
}
