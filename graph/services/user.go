package service

import (
	"errors"
	"event_management/database"
	"event_management/graph/model"
	"fmt"
	"os"

	"github.com/doug-martin/goqu"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func GetUserById(userId string) (*model.User, error) {

	fmt.Println("trigger")
	db := database.Db

	rows, err := db.Query(`Select id, name, email, password, phone from "User" where id = $1`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var user model.User

	if !rows.Next() {
		return nil, errors.New("No User Found")
	}

	err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Phoneno)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser() ([]*model.User, error) {
	db := database.Db
	query := db.From("User").Select(
		goqu.I("id").As("id"),
		goqu.I("name").As("name"),
		goqu.I("email").As("email"),
		goqu.I("password").As("password"),
		goqu.I("phone").As("phoneno"),
	)

	var users []*model.User

	err := query.ScanStructs(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func RegisterUser(name, email, hashedPassword, phoneNo string) (*model.User, error) {
	db := database.Db

	var userId int
	err := db.QueryRow("INSERT INTO \"User\" (name,email,password,phone) VALUES ($1,$2,$3,$4) RETURNING id", name, email, hashedPassword, phoneNo).Scan(&userId)

	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:       fmt.Sprint(userId),
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		Phoneno:  phoneNo,
	}, nil
}

func LoginUser(email, password string) (*model.LoginResponse, error) {
	secret := os.Getenv("SECRET")
	fmt.Println(secret)
	db := database.Db

	rows, err := db.Query("SELECT id, name, email, password, phone from \"User\" where email = $1", email)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user model.User
	for rows.Next() {

		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Phoneno)
		if err != nil {

			return nil, err
		}

	}

	if user.ID == "" {
		return nil, errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("Invalid password")
	}
	var tokenString model.LoginResponse

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID, "name": user.Name,
	})

	tokenString.Token, err = token.SignedString([]byte(secret))
	if err != nil {
		return nil, errors.New("error while signing token")
	}

	return &tokenString, nil
}
