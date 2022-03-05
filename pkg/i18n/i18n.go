package i18n

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	goI18n "github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *goI18n.Bundle
var localizer *goI18n.Localizer

func init() {
	bundle = goI18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	// check if file "locales/en.toml" exists
	if _, err := os.Stat("locales/en.toml"); err == nil {
		bundle.LoadMessageFile("locales/en.toml")
	}
}

// SetLocale sets the locale to use for the application.
func SetLocale(lang string) {
	if _, err := os.Stat("locales/" + lang + ".toml"); err == nil {
		bundle.LoadMessageFile("locales/" + lang + ".toml")
	} else {
		log.Printf("Locale file not found: locales/%s.toml", lang)
	}
}

// Localize returns the localized string for the given id.
func GetMessageWithKey(key string, templateData map[string]interface{}, pluralCount int) string {
	message := localizer.MustLocalize(&goI18n.LocalizeConfig{
		MessageID:    key,
		TemplateData: templateData,
	})
	return message
}

// Localize returns the localized string for the given message
func GenerateMessage(id, description, one, other string, data map[string]interface{}, pluralCount int) string {
	return localizer.MustLocalize(&goI18n.LocalizeConfig{
		DefaultMessage: &goI18n.Message{
			ID:          id,
			Description: description,
			One:         one,
			Other:       other,
		},
		TemplateData: data,
		PluralCount:  pluralCount,
	})
}
