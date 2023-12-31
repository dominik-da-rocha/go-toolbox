package toolbox_test

import (
	"testing"

	"github.com/dominik-da-rocha/go-toolbox/toolbox"
	"github.com/stretchr/testify/assert"
)

type TestValue struct {
	val1      int `my-tag:"hello"`
	val2      int `my-tag:"world"`
	val3      int `my-tag:"!"`
	ignoreTag int `my-tag:"-"`
	notTagged int
}

func Test_GetStructNames(t *testing.T) {
	val := TestValue{
		val1:      1,
		val2:      2,
		val3:      3,
		ignoreTag: -1,
		notTagged: 4,
	}
	valPtr := &val
	names := toolbox.GetStructTags(valPtr, "my-tag")
	assert.Equal(t, 4, len(names))
	assert.Equal(t, "hello", names[0])
	assert.Equal(t, "world", names[1])
	assert.Equal(t, "!", names[2])
	assert.Equal(t, "notTagged", names[3])
}

func Test_GetStructNamesWithArrayPanics(t *testing.T) {
	assert.Panics(t, func() {
		value := []string{}
		toolbox.GetStructTags(value, "my-tag")
	})
}
