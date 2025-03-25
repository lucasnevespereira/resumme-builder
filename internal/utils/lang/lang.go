package lang

import (
	"path/filepath"
	"resumme-builder/internal/utils/json"
	"resumme-builder/internal/utils/logger"

	"golang.org/x/text/language"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var (
	bundle             *i18n.Bundle
	localesDir         string
	supportedLanguages = []string{"en", "fr"}
)

func Init(uiDir string) {
	localesDir = filepath.Join(uiDir, "locales")

	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	for _, lang := range supportedLanguages {
		translation, err := loadTranslation(lang)
		if err != nil {
			logger.Log.Error(err)
			continue
		}
		err = bundle.AddMessages(language.Make(lang), translation.Messages...)
		if err != nil {
			logger.Log.Error(err)
			continue
		}
	}
}

func loadTranslation(lang string) (*i18n.MessageFile, error) {
	translationFile := filepath.Join(localesDir, lang+".json")
	return bundle.LoadMessageFile(translationFile)
}

func Translate(lang string, messageID string) string {
	return i18n.NewLocalizer(bundle, lang).
		MustLocalize(&i18n.LocalizeConfig{
			MessageID: messageID,
		})
}
