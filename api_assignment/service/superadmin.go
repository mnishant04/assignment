package service

import (
	"api_assignment/models"
	"log"

	"gorm.io/gorm"
)

type Database struct {
	Conn *gorm.DB
}

func (db *Database) AddUser(user *models.User) error {
	log.Printf("Db is %v", db)
	err := db.Conn.Create(&user)
	if err.Error != nil {
		log.Printf("Error adding User in to database %s", err.Error)
		return err.Error
	}
	return nil
}

func (db *Database) ListUser() (allUsers []models.User, err error) {

	if err = db.Conn.Find(&allUsers).Error; err != nil {
		log.Printf("Error fetching User from database %s", err)
		return allUsers, err
	}

	return allUsers, nil
}

func (db *Database) ListUserProducts() (result []models.Product, err error) {

	if err = db.Conn.Find(&result).Error; err != nil {
		log.Printf("Error fetching User from database %s", err)
		return []models.Product{}, err
	}

	return result, nil
}

func (db *Database) TrackUserActivity() (allUsersActivity []models.LoginActivity, err error) {

	if err = db.Conn.Find(&allUsersActivity).Error; err != nil {
		log.Printf("Error fetching User from database %s", err)
		return []models.LoginActivity{}, err
	}

	return allUsersActivity, nil
}
