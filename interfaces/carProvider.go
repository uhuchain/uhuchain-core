package interfaces

import "github.com/uhuchain/uhuchain-core/models"

// CarProvider defines the interface how to handle cars
type CarProvider interface {
	SaveCar(models.Car) error
	GetCar(int64) (models.Car, error)
}
