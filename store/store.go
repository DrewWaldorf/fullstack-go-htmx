package store

import (
	"database/sql"
	"time"

	"github.com/sikozonpc/fullstackgo/types"
)

type Storage struct {
	db *sql.DB
}

type Store interface {
	CreateCar(car *types.Car) (*types.Car, error)
	GetCars() ([]types.Car, error)
	DeleteCar(id string) error
	FindCarsByNameMakeOrBrand(search string) ([]types.Car, error)
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) DeleteCar(id string) error {
	_, err := s.db.Exec("DELETE FROM cars WHERE id = $1", id)
	return err
}

func (s *Storage) CreateCar(c *types.Car) (*types.Car, error) {
	var id int
	err := s.db.QueryRow("INSERT INTO cars (brand, make, model, year, imageURL, created) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", c.Brand, c.Make, c.Model, c.Year, c.ImageURL, time.Now()).Scan(&id)
	if err != nil {
		return nil, err
	}
	c.ID = id

	return c, nil
}

func (s *Storage) GetCars() ([]types.Car, error) {
	rows, err := s.db.Query("SELECT * FROM cars")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cars []types.Car
	for rows.Next() {
		car, err := scanCar(rows)
		if err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}

	return cars, nil
}

func (s *Storage) FindCarsByNameMakeOrBrand(search string) ([]types.Car, error) {
	rows, err := s.db.Query("SELECT * FROM cars WHERE LOWER(brand) LIKE LOWER($1) OR LOWER(model) LIKE LOWER($2) OR LOWER(make) LIKE LOWER($3)", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cars []types.Car
	for rows.Next() {
		car, err := scanCar(rows)
		if err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}

	return cars, nil
}

func scanCar(row *sql.Rows) (types.Car, error) {
	var car types.Car
	err := row.Scan(&car.ID, &car.Brand, &car.Make, &car.Model, &car.Year, &car.ImageURL, &car.CreatedAt)
	return car, err
}
