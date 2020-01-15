package urlshort

import (
	"fmt"
	"net/http"
	"gopkg.in/yaml.v2"
)

// MapHandler redirects to longer url if we find one, 
// otherwise passes request on to fallback handler
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request){
		//html.EscapeString on request path?
		//does this include query strings? 
		if url, ok := pathsToUrls[req.URL.Path]; ok {
			http.Redirect(w, req, url, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, req)
	}
}

type Redirect struct {
	Path string
	URL string
}
// Redirects see if the marshaller needs this exported - presumably not?
type Redirects []Redirect

func parseYaml(yml []byte) (Redirects, error) {
	p := Redirects{}

	err := yaml.Unmarshal(yml, &p)
	if err != nil {
		// replace this with Injected logger?
		fmt.Println("error: ", err)
	}
	return p, err
}

func buildMap(parsedYml Redirects) map[string]string {
	m := make(map[string]string)

	for _, redirect := range parsedYml {
		m[redirect.Path] = redirect.URL
	}

	return m
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all relate to having
// invalid YAML data.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYml, err := parseYaml(yml)
	if err != nil {
		return nil, err
	}

	pathMap := buildMap(parsedYml)
	return MapHandler(pathMap, fallback), nil
}