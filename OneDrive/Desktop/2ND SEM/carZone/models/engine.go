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

func validateDisplacement(displacement int64) error {
	if displacement <= 0 {
		return errors.New("displacement is required")
	}
	return nil
}

