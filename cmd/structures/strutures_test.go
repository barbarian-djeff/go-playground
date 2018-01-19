package structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Structures(t *testing.T) {
	var car HasWheels
	car = NewCar()
	assert.Equal(t, car.CanStandWithoutMoving(), true)
	assert.Equal(t, NewBike().CanStandWithoutMoving(), false)

	typedCar := NewCar()
	assert.Equal(t, typedCar.SpeedLimit, 220)
	assert.Equal(t, NewBike().ForTwoPeople, true)
}
