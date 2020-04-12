package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
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
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathsUrls, err := ParseYaml(yml)
	if err != nil {
		return nil, err
	}
	pathsToURL := buildMap(pathsUrls)
	return MapHandler(pathsToURL, fallback), nil
}

func JSONHandler(json []byte, fallback http.Handler) (http.HandlerFunc, error) {

}

func buildMap(pathsUrls []PathURL) map[string]string {
	pathsToURL := make(map[string]string)
	for _, path := range pathsUrls {
		pathsToURL[path.Path] = path.url
	}
	return pathsToURL
}

// ParseYaml YAML to PathURL
func ParseYaml(data []byte) ([]PathURL, error) {
	var pathsUrls []PathURL
	err := yaml.Unmarshal(data, &pathsUrls)
	if err != nil {
		return nil, err
	}
	return pathsUrls, nil
}

// ParseJson Json to PathURL
func ParseJson(data []byte) ([]PathURL, error) {
	var pathsUrls []PathURL
	err := json.Unmarshal(data, &pathsUrls)
	if err != nil {
		return nil, err
	}
	return pathsUrls, nil
}

// PathURL Path URL
type PathURL struct {
	Path string `yaml:"path,omitempty"`
	url  string `yaml:"url,omitempty"`
}
