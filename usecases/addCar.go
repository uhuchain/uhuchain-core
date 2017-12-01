package usecases

import "github.com/uhuchain/uhuchain-api/core/interfaces"

type AddCarUsecase struct {
	carProvider interfaces.CarProvider
}

func NewAddCarUsecase(carProvider interfaces.CarProvider) AddCarUsecase {
	return AddCarUsecase{
		carProvider: carProvider,
	}
}

func (usecase *AddCarUsecase) addCar() {

}
