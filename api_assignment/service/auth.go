package service

import (
	"api_assignment/models"
	"errors"
	"log"
	"strings"
	"time"

	"gorm.io/gorm"
)

func (db Database) Login(login *models.LoginRequest) (*uint, error) {
	var res *gorm.DB
	var id *uint
	switch strings.ToLower(login.Type) {
	case "superadmin":
		var admin models.SuperAdmin
		res = db.Conn.Where("email=? AND password=?", login.Email, login.Password).Find(&admin)
		id = &admin.ID
	case "user":
		var user models.User
		res = db.Conn.Where("email=? AND password=?", login.Email, login.Password).Find(&user)
		id = &user.ID
	default:
		return nil, errors.New("Unknown Account Type")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("Wrong Username or Password")
	}

	return id, nil
}

func (db Database) TrackActivity(user *uint) (*uint, error) {

	activity := models.LoginActivity{
		UserID:    user,
		LoginTime: time.Now(),
	}
	res := db.Conn.Create(&activity)
	if res.Error != nil {
		log.Println("Error Storing activity Details")
	}

	return &activity.ID, nil
}

func (db Database) LogoutActivity(activtyID *uint) error {
	res := db.Conn.Model(&models.LoginActivity{}).Where("id=?", activtyID).Update("logout_time", time.Now())
	if res.Error != nil || res.RowsAffected == 0 {
		log.Printf("Error Storing in DB %s\n", res.Error)
		return res.Error
	}
	return nil
}
