package sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToSingularPascalCase(t *testing.T) {
	assert.Equal(t, "Lion", ToSingularPascalCase("lions"))
	assert.Equal(t, "ID", ToSingularPascalCase("ids"))
	assert.Equal(t, "Knife", ToSingularPascalCase("knives"))
	assert.Equal(t, "PistolsID", ToSingularPascalCase("pistols_ids"))
	assert.Equal(t, "PistolsID", ToSingularPascalCase("pistolsIds"))
	assert.Equal(t, "PistolsID", ToSingularPascalCase("PistolsIds"))
	assert.Equal(t, "PistolsID", ToSingularPascalCase("PistolsID"))
}

func TestToSingularSnakeCase(t *testing.T) {
	assert.Equal(t, "lion", ToSingularSnakeCase("lions"))
	assert.Equal(t, "id", ToSingularSnakeCase("ids"))
	assert.Equal(t, "knife", ToSingularSnakeCase("knives"))
	assert.Equal(t, "pistols_id", ToSingularSnakeCase("pistols_ids"))
	assert.Equal(t, "pistols_id", ToSingularSnakeCase("pistolsIds"))
	assert.Equal(t, "pistols_id", ToSingularSnakeCase("PistolsIds"))
	assert.Equal(t, "pistols_id", ToSingularSnakeCase("PistolsID"))
}

func TestToSnakeCase(t *testing.T) {
	assert.Equal(t, "pistols_ids", ToSnakeCase("pistols_ids"))
	assert.Equal(t, "pistols_ids", ToSnakeCase("pistolsIds"))
	assert.Equal(t, "pistols_ids", ToSnakeCase("PistolsIds"))
	assert.Equal(t, "pistols_ids", ToSnakeCase("PistolsIDs"))
}
