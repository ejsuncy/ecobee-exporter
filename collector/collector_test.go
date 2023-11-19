package collector

import "testing"
import "github.com/stretchr/testify/assert"

func Test_getSkyCondition(t *testing.T) {
	sky := -1
	want := "UNDEFINED"
	assert.Equal(t, getSkyCondition(sky), want)

	sky = 35
	assert.Equal(t, getSkyCondition(sky), want)

	sky = 34
	want = "LOW_LEVEL_HAZE"
	assert.Equal(t, getSkyCondition(sky), want)
}
