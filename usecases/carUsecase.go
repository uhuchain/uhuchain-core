package usecases

import "github.com/uhuchain/uhuchain-core/interfaces"

type CarUsecase struct {
	carProvider interfaces.CarProvider
}

func NewCarUsecase(carProvider interfaces.CarProvider) *CarUsecase {
	return &CarUsecase{
		carProvider: carProvider,
	}
}

func (usecase *CarUsecase) AddCar(car string) {
	usecase.carProvider.SaveCar(car)
}

func (usecase *CarUsecase) GetCar(id string) {
	usecase.carProvider.GetCar(id)
}
