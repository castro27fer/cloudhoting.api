package translations

import (
	"fmt"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator"
)

type translation struct {
	tag             string
	translation     string
	override        bool
	customRegisFunc validator.RegisterTranslationsFunc
	customTransFunc validator.TranslationFunc
}

var ES = []translation{
	{
		tag:         "required",
		translation: "El campo {0} es requerido",
		override:    false,
	},
}

func Location_ES() []translation {
	return ES
}

// func RegisterDefaultTranslations(v *validator.Validate, trans ut.Translator) (err error) {

// 	for _, t := range ES {

// 		if t.customTransFunc != nil && t.customRegisFunc != nil {
// 			err = v.RegisterTranslation(t.tag, trans, t.customRegisFunc, t.customTransFunc)
// 		} else if t.customTransFunc != nil && t.customRegisFunc == nil {
// 			err = v.RegisterTranslation(t.tag, trans, registrationFunc(t.tag, t.translation, t.override), t.customTransFunc)
// 		} else if t.customTransFunc == nil && t.customRegisFunc != nil {
// 			err = v.RegisterTranslation(t.tag, trans, t.customRegisFunc, translateFunc)
// 		} else {
// 			err = v.RegisterTranslation(t.tag, trans, registrationFunc(t.tag, t.translation, t.override), translateFunc)
// 		}

// 		if err != nil {
// 			return
// 		}
// 	}

// 	return
// }

func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		fmt.Printf("warning: error translating FieldError: %#v", fe)
		return fe.(error).Error()
	}

	return t
}
