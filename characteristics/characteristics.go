package characteristics

type Characteristic interface {
	Get(name string) (map[string]struct{}, error)
}

// TODO: implement factory pattern to return the correct characteristic based on the name
func NewCharacteristic(name string) Characteristic {
	return nil
}
