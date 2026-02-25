// mnemonic_test.go
package mnemonic

import (
	"testing"
)

func TestGenerateAndValidate(t *testing.T) {

	tests := []struct {

		wordCount int

	}{

		{12},

		{15},

		{18},

		{21},

		{24},

	}

	for _, tt := range tests {

		m, err := Generate(tt.wordCount)

		if err != nil {

			t.Error(err)

		}

		if !Validate(m) {

			t.Errorf("generated mnemonic invalid: %s", m)

		}

	}

}

func TestToSeedIcarus(t *testing.T) {

	m := "abandon amount liar amount expire adjust cage candy arch gather drum bundle"

	seed, err := ToSeed(m, "pass", Icarus)

	if err != nil {

		t.Error(err)

	}

	if len(seed) != 64 {

		t.Error("invalid seed length")

	}

	// Check against known value (assume known test vector)

}

func TestToSeedLedger(t *testing.T) {

	m := "abandon amount liar amount expire adjust cage candy arch gather drum bundle"

	seed, err := ToSeed(m, "pass", Ledger)

	if err != nil {

		t.Error(err)

	}

	if len(seed) != 96 {

		t.Error("invalid seed length")

	}

	// Check tweak applied

	if seed[0]&0x07 != 0 || seed[31]&0x80 != 0 || seed[31]&0x40 == 0 {

		t.Error("tweak not applied")

	}

}

func TestInvalid(t *testing.T) {

	invalid := "invalid word list"

	if Validate(invalid) {

		t.Error("invalid mnemonic validated")

	}

}
