/*
 * Various MOD11 computations.
 *
 * Used for validation of social security numbers, account numbers, organization numbers
 * and KID numbers.
 *
 * Copyright (C) 2020 Folio AS.  All rights reserved.
 */

package checksum

func strToInts(s string) []int {
	if len(s) == 0 {
		return nil
	}
	a := make([]int, len(s))
	for i, v := range s {
		if v < '0' || v > '9' {
			return nil
		}
		a[i] = int(v - '0')
	}
	return a
}

func mod11Check(sum, checksum int) bool {
	sum %= 11
	if sum == 0 {
		sum = 11
	}
	return (11 - sum) == checksum
}

func mod11Sum(values []int) (sum int) {
	end := len(values) - 1
	for i, v := range values {
		sum += v * ((end-i)%6 + 2)
	}

	return
}

// Mod11 verifies the checksum
func Mod11(value string) bool {
	end := len(value) - 1
	if end < 0 {
		return false
	}

	checksum := 10
	v := value[end]
	if v >= '0' && v <= '9' {
		checksum = int(v - '0')
	} else if v != '-' {
		return false
	}

	values := strToInts(value[:end])
	if values == nil {
		return false
	}

	return mod11Check(mod11Sum(values), checksum)
}
