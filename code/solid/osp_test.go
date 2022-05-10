package main

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	type fields struct {
		name string
	}
	tests := []struct {
		sword fmt.Stringer
		exp   string
		name  string
	}{
		{
			sword: Sword{
				name: "sword",
			},
			exp:  "2",
			name: "test sword dame",
		},
		{
			sword: EnchantedSword{
				Sword: Sword{
					name: "enchanted sword",
				},
			},
			exp:  "42",
			name: "test enchanted sword dame",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.sword.String(); got != tt.exp {
				t.Errorf("Sword.String() = %v, want %v", got, tt.exp)
			}
		})
	}
}

func TestSword_Damage(t *testing.T) {
	type fields struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Sword{
				name: tt.fields.name,
			}
			if got := s.Damage(); got != tt.want {
				t.Errorf("Sword.Damage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSword_String(t *testing.T) {
	type fields struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Sword{
				name: tt.fields.name,
			}
			if got := s.String(); got != tt.want {
				t.Errorf("Sword.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnchantedSword_Damage(t *testing.T) {
	type fields struct {
		Sword Sword
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := EnchantedSword{
				Sword: tt.fields.Sword,
			}
			if got := e.Damage(); got != tt.want {
				t.Errorf("EnchantedSword.Damage() = %v, want %v", got, tt.want)
			}
		})
	}
}
