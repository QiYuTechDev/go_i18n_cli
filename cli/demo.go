package cli

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"os"
)

// GetLanguage 获取运行此任务使用的语言
func GetLanguage() language.Tag {
	for _, key := range []string{"LANGUAGE", "LC_ALL", "LANG"} {
		if tag, err := language.Parse(os.Getenv(key)); err == nil {
			return tag
		}
	}
	return language.English
}

// GetPrinter 获取 i18n 翻译 Printer
func GetPrinter() *message.Printer {
	return message.NewPrinter(GetLanguage())
}

func I18nHello(name string) {
	i18n := GetPrinter()
	println(i18n.Sprintf("Hello: %s", name))
}
