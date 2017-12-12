package controller

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/uhuchain/uhuchain-core/interfaces"
	"github.com/uhuchain/uhuchain-core/models"
	"github.com/uhuchain/uhuchain-core/usecases"
)

// CarController handles car usecases
type CarController struct {
	provider interfaces.CarProvider
}

// NewCarController creates a new car controller
func NewCarController(provider interfaces.CarProvider) *CarController {
	return &CarController{
		provider: provider,
	}
}

// Run the car controller
func (controller *CarController) Run(function string, args []string) ([]byte, error) {
	carUsecase := usecases.NewCarUsecase(controller.provider)
	if function == "saveCar" {
		return nil, controller.saveCar(carUsecase, args)
	} else if function == "getCar" {
		return controller.getCar(carUsecase, args)
	}
	return nil, fmt.Errorf("Invalid invoke function name \"%s\". Expecting \"invoke\" \"delete\" \"query\"", function)
}

// Write a car with a given ID onto the ledger
func (controller *CarController) saveCar(usecase *usecases.CarUsecase, args []string) error {
	if len(args) != 1 {
		return errors.New("Incorrect number of arguments. Expecting ")
	}
	carValue := []byte(args[0])
	car := models.Car{}
	err := car.UnmarshalBinary(carValue)
	if err != nil {
		return err
	}
	err = usecase.SaveCar(car)
	if err != nil {
		return err
	}
	return nil
}

func (controller *CarController) getCar(usecase *usecases.CarUsecase, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting ")
	}
	carID, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return nil, err
	}
	car, err := usecase.GetCar(carID)
	if err != nil {
		return nil, err
	}
	carValue, err := car.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return carValue, nil
}
