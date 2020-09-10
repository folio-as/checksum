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
	return Mod11(value) ||
		Luhn(value)
}
