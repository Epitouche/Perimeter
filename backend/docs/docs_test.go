package docs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"area/docs"
)

func TestSwaggerInfo(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)

	// Check if SwaggerInfo fields are initialized correctly
	assert.Equal("", docs.SwaggerInfo.Version, "Version should be empty")
	assert.Equal("", docs.SwaggerInfo.Host, "Host should be empty")
	assert.Equal("", docs.SwaggerInfo.BasePath, "BasePath should be empty")
	assert.Equal([]string{}, docs.SwaggerInfo.Schemes, "Schemes should be empty")
	assert.Equal("", docs.SwaggerInfo.Title, "Title should be empty")
	assert.Equal("", docs.SwaggerInfo.Description, "Description should be empty")
	assert.Equal("swagger", docs.SwaggerInfo.InfoInstanceName, "InfoInstanceName should be 'swagger'")
	// assert.Equal(
	// 	docTemplate,
	// 	docs.SwaggerInfo.SwaggerTemplate,
	// 	"SwaggerTemplate should match the docTemplate",
	// )
	assert.Equal("{{", docs.SwaggerInfo.LeftDelim, "LeftDelim should be '{{'")
	assert.Equal("}}", docs.SwaggerInfo.RightDelim, "RightDelim should be '}}'")
}
