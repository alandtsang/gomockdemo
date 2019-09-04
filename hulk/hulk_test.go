package hulk

import (
	"testing"

	"github.com/alandtsang/gomockdemo/mock_hero"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestHulk(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHero := mock_hero.NewMockHero(ctrl)
	mockHero.EXPECT().GetName().Return("Hulk")
	mockHero.EXPECT().GetWeapon().Return("fist")

	hulk := NewHulk()
	assert.Equal(t, hulk.GetName(), mockHero.GetName())
	assert.Equal(t, hulk.GetWeapon(), mockHero.GetWeapon())
}
