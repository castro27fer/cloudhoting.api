package translations

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator"
)

// use a single instance , it caches struct info

var uni *ut.UniversalTranslator
var trans ut.Translator

type translation struct {
	tag             string
	translation     string
	override        bool
	customRegisFunc validator.RegisterTranslationsFunc
	customTransFunc validator.TranslationFunc
}

var translations_ES = []translation{
	{
		tag:         "required",
		translation: "El campo {0} es requerido",
		override:    false,
	},
}

var translations_EN = []translation{
	{
		tag:         "required",
		translation: "The field {0} is required",
		override:    false,
	},
}

type Languages struct {
	es []translation
	en []translation
}

var languages = Languages{
	es: translations_ES,
	en: translations_EN,
}


//var variable string ="una variable cadena" // declaracion de una variable....
//var puntero *string // declaración de una variable puntero vacia ....
// puntero = &myValue // asignar valor a un puntero......
// var puntero2 = &puntero // obtener la dirección en memoria de un puntero...
// var valorPuntero = *puntero // obtiene el valor al q apunta el puntero....

func Load_translates(language string) (err error) {

	trans, _ = uni.GetTranslator(language)

	for _, t := range languages.[language] {

		if err = trans.Add(t.tag, t.translation, t.override); err != nil {
			return
		}
	}
	return
}

func Init_translate_default() ut.Translator {

	//Instanciar locales.Translator
	ES := es.New() // en español
	EN := en.New() // en ingles

	//Instanciar un *ut.UniversalTranslator
	uni = ut.New(EN, EN, ES)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ = uni.GetTranslator("en")

	return trans
}

func Change_translate(translator string, v *validator.Validate) ut.Translator {
	//change translate
	// trans, _ = uni.GetTranslator(translator)
	// es_translations.RegisterDefaultTranslations(v, trans)
	return trans
}

func Get_translator() ut.Translator {
	return trans
}
