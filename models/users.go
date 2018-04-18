package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrNotFound is returned when a resource cannot be found in the database.
	ErrNotFound = errors.New("models: resource not found")
	// ErrInvalidID is returned when an invalid ID is provided to a method like Delete.
	ErrInvalidID = errors.New("models: ID provided was invalid")
	// ErrInvalidPassword is returned when an invalid password is used when attempting to authenticate a user.
	ErrInvalidPassword = errors.New("models: incorrect password provided")
	)

	var userPwPepper = "secret-random-string"

//User type reps user resource
type User struct {
	gorm.Model //includes the id, created_at, updated_at, and created_at fields
	Name string
	Email string `gorm:"not null;unique_index"`
	Password string `gorm:"-"` //stores the raw (unhashed) password. The "-" tag denotes that we will NEVER save this field in the database,
	PasswordHash string `gorm:"not null"`
}

//UserService defines the abstraction layer for our users database - a way of hiding implementation details.
type UserService struct {
	db *gorm.DB
}

//NewUserService
func NewUserService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil{
		return nil, err
	}
	db.LogMode(true)
	return &UserService{
		db: db,
	}, nil
}
// Close closes the UserService database connection
//if we were to defer closing the DB inside of our NewUserService function, it would end up closing the database connection right before returning the new user service.
func (us *UserService) Close() error {
	return us.db.Close()
}

// first will query using the provided gorm.DB and it will get the first item returned and place it into dst. 
// If nothing is found in the query, it will return ErrNotFound
func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err
}

// ByID will look up a user with the provided ID.
// If the user is found, we will return a nil error
// If the user is not found, we will return ErrNotFound
// If there is another error, we will return an error with more information about what went wrong
// As a general rule, any error but ErrNotFound should probably result in a 500 error.
func (us *UserService) ByID(id uint) (*User, error) {
	var user User
	db := us.db.Where("id = ?", id)
	err := first(db, &user)
	if err!=nil{  
		return nil, err
	}
	return &user, nil
}

// ByEmail looks up a user with the given email address and
// returns that user.
// If the user is found, we will return a nil error
// If the user is not found, we will return ErrNotFound
// If there is another error, we will return an error with
// more information about what went wrong. This may not be
// an error generated by the models package.
func (us *UserService) ByEmail(email string) (*User, error) {
	var user User
	db := us.db.Where("email = ?", email)
	err := first(db, &user)
	return &user, err
}

// DestructiveReset drops the user table and rebuilds it. Useful for quick reset of table during development.
// Not to be used in production
func (us *UserService) DestructiveReset() error {
	err:= us.db.DropTableIfExists(&User{}).Error
	if err != nil {
		return err
	}
	return us.AutoMigrate()
}
// AutoMigrate will attempt to automatically migrate the users table
func (us *UserService) AutoMigrate() error {
	if err := us.db.AutoMigrate(&User{}).Error; err != nil {
		return err
	}
	return nil
}
// Create will create the provided user and backfill data
// like the ID, CreatedAt, and UpdatedAt fields.
func (us *UserService) Create(user *User) error{
	pwBytes := []byte(user.Password + userPwPepper)
	hashedBytes, err := bcrypt.GenerateFromPassword(
		pwBytes, bcrypt.DefaultCost) //DefaultCost dictates how much work (and sometimes memory) must be used to hash a password.
	if err != nil{
		return err
	}
	user.PasswordHash = string(hashedBytes)
	user.Password = ""
	return us.db.Create(user).Error
}
// Update will update the provided user with all of the data in the provided user object.
func (us *UserService) Update(user *User) error {
	return us.db.Save(user).Error
}

// Delete will delete the user with the provided ID
func (us *UserService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}
	user := User{Model: gorm.Model{ID: id}}
	return us.db.Delete(&user).Error
}

// Authenticate can be used to authenticate a user with the provided email address and password.
// If the email address provided is invalid, this will return nil, ErrNotFound
// If the password provided is invalid, this will return nil, ErrInvalidPassword
func (us *UserService) Authenticate(email, password string) (*User, error){
	foundUser, err := us.ByEmail(email) // returns the ErrNotFound error when a user isn’t found with the email address.
	if err != nil{
		return nil, err
	}
	//Check if passwords match
	err = bcrypt.CompareHashAndPassword(
		[]byte(foundUser.PasswordHash),
		[]byte(password+userPwPepper))
	switch err {
	case nil:
		return foundUser, nil
	case bcrypt.ErrMismatchedHashAndPassword:
		return nil, ErrInvalidPassword
	default:
		return nil, err
	}
	return nil, nil
}