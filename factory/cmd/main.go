package main

import (
	"GoDesignPatterns/factory/internal/factory"
	"GoDesignPatterns/factory/internal/vehicle"
)

func main() {
	scheduleDelivery(factory.GetVehicle(2.3))
	scheduleDelivery(factory.GetVehicle(21.1))
	scheduleDelivery(factory.GetVehicle(200))

}

func scheduleDelivery(deliveryVehicle vehicle.IVehicle, err error) {
	if err != nil {
		println(err.Error())
	} else {
		println(deliveryVehicle.ScheduleDeliveryVehicle())
	}
}
