package interfaces

// CarProvider defines the interface how to handle cars
type CarProvider interface {
	SaveCar(string) error
	GetCar(string) (string, error)
}
