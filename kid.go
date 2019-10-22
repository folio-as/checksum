/*
 * KID validator.
 *
 * KIDs can use either Luhn or MOD11.
 *
 * Copyright (C) 2018 Folio AS.  All rights reserved.
 */

package checksum

// ValidateKID computes the MOD11 or Luhn checksums and returns true if either are correct
func ValidateKID(value string) bool {
	values := strToInts(value)
	if values == nil {
		return false
	}
	end := len(values) - 1

	return mod11Check(mod11Sum(values[:end]), values[end]) ||
		luhnCheck(values)
}
