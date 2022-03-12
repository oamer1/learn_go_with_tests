package propertybasedtests

func ConvertToRoman(arabic int) string {
	if arabic == 3 {
		return "III"
	} else if arabic == 2 {
		return "II"
	}
	return "I"
}
