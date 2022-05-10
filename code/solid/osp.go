package main

import "fmt"

// a software module should be open for extension but closed for modification

type (
	Sword struct {
		name string // Important tip for RPG players: always name your swords!
	}

	EnchantedSword struct {
		Sword
	}
)

// Damage returns the damage dealt by this sword.
func (Sword) Damage() int {
	return 2
}

// String implements fmt.Stringer for the Sword type.
func (s Sword) String() string {
	return fmt.Sprint(s.Damage())
}

// Damage returns the damage dealt by the enchanted sword.
func (EnchantedSword) Damage() int {
	return 42
}

func (s EnchantedSword) String() string {
	return fmt.Sprint(s.Damage())
}
