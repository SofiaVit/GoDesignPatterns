package factory

import (
	"GoDesignPatterns/factory/internal/vehicle"
	"errors"
)

func GetVehicle(km float32) (vehicle.IVehicle, error) {
	if km <= 3.5 {
		return &vehicle.Bicycle{}, nil
	} else if km > 3.5 && km <= 25 {
		return &vehicle.Car{}, nil
	} else {
		return nil, errors.New("delivery location is too far")
	}
}
