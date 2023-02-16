package validations

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func init() {

	validate = validator.New()
	en := en.New()
	uni = ut.New(en, en)
	trans, _ = uni.GetTranslator("en")

	en_translations.RegisterDefaultTranslations(validate, trans)

	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {

		return ut.Add("required", "{0} must have a value!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})

}

func Validate(paylaod interface{}) []string {

	err := validate.Struct(paylaod)
	var errorsArray []string

	if err != nil {

		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			// can translate each error one at a time.
			errorsArray = append(errorsArray, e.Translate(trans))

		}
	}

	return errorsArray

}
