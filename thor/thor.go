package thor

import "github.com/alandtsang/gomockdemo/hero"

// Thor is the son of the Nordic god Odin
type Thor struct {
	hero hero.Hero
}

// NewThor creates Thor instance
func NewThor(h hero.Hero) *Thor {
	return &Thor{
		hero: h,
	}
}

// GetName return Thor's name
func (t *Thor) GetName() string {
	return t.hero.GetName()
}

// GetName return Thor's weapon
func (t *Thor) GetWeapon() string {
	return t.hero.GetWeapon()
}
