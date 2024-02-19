package translations

import (
	"fmt"

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
		tag:         "Email",
		translation: "Correo",
		override:    false,
	},
	{
		tag:         "Password",
		translation: "Contraseña",
		override:    false,
	},
	{
		tag:         "required",
		translation: "El campo {0} es requerido",
		override:    false,
	},
	{
		tag:         "min",
		translation: "El campo {0} require minimo {1} caracteres",
		override:    false,
	},
}

var translations_EN = []translation{
	{
		tag:         "Email",
		translation: "Email",
		override:    false,
	},
	{
		tag:         "Password",
		translation: "Password",
		override:    false,
	},
	{
		tag:         "required",
		translation: "The field {0} is required",
		override:    false,
	},
	{
		tag:         "min",
		translation: "The field {0} require minimum {1} characters",
		override:    false,
	},
}

type Languages struct {
	es []translation `json:"es"`
	en []translation `json:"en"`
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

// load all languages into the traslator
func Load_languages() (err error) {

	//translations in english.....
	trans, _ = uni.GetTranslator("en")
	if err = Load_traslator(languages.en, trans); err != nil {
		return err
	}

	//translations in spanish
	trans, _ = uni.GetTranslator("es")
	if err = Load_traslator(languages.es, trans); err != nil {
		return err
	}

	//translator by default...
	trans, _ = uni.GetTranslator("es")

	return nil
}

// translator = traductor
// language = list of traslations
// load all translations into the translator
func Load_traslator(language []translation, translator ut.Translator) (err error) {

	for _, traslation := range language {

		//add translation into translator
		if err = translator.Add(traslation.tag, traslation.translation, traslation.override); err != nil {
			fmt.Printf("%v", err)
			return err
		}
	}
	return nil
}

func Init_translate_default() ut.Translator {

	//create new instances of translators
	ES := es.New() // in spanish
	EN := en.New() // in english

	//create new instance of universal translator
	uni = ut.New(EN, EN, ES)

	//get translator in english
	trans, _ = uni.GetTranslator("en")

	return trans
}

func Change_translate(translator string, v *validator.Validate) ut.Translator {
	//change translate
	// trans, _ = uni.GetTranslator(translator)
	// es_translations.RegisterDefaultTranslations(v, trans)
	return trans
}

// change translator
func Change_translator(language string) ut.Translator {

	trans, _ = uni.GetTranslator(language)
	return trans
}

// get the current translator
func Get_translator() ut.Translator {
	return trans
}
