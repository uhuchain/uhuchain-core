package usecases

import "github.com/uhuchain/uhuchain-core/interfaces"
import "github.com/uhuchain/uhuchain-core/models"

// CarUsecase handles cars on uhuchain
type CarUsecase struct {
	carProvider interfaces.CarProvider
}

// NewCarUsecase creates a new car usecase with a given dataprovide
func NewCarUsecase(carProvider interfaces.CarProvider) *CarUsecase {
	return &CarUsecase{
		carProvider: carProvider,
	}
}

// SaveCar writes a car onto the uhuchain ledger
func (usecase *CarUsecase) SaveCar(car models.Car) error {
	err := usecase.carProvider.SaveCar(car)
	return err
}

// GetCar retrieves a car with a given ID from the ledger
func (usecase *CarUsecase) GetCar(id int64) (models.Car, error) {
	car, err := usecase.carProvider.GetCar(id)
	return car, err
}
