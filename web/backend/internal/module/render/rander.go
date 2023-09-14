package render

import (
	"envoy-go-fliter-hub/internal/module/parse"
)

func (r render) Render(metadata []parse.Metadata) error {
	details, list, err := r.renderIntoStruct(metadata)
	if err != nil {
		return err
	}
	err = r.writeToFile(details, list)
	if err != nil {
		return err
	}
	return nil
}
