/*
 * KID validator.
 *
 * KIDs can use either Luhn or MOD11.
 *
 * Copyright (C) 2018 Folio AS.  All rights reserved.
 */

package checksum

// CheckKID verifies the MOD11 or Luhn checksum
func CheckKID(value string) bool {
	values := strToInts(value)
	end := len(values) - 1

	return mod11Check(mod11Sum(values[:end]), values[end]) ||
		luhnCheck(values)
}
