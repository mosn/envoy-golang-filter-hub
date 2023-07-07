package validate

import (
	"github.com/go-playground/validator/v10"
)

var (
	Validate = validator.New()
	//Validation = make(map[string]validator.Func)
)

//func init() {
//	for k, v := range Validation {
//		utils.PanicIfErr(Validate.RegisterValidation(k, v))
//	}
//}
