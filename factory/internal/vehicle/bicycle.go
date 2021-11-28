package vehicle

type Bicycle struct{}

func (b *Bicycle) ScheduleDeliveryVehicle() string {
	return "Delivery by bicycle scheduled"
}
