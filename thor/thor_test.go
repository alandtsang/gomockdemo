package thor

import (
	"testing"

	"github.com/alandtsang/gomockdemo/mock_hero"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestThor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHero := mock_hero.NewMockHero(ctrl)
	mockHero.EXPECT().GetName().Return("Thor").AnyTimes()
	mockHero.EXPECT().GetWeapon().Return("hammer").AnyTimes()

	thor := NewThor(mockHero)
	assert.Equal(t, thor.GetName(), mockHero.GetName())
	assert.Equal(t, thor.GetWeapon(), mockHero.GetWeapon())
}
