package docs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwaggerInfo(t *testing.T) {
	assert := assert.New(t)

	// Check if SwaggerInfo fields are initialized correctly
	assert.Equal("", SwaggerInfo.Version, "Version should be empty")
	assert.Equal("", SwaggerInfo.Host, "Host should be empty")
	assert.Equal("", SwaggerInfo.BasePath, "BasePath should be empty")
	assert.Equal([]string{}, SwaggerInfo.Schemes, "Schemes should be empty")
	assert.Equal("", SwaggerInfo.Title, "Title should be empty")
	assert.Equal("", SwaggerInfo.Description, "Description should be empty")
	assert.Equal("swagger", SwaggerInfo.InfoInstanceName, "InfoInstanceName should be 'swagger'")
	assert.Equal(
		docTemplate,
		SwaggerInfo.SwaggerTemplate,
		"SwaggerTemplate should match the docTemplate",
	)
	assert.Equal("{{", SwaggerInfo.LeftDelim, "LeftDelim should be '{{'")
	assert.Equal("}}", SwaggerInfo.RightDelim, "RightDelim should be '}}'")
}
