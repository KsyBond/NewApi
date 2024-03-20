package render

import (
	"encoding/json"
	"errors"
	"io"
)

func DecodeJSON(r io.Reader, v any) error {
	const op = "utils.render.decodeJSON"
	b, err := io.ReadAll(r)
	if err != nil {
		return errors.New(op + "- failed to decode json")
	}

	err = json.Unmarshal(b, &v)
	if err != nil {
		return errors.New(op + "- failed to unmarshall json")
	}
	return nil
}
