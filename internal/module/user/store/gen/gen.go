package main

import (
	"envoy-golang-filter-hub/internal/module/user/store/model"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "internal/module/user/store/dao",
		//WithUnitTest:     true,
		Mode:             gen.WithoutContext /*| gen.WithDefaultQuery*/ | gen.WithQueryInterface, // generate mode
		FieldWithTypeTag: true,
		// generate field with type tag
	})

	g.ApplyBasic(&model.User{})
	g.Execute()
}
