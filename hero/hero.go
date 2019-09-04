package hero

//go:generate mockgen -destination ../mock_hero/mock_hero.go -package mock_hero -source hero.go

// Hero is a Marvel hero interface
type Hero interface {
	GetName() string
	GetWeapon() string
}
