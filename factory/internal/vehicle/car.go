package vehicle

type Car struct{}

func (c *Car) ScheduleDeliveryVehicle() string {
	return "Delivery by Car scheduled"
}
