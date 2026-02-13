package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	res := false
	if kind == "car" {
		res = true
	} else if kind == "truck" {
		res = true
	}
	return res
	// panic("NeedsLicense not implemented")
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	suffixString := " is clearly the better choice."
	var res string
	if option1 > option2 {
		res = option2 + suffixString
	} else {
		res = option1 + suffixString
	}
	return res
	// panic("ChooseVehicle not implemented")
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var res float64
	if age < 3 {
		res = originalPrice * 0.8
	} else if age < 10 {
		res = originalPrice * 0.7
	} else {
		res = originalPrice * 0.5
	}
	return res
	// panic("CalculateResellPrice not implemented")
}
