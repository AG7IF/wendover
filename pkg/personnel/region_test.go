package personnel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegion_WingIsInRegion(t *testing.T) {
	// UT is in RMR
	var region Region = RMR
	res := region.WingIsInRegion(UTWG)
	assert.True(t, res)
}
