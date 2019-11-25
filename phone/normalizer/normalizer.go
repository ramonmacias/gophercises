package normalizer

func Normalize(phoneNumber string) string {
	for _, filter := range FilterFuncs() {
		phoneNumber = filter(phoneNumber)
	}
	return phoneNumber
}

func BatchNormalize(phoneNumbers []string) (normalizedNumbers []string) {
	for _, phone := range phoneNumbers {
		normalizedNumbers = append(normalizedNumbers, Normalize(phone))
	}
	return normalizedNumbers
}
