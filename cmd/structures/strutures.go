package structures

type HasWheels interface {
	CanStandWithoutMoving() bool
}

type OnWheels struct {
	number int
}

func (hasWheels *OnWheels) CanStandWithoutMoving() bool {
	return hasWheels.number > 3
}

type Car struct {
	*OnWheels
	SpeedLimit int
}

type Bike struct {
	*OnWheels
	ForTwoPeople bool
}

func NewCar() *Car {
	return &Car{
		&OnWheels{4}, 220,
	}
}

func NewBike() *Bike {
	return &Bike{
		&OnWheels{2}, true,
	}
}
