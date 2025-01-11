package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Car struct {
	ID 		uuid.UUID `json:"id"`
	Name 	string `json:"name"`
	Year	string `json:"year"`
	Brand	string `json:"brand"`
	FuelType string `json:"fuel_type"`
	Engine Engine `json:"engine"`
	Price float64 `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
 
type CarRequest struct {
	
	Name 	string `json:"name"`
	Year	string `json:"year"`
	Brand	string `json:"brand"`
	FuelType string `json:"fuel_type"`
	Engine Engine `json:"engine"`
	Price float64 `json:"price"`

}

func ValidateRequest(carReq CarRequest) error {
	if err := validateName(carReq.Name); err != nil {
		return err
	}
	if err := validateYear(carReq.Year); err != nil {
		return err
	}
	if err := validateBrand(carReq.Brand); err != nil {
		return err
	}
	if err := validateFuelType(carReq.FuelType); err != nil {
		return err
	}
	if err := validateEngine(carReq.Engine); err != nil {
		return err
	}
	if err := validatePrice(carReq.Price); err != nil {
		return err
	}
	return nil
}
	
func validateName(name string) error {
	if name == "" {
		return errors.New("name is required")
	}
	return nil
}
func validateYear(year string) error {
	if year == "" {
		return errors.New("year is required")
	}
	_, err := strconv.Atoi(year)
	if err != nil {
		return errors.New("year must be a number")
	}
	currentYear := time.Now().Year()
	yearInt, _ := strconv.Atoi(year)
	if yearInt < 1886 || yearInt > currentYear {
		return errors.New("year must be between 1886 and current year")
	}
	return nil
}

func validateBrand(brand string) error {
	if brand == "" {
		return errors.New("brand is required")
	}
	return nil
}

func validateFuelType(fuelType string) error {
	validateFuelTypes := []string{"petrol", "diesel", "electric", "hybrid"}
	for _, validType := range validateFuelTypes {
		if fuelType == validType {
			return nil
		}
	}
	return errors.New("fuel type must be petrol, diesel, electric or hybrid")
}

func validateEngine(engine Engine) error {
	if engine.EngineID == uuid.Nil {
		return errors.New("engine_id is required")
	}
	if engine.Displacement <= 0 {
		return errors.New("displacement must be greater than zero")
	}
	if engine.NoOfCylinders <= 0 {
		return errors.New("no of cylinders is must be greater than zero")
	}
	if engine.CarRange <= 0 {
		return errors.New("car range must be greater than zero")
	}
	return nil
}

func validatePrice(price float64) error {
	if price <= 0 {
		return errors.New("price must be greater than zero")
	}
	return nil
}