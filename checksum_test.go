/*
 * Test checksum computations.
 *
 * Copyright (C) 2018 Folio AS.  All rights reserved.
 */

package checksum

import "testing"

func TestChecksum(t *testing.T) {
	if !Mod11("01010750160") {
		t.Errorf("failed")
	}
	if !CheckNorwegianSSN("01010750241") {
		t.Errorf("failed")
	}
	if !Mod11("919638095") {
		t.Errorf("failed")
	}
	if CheckNorwegianSSN("01010750161") {
		t.Errorf("failed")
	}
	if !Luhn("79927398713") {
		t.Errorf("failed")
	}
	if !Mod11("0095636818940001513") {
		t.Errorf("failed")
	}
	if !CheckKID("79927398713") {
		t.Errorf("failed")
	}
	if !CheckKID("0095636818940001513") {
		t.Errorf("failed")
	}
}