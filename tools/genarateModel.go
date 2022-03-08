package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"leaf-go/utils"
	"strings"
)

type Fields []*Field
type FieldString string

func (f FieldString) Name() string {
	s := string(f)
	if s == "id" {
		return "ID"
	}

	return toTitle(s)
}

type TypeString string

func (t TypeString) GetModelType() string {
	s := string(t)
	if -1 < strings.Index(s, "int") {
		return "int"
	}

	if _, ok := utils.Regexp.Match(".*(char|text|json|blob).*", s); ok {
		return "string"
	}

	if _, ok := utils.Regexp.Match(".*(time|date).*", s); ok {
		return "time.Time"
	}

	return ""
}

type Field struct {
	Field FieldString `gorm:"Field"`
	Type  TypeString  `gorm:"Type"`
}

func toTitle(s string) string {
	ss := strings.Split(s, "_")
	name := ""
	for _, s1 := range ss {
		name += strings.Title(s1)
	}

	return name
}

func tableInfo(table string) Fields {
	fields := make(Fields, 0)
	db.Raw(fmt.Sprintf("DESC %s", table)).Scan(&fields)
	return fields
}

func getTableList() []string {
	var s []string
	db.Raw("show tables").Scan(&s)
	return s
}

var (
	db = withGorm()
)

func main() {
	//db.
	tables := getTableList()
	for _, table := range tables {
		fields := tableInfo(table)
		s := template(table, fields)
		tableTitle := toTitle(table)

		filePath := utils.File.AppPath() + "/data/models/" + tableTitle + ".go"
		if exists, err := utils.File.Exists(filePath); err != nil || exists {
			if exists {
				//utils.Output.Errorf("%s exists", filePath)
				continue
			}

			if err != nil {
				utils.Output.Errorf("check %s exists err: %+v", filePath, err)
				continue
			}
		}

		if err := utils.File.Create(filePath, []byte(s)); err != nil {
			utils.Output.Errorf("create %s failed err: %+v", filePath, err)
			continue
		}

		utils.Output.Successf("create %s success ,path: %+v\n", tableTitle, filePath)
	}

}

func withGorm() *gorm.DB {
	var user, password, host, port, dbname string
	user = "root"
	password = "123123"
	host = "127.0.0.1"
	port = "3306"
	dbname = "stu"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local", user, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(
			fmt.Sprintf("couldn't open database: %v\n", err),
		)
	}
	return db
}

func template(table string, fields Fields) string {
	tableTitle := toTitle(table)

	fw := string(table[0])

	fieldString := ""
	formatted := "\t%s\t%s\t`json:\"%s\" gorm:\"column:%s\"`\n"
	rows, cols := "", ""

	for _, field := range fields {
		t := field.Type.GetModelType()
		name := field.Field.Name()

		if name == "Status" {
			t = "bool"
		}

		fieldString += fmt.Sprintf(
			formatted,
			name,
			t,
			field.Field,
			field.Field,
		)

		s1 := field.Field.Name()
		param := strings.ToLower(string(s1[0])) + s1[1:]

		if s1 == "CreatedAt" || s1 == "UpdatedAt" || s1 == "ID" {
			continue
		}

		rows += param + " " + t + ","
		cols += fw + "." + name + " = " + param + "\n"
	}

	rows = strings.TrimRight(rows, ",")

	//
	temp := `
	package models
	
	import (
		"gorm.io/gorm"
		"time"
	)
	

	func New` + tableTitle + `() *` + tableTitle + ` {
		return &` + tableTitle + `{}
	}
	
	type ` + tableTitle + `s []*` + tableTitle + `
	func (` + fw + ` ` + tableTitle + `s) Ids() []int {
		ids := make([]int, len(` + fw + `))
		for i, v := range ` + fw + ` {
			ids[i] = v.ID
		}

		return ids
	}	

	type ` + tableTitle + ` struct {
		Model
	` + fieldString + `
	}
	
	func (` + fw + ` *` + tableTitle + `) TableName() string {
		return "` + table + `"
	}
	
	func (` + fw + ` *` + tableTitle + `) Table() *gorm.DB {
		return db.Table(` + fw + `.TableName())
	}
	
	func (` + fw + ` *` + tableTitle + `) IsEmpty() bool {
		return ` + fw + `.ID == 0
	}
	
	func (` + fw + ` *` + tableTitle + `) IdGet(id int) error {
		return db.Where("id = ? AND status = 1", id).First(&` + fw + `).Error
	}

	func (` + fw + ` *` + tableTitle + `) Create(` + rows + `) error {
		` + cols + `	
		return db.Create(&` + fw + `).Error
	}

	func (` + fw + ` *` + tableTitle + `) Edit(` + rows + `) error {
		` + cols + `	
		return db.Save(&` + fw + `).Error
	}

	func (` + fw + ` *` + tableTitle + `) Save() error {
		return db.Save(&` + fw + `).Error
	}

	`

	return temp
}
