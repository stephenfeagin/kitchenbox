package ldjson

import (
	"encoding/json"
	"errors"
	"strings"
)

// unmarshalToStringSlice converts a json.RawMessage, which could be either a string or an array of
// strings, possibly empty, into a Go []string.
func unmarshalToStringSlice(raw json.RawMessage) ([]string, error) {
	var syntaxError *json.SyntaxError
	// try to unmarshal into a string slice
	var slc []string
	if err := json.Unmarshal(raw, &slc); err == nil {
		return slc, nil
	} else if errors.Is(err, syntaxError) {
		return nil, err
	}
	// if not successful, cast to string and separate by commas
	str := string(raw)
	if str == "" || str == `""` {
		return slc, nil
	}
	slc = strings.Split(str, ",")
	return slc, nil
}
