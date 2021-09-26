package cars

const carsProducedPerHour = 221

// CalculateProductionRatePerHour returns the number of cars completed in an hour.
func CalculateProductionRatePerHour(speed int) float64 {
	return float64(carsProducedPerHour*speed) * successRate(speed)
}

// CalculateProductionRatePerMinute returns the number of cars completed in an öimute.
func CalculateProductionRatePerMinute(speed int) int {
	return int(CalculateProductionRatePerHour(speed) / 60)
}

// successRate returns the success rate
func successRate(speed int) float64 {
	switch {
	case speed > 0 && speed < 5:
		return 1.0
	case speed > 4 && speed < 9:
		return 0.9
	case speed > 8 && speed < 11:
		return 0.77
	default:
		return 0
	}
}
