package sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAnnotation(t *testing.T) {
	a, ok := ParseAnnotation("one")
	assert.Equal(t, a, AnnotationOne)
	assert.True(t, ok)
	a, ok = ParseAnnotation("many")
	assert.Equal(t, a, AnnotationMany)
	assert.True(t, ok)
	a, ok = ParseAnnotation("exec")
	assert.Equal(t, a, AnnotationExec)
	assert.True(t, ok)
	a, ok = ParseAnnotation("batch")
	assert.Equal(t, a, AnnotationBatch)
	assert.True(t, ok)
	_, ok = ParseAnnotation("unknown")
	assert.False(t, ok)
}
