package luhn

func IsValid(number int) bool {
	sum := number % 10
	number /= 10
	even := true
	for number > 0 {
		digit := number % 10
		number /= 10
		if even {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		even = !even
	}
	return sum%10 == 0
}

func IsValidStr(number string) bool {
	n := len(number)
	if n == 0 {
		return false
	}
	sum := 0
	even := false
	for i := n - 1; i >= 0; i-- {
		digit := int(number[i] - '0')
		if digit < 0 || digit > 9 {
			return false
		}
		if even {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		even = !even
	}
	return sum%10 == 0
}
