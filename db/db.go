package db

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB
var err error

type Movie struct {
	Id          string `json:"id" gorm:"primarykey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CreateMovie(movie *Movie) (*Movie, error) {
	movie.Id = uuid.New().String()
	res := db.Create(&movie)
	if res.Error != nil {
		return nil, res.Error
	}
	return movie, nil
}

func GetMovie(id string) (*Movie, error) {
	var movie Movie
	res := db.First(&movie, "id = ?", id)
	if res.RowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("movie of id %s not found", id))
	}
	return &movie, nil
}

func GetMovies() ([]*Movie, error) {
	var movies []*Movie
	res := db.Find(&movies)
	if res.Error != nil {
		return nil, errors.New("no movies found")
	}
	return movies, nil
}

func UpdateMovie(movie *Movie) (*Movie, error) {
	var movieToUpdate Movie
	result := db.Model(&movieToUpdate).Where("id = ?", movie.Id).Updates(movie)
	if result.RowsAffected == 0 {
		return &movieToUpdate, errors.New("movie not updated")
	}
	return movie, nil
}

func DeleteMovie(id string) error {
	var deletedMovie Movie
	result := db.Where("id = ?", id).Delete(&deletedMovie)
	if result.RowsAffected == 0 {
		return errors.New("movie not deleted")
	}
	return nil
}

func InitMysqlDB() {
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		dbUser   = os.Getenv("DB_USER")
		dbName   = os.Getenv("DB_NAME")
		password = os.Getenv("DB_PASSWORD")
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		password,
		host,
		port,
		dbName,
	)

	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(Movie{})
}
