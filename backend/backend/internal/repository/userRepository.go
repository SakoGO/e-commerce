package repository

import (
	model "e-commerce/backend/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) (*UserRepository, error) {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	}

	/*
		query := `CREATE TABLE users (
				email VARCHAR(255) UNIQUE NOT NULL,
				password VARCHAR(255) NOT NULL,
				username VARCHAR(255) UNIQUE NOT NULL,
				phone VARCHAR(255) UNIQUE NOT NULL,
				id INT PRIMARY KEY AUTO_INCREMENT,
				created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
				updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
				deleted_at DATETIME DEFAULT NULL,
				user_id INT NOT NULL,
				blocked BOOLEAN DEFAULT FALSE,
				is_admin BOOLEAN DEFAULT FALSE
			);`

			if err := db.Exec(query).Error; err != nil {
				log.Fatalf("Error executing query: %v", err)
			} else {
				log.Println("Table created successfully")
			}
	*/
	return &UserRepository{db: db}, nil
}

func (r *UserRepository) UserSignUP(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) UserFindByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UserFindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UserFindByPhone(phone string) (*model.User, error) {
	var user model.User
	err := r.db.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

/*
func (r *UserRepository) UserFindByID(userID int) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}


func (r *UserRepository) UserDelete(userID int) error {
	var user model.User
	err := r.db.First(&user, userID).Error
	if err != nil {
		return err
	}

	err = r.db.Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
*/
