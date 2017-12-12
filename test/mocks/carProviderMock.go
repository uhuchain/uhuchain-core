package mocks

import (
	"github.com/uhuchain/uhuchain-core/models"
)

// CarProviderMock mocks the car provider
type CarProviderMock struct {
}

// SaveCar implements the the provider using hlf
func (p *CarProviderMock) SaveCar(car models.Car) error {
	return nil
}

// GetCar implements the the provider using hlf
func (p *CarProviderMock) GetCar(id int64) (models.Car, error) {
	car := createCar()
	return car, nil
}
