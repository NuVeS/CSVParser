package finder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	assert := assert.New(t)

	rows := []string{"id", "Name", "Age"}
	index := defineColumnIndex("Name", rows)
	assert.Equal(index, 1)

	index = defineColumnIndex("id", rows)
	assert.Equal(index, 0)
}
