# go-cardano-mnemonic

Lightweight, dependency-free Go library for Cardano mnemonic generation, validation, and seed derivation supporting Icarus and Ledger styles as per CIP-3.

## Installation

```sh
go get github.com/njchilds90/go-cardano-mnemonic

package main

import (
	"fmt"
	"github.com/njchilds90/go-cardano-mnemonic/mnemonic"
)

func main() {
	// Generate a 15-word mnemonic
	m, err := mnemonic.Generate(15)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Mnemonic:", m)

	// Validate
	if !mnemonic.Validate(m) {
		fmt.Println("Invalid mnemonic")
		return
	}

	// Derivation to seed (Icarus)
	seed, err := mnemonic.ToSeed(m, "", mnemonic.Icarus)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Icarus seed: %x\n", seed)

	// Ledger
	seed, err = mnemonic.ToSeed(m, "", mnemonic.Ledger)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Ledger seed: %x\n", seed)
}

Features

Deterministic generation and derivation.
Structured errors for machine parsing.
Pure functions, no global state.


```markdown
# CHANGELOG.md

## v1.0.0 (2026-02-25)

- Initial release with mnemonic generation, validation, and seed derivation for Icarus and Ledger styles.

