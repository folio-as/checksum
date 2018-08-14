/*
 * Various MOD11 computations.
 *
 * Used for validation of social security numbers, account numbers, organization numbers
 * and KID numbers.
 *
 * Copyright (C) 2018 Folio AS.  All rights reserved.
 */

package checksum

func strToInts(s string) []int {
	a := make([]int, len(s))
	for i, v := range s {
		a[i] = int(v - '0')
	}
	return a
}

func mod11Check(sum, checksum int) bool {
	sum %= 11
	if sum == 0 {
		return 0 == checksum
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
	values := strToInts(value)
	end := len(values) - 1

	return mod11Check(mod11Sum(values[:end]), values[end])
}
