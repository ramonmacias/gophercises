package urlshort

import (
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
	return func(resp http.ResponseWriter, req *http.Request) {
		val, exist := pathsToUrls[req.URL.Path]
		if exist {
			http.Redirect(resp, req, val, http.StatusMovedPermanently)
		} else {
			fallback.ServeHTTP(resp, req)
		}
	}
}

// RedirectInfo is a struct used for unmarshall info from the YAML file
type RedirectInfo struct {
	Path string
	Url  string
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
	parsedYml, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	paths := buildMap(parsedYml)
	return MapHandler(paths, fallback), nil
}

// parseYAML function receive a yaml in form of array of bytes and return
// an array of objects RedirectInfo
func parseYAML(yml []byte) (redirectInfo []RedirectInfo, err error) {
	err = yaml.Unmarshal(yml, &redirectInfo)
	return redirectInfo, err
}

// buildMap Converts an slice of RedirectInfo objects into a map[string]string
func buildMap(redirectInfo []RedirectInfo) (paths map[string]string) {
	paths = make(map[string]string)
	for _, value := range redirectInfo {
		paths[value.Path] = value.Url
	}
	return paths
}
