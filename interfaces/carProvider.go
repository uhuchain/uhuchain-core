package interfaces

import "github.com/uhuchain/uhuchain-core/models"

// CarProvider defines the interface how to handle cars
type CarProvider interface {
	saveCar(models.Car) error
	getCar(string) (models.Car, error)
}
