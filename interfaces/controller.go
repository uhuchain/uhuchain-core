package interfaces

// Controller defines the interface for a controller
type Controller interface {
	Run(function string, args []string) ([]byte, error)
}
