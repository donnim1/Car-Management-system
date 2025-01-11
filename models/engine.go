package models

import ( "github.com/google/uuid"	
	"errors"
)
 



type Engine struct {
	EngineID 		uuid.UUID `json:"engine_id"`
	Displacement 	int64 `json:"displacement"`
	NoOfCylinders 	int64 `json:"noOfCylinders"`
	CarRange        int64 `json:"carRange"`

}
type EngineRequest struct {
	Displacement 	int64 `json:"displacement"`
	NoOfCylinders 	int64 `json:"noOfCylinders"`
	CarRange        int64 `json:"carRange"`

}
func ValidateEngineRequest

func validateDisplacement(displacement int64) error {
	if displacement <= 0 {
		return errors.New("displacement is required")
	}
	return nil
}

func validateNoOfCylinders(noOfCylinders int64) error {
	if noOfCylinders <= 0 {
		return errors.New("noOfCylinders is required")
	}
	return nil
}


func validateCarRange(carRange int64) error {
	if carRange <= 0 {
		return errors.New("carRange is required")
	}
	return nil
}
