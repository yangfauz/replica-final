package main

import (
	"log"
	"replica-finalproject/config"
	"replica-finalproject/database/migration"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestMigration(t *testing.T) {
	//setup configuration
	configuration := config.New()
	database := config.NewMysqlDatabase(configuration)

	database.AutoMigrate(&migration.Role{}, &migration.Status{}, &migration.User{}, &migration.Payment{})
}

func TestRoleSeeder(t *testing.T) {
	//setup configuration
	configuration := config.New()
	database := config.NewMysqlDatabase(configuration)

	var roles = []migration.Role{
		{Name: "Admin"},
		{Name: "Unit Kerja (Customer)"},
		{Name: "General Support"},
		{Name: "Accounting"},
	}

	err := database.Create(&roles).Error

	if err != nil {
		log.Println("Role Seed Failed")
	}

	log.Println("Role Seed Success")
}

func TestStatusSeeder(t *testing.T) {
	//setup configuration
	configuration := config.New()
	database := config.NewMysqlDatabase(configuration)

	var statuses = []migration.Status{
		{Name: "Menunggu Konfirmasi"},
		{Name: "Reject By General Support"},
		{Name: "Diteruskan ke Accounting"},
		{Name: "Reject By Accounting"},
		{Name: "Disetujui Accounting"},
	}

	err := database.Create(&statuses).Error

	if err != nil {
		log.Println("Status Seed Failed")
	}

	log.Println("Status Seed Success")
}

func TestUserSeeder(t *testing.T) {
	//setup configuration
	configuration := config.New()
	database := config.NewMysqlDatabase(configuration)

	passwordHash, err := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.MinCost)

	var user = migration.User{
		Name:     "Admin",
		Username: "admin",
		Password: string(passwordHash),
		RoleID:   1,
	}

	err = database.Create(&user).Error

	if err != nil {
		log.Println("User Seed Failed")
	}

	log.Println("User Seed Success")

}
