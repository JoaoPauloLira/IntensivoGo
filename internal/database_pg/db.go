package database_pg

import (
	"gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	connectString := "host=postgres user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectString))
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados: " + err.Error())
	}
}
