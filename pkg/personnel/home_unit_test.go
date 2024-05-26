package personnel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseHomeUnit(t *testing.T) {
	charter := "RMR-UT-067"

	unit, err := ParseHomeUnit(charter, "")
	assert.NoError(t, err)

	assert.Equal(t, RockyMountainRegion, unit.Region())
	assert.Equal(t, UTWG, unit.Wing())
	assert.Equal(t, uint(67), unit.UnitNumber())

	// UT is not in SER
	charter = "SER-UT-067"

	_, err = ParseHomeUnit(charter, "")
	assert.Error(t, err)
}

func TestHomeUnit_ShortCharterNumber(t *testing.T) {
	charter := "RMR-UT-067"

	unit, err := ParseHomeUnit(charter, "")
	assert.NoError(t, err)

	out := unit.ShortCharterNumber()

	assert.Equal(t, "UT-067", out)
}

func TestHomeUnit_FullCharterNumber(t *testing.T) {
	charter := "RMR-UT-067"

	unit, err := ParseHomeUnit(charter, "")
	assert.NoError(t, err)

	out := unit.FullCharterNumber()

	assert.Equal(t, "RMR-UT-067", out)
}
