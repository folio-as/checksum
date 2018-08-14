/*
 * Luhn checksum validators.
 *
 * Used for validation of social security numbers and KID numbers.
 *
 * Copyright (C) 2018 Folio AS.  All rights reserved.
 */

package checksum

func luhnSum(value []int) (checksum int) {
	end := len(value) - 1
	for i := range value {
		v := value[end-i]
		if i%2 == 1 {
			if v += v; v >= 10 {
				v -= 9
			}
		}
		checksum += v
	}

	return
}

func luhnCheck(values []int) bool {
	return (luhnSum(values) % 10) == 0
}

// Luhn verifies the checksum
func Luhn(value string) bool {
	values := strToInts(value)

	return luhnCheck(values)
}
