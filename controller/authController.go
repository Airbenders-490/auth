package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"mocklogin/database"
	"mocklogin/model"
	"os"
	"strings"
	"time"
)

// SecretKey for jwt encoding
var SecretKey = os.Getenv("SECRET_KEY")

// Register user for account
func Register(context *fiber.Ctx) error {
	// context has our request data

	// extract the http request data
	var data map[string]string // declare var data of type map (key:string, value:string)

	// handle error if any
	if err := context.BodyParser(&data); err != nil {
		return err
	}

	// hash the password before inserting into the database
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	nameArr := strings.Fields(data["name"])

	// initialize User object with implicit type declaration
	user := model.User{
		ID: uuid.NewString(),
		FirstName: nameArr[0],
		LastName: nameArr[len(nameArr)-1],
		Email: data["email"],
		Password: password,
	}

	// insert reference to User object into database
	database.DatabaseConnection.Create(&user)


	// return http body content
	return context.JSON(user)
}

// Login user to account and return token
func Login(context *fiber.Ctx) error {
	var incomingCredentials map[string]string

	if err := context.BodyParser(&incomingCredentials); err != nil {
		return err
	}

	var user model.User

	// Get first matched record and assign it to user variable
	// ie SELECT * FROM users WHERE email = 'someEmail' ORDER BY id LIMIT 1;
	database.DatabaseConnection.
		Where("email = ?", incomingCredentials["email"]).
		First(&user)

	// if user doesn't exist
	// user looks like : &{""   []}
	// success user looks like :
	// &{3 carlin lee myemail2@gmail.com [36 50 97 36 49 52 ]}
	if user.ID == "" {
		context.Status(fiber.StatusNotFound)
		return context.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// if passwords don't match
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(incomingCredentials["password"])); err != nil {
		context.Status(fiber.StatusBadRequest)
		return context.JSON(fiber.Map{
			"message": "Incorrect Password",
		})
	}

	// in a payload, jwt Claims are used to provide authentication to the party receiving the token
	// The claim is digitally signed by the issuer of the token, and the party receiving this token
	// can later use this digital signature to prove the ownership on the claim.
	// https://www.softwaresecured.com/security-issues-jwt-authentication/

	type MyCustomClaims struct {
		FirstName 	string `json:"first_name"`
		LastName 	string `json:"last_name"`
		jwt.StandardClaims
	}

	// Create the Custom Claims
	claims := MyCustomClaims{
		user.FirstName,
		user.LastName,
		jwt.StandardClaims{
			Issuer:    user.ID, // reserved claim
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // expires in 24hours, reserved claim
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// The signature is used to verify the message wasn't changed along the way, and,
	// in the case of tokens signed with a private key, it can also verify that the
	// sender of the JWT is who it says it is
	signedToken, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		context.Status(fiber.StatusInternalServerError)
		return context.JSON(fiber.Map{
			"message": "Unable to Login",
		})
	}

	// frontend gets token then stores it their way
	return context.JSON(fiber.Map{
		"token": signedToken,
	})
}

// RetrieveUser info
func RetrieveUser(context *fiber.Ctx) error {
	// fetch user's jwt cookie string value by key
	jwtTokenString := context.Cookies("jwt")

	// want to be able to regenerate our SecretKey else error
	token, err := jwt.ParseWithClaims(jwtTokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	// if user was not logged in
	if err != nil {
		context.Status(fiber.StatusUnauthorized)
		return context.JSON(fiber.Map{
			"message": "Unauthenticated User; not logged in",
		})
	}

	standardClaims := token.Claims.(*jwt.StandardClaims) // Type assertion; interface to struct

	var user model.User

	database.DatabaseConnection.
		Where("id = ?", standardClaims.Issuer).
		First(&user)

	return context.JSON(user)
}