package lang

import (
	"path/filepath"
	"resumme-builder/internal/utils/json"
	"resumme-builder/internal/utils/logger"
	"strings"

	"golang.org/x/text/language"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

const localesDir = "internal/utils/lang/locales"

var bundle *i18n.Bundle

var supportedLanguages = []string{"en_US", "fr_FR"}

func init() {
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
	translationFile := filepath.Join(localesDir, strings.ToLower(lang)+".json")
	return bundle.LoadMessageFile(translationFile)
}

func Translate(lang string, messageID string) string {
	return i18n.NewLocalizer(bundle, lang).
		MustLocalize(&i18n.LocalizeConfig{
			MessageID: messageID,
		})
}
