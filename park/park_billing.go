package park

type VehicleType int

const (
	Car VehicleType = iota
	Motorcycle
	Bicycle
)

func CalculateParkingBill(vehicle VehicleType, duration int) float32 {
	var fee float32

	if duration == 0 {
		if vehicle == VehicleType(0) {
			fee = 7000
		} else if vehicle == VehicleType(1) {
			fee = 3000
		} else {
			fee = 500
		}
	} else {
		if vehicle == VehicleType(0) {
			fee = float32((duration-1)*10000 + 7000 + duration/24*50000)
		} else if vehicle == VehicleType(1) {
			fee = float32((duration-1)*5000 + 3000 + duration/24*20000)
		} else {
			fee = float32((duration-1)*1000 + 500)
		}
	}

	return fee
}
