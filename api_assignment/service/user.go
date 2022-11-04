package service

import (
	"api_assignment/models"
	"log"
)

func (db *Database) CreateProd(prod *models.Product) error {

	err := db.Conn.Create(&prod)
	if err.Error != nil {
		log.Printf("Error Adding Product into database %s", err.Error)
		return err.Error
	}
	return nil

}

func (db *Database) ProdList(user uint) ([]models.Product, error) {
	var Prodlst []models.Product

	err := db.Conn.Where("user_id=?", user).Find(&Prodlst)
	if err.Error != nil {
		log.Printf("Error Getting Product list from database %s", err.Error)
		return []models.Product{}, err.Error
	}
	return Prodlst, nil

}

func (db *Database) UpdateList(user *uint, updateProd *models.Product) error {

	err := db.Conn.Model(&updateProd).Where("user_id=?", user).Updates(models.Product{Image: updateProd.Image, Name: updateProd.Name,
		Price: updateProd.Price})

	if err.Error != nil {
		log.Printf("Error Getting Product list from database %s", err.Error)
		return err.Error
	}
	return nil

}

func (db *Database) DeleteList(productID int, userID uint) error {

	err := db.Conn.Where("user_id=?", userID).Delete(&models.Product{}, productID)

	if err.Error != nil {
		log.Printf("Error Getting Product list from database %s", err.Error)
		return err.Error
	}
	return nil

}
