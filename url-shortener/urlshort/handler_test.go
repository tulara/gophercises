package urlshort

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestYamlHandler(t *testing.T) {
	yaml := `
- path: /urlshort
  url: https://github.com/thisdoesntexist
`
	t.Run("yaml is parsed into a redirect map that can be passed to another handler",
		func(t *testing.T) {
			

			YAMLHandler([]byte(yaml), nil)

	})
}