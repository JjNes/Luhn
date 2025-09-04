package luhn

func IsValidInt(number int) bool {
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

func GenerateLuhnNumber(number int) int {
	sum := luhnCheckDigit(number)

	fullNumder := number*10 + sum

	return fullNumder
}

func luhnCheckDigit(number int) int {
	sum := 0
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
	return (10 - (sum % 10)) % 10
}
