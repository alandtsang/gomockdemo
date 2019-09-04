package hulk

import "github.com/alandtsang/gomockdemo/hero"

// Hulk is the Hulk
type Hulk struct {
	name   string
	weapon string
}

// NewHulk creates Hulk instance
func NewHulk() hero.Hero {
	return &Hulk{name: "Hulk", weapon: "fist"}
}

// GetName return Hulk's name
func (hk *Hulk) GetName() string {
	return hk.name
}

// GetWeapon return Hulk's weapon
func (hk *Hulk) GetWeapon() string {
	return hk.weapon
}
