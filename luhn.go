package luhn

func IsValid(number int) bool {
	sum := number % 10
	number /= 10

	for i := 0; number > 0; i++ {
		digit := number % 10
		number /= 10

		if i%2 == 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}

	return sum%10 == 0
}

func IsValidStr(number string) bool {
	n := len(number)
	if n == 0 {
		return false
	}

	sum := 0
	for i := n - 1; i >= 0; i-- {
		digit := int(number[i] - '0')
		if digit < 0 || digit > 9 {
			return false
		}

		if (n-1-i)%2 == 1 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
	}

	return sum%10 == 0
}
