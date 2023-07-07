package main

import (
	"envoy-golang-filter-hub/internal/module/plugin/store/model"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "internal/module/plugin/store/dao",
		//WithUnitTest:     true,
		Mode:             gen.WithoutContext /*| gen.WithDefaultQuery*/ | gen.WithQueryInterface, // generate mode
		FieldWithTypeTag: true,
		// generate field with type tag
	})

	g.ApplyBasic(&model.Plugin{}, &model.Version{}, &model.Tag{})
	g.Execute()
}
