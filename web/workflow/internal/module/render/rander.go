package render

import (
	"encoding/json"
	"envoy-go-fliter-hub/internal/model"
	"fmt"
)

func (r render) Render(metadata []model.Metadata) error {
	details, list, err := r.renderIntoStruct(metadata)
	if err != nil {
		return err
	}

	jsonData, _ := json.MarshalIndent(details, "", "  ")
	fmt.Println(string(jsonData))

	err = r.writeToFile(details, list)
	if err != nil {
		return err
	}
	return nil
}
