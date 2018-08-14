/*
 * Norwegian SSN validator.
 *
 * Copyright (C) 2018 Folio AS.  All rights reserved.
 */

package checksum

func ssnSum(value []int) (sum int) {
	weights := []int{3, 7, 6, 1, 8, 9, 4, 5, 2}

	for i, v := range value {
		sum += v * weights[i]
	}

	return
}

// ValidateNorwegianSSN computes MOD11 checksums and returns true if both are correct
func ValidateNorwegianSSN(value string) bool {
	values := strToInts(value)
	end := len(values) - 1

	if end != 10 {
		return false
	}

	return mod11Check(mod11Sum(values[:end]), values[end]) &&
		mod11Check(ssnSum(values[:end-1]), values[end-1])
}
