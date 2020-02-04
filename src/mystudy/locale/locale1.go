package locale1

import (
	"fmt"
	"time"
)

var locales map[string]map[string]string

func locale1() {
	locales = make(map[string]map[string]string, 2)
	en := make(map[string]string, 10)
	en["pea"] = "pea"
	en["bean"] = "bean"
	locales["en"] = en
	cn := make(map[string]string, 10)
	cn["pea"] = "豌豆"
	cn["bean"] = "毛豆"
	locales["zh-CN"] = cn
	lang := "zh-CN"
	fmt.Println(msg(lang, "pea"))
	fmt.Println(msg(lang, "bean"))

	en["time_zone"] = "America/Chicago"
	cn["time_zone"] = "Asia/Shanghai"

	loc, _ := time.LoadLocation(msg(lang, "time_zone"))
	t := time.Now()
	t = t.In(loc)
	fmt.Println(t.Format(time.RFC3339))

	//en["date_format"] = "%Y-%m-%d %H:%M:%S"
	//cn["date_format"] = "%Y年%m月%d日 %H时%M分%S秒"
	//
	////fmt.Println(date(msg(lang,"date_format"),t))
	//
	//en["money"] = "USD %d"
	//cn["money"] = "￥%d元"
	//
	//fmt.Println(money_format(msg(lang, "date_format"), 100))
}

func money_format(fomate string, money int64) string {
	return fmt.Sprintf(fomate, money)
}

func msg(locale, key string) string {
	if v, ok := locales[locale]; ok {
		if v2, ok := v[key]; ok {
			return v2
		}
	}
	return ""
}

//func date(fomate string,t time.Time) string{
//	year, month, day := t.Date()
//	hour, min, sec :=t.Clock()
//	//解析相应的%Y %m %d %H %M %S然后返回信息
//	//%Y 替换成2012
//	//%m 替换成10
//	//%d 替换成24
//}
/*
func f() {
	Tr := i18n.NewLocale()
	Tr.LoadPath("config/locales")

	fmt.Println(Tr.Translate("submit"))
	//输出Submit
	Tr.SetLocale("zh")
	fmt.Println(Tr.Translate("submit"))
}
*/

//TODO i18n
//加载默认配置文件，这些文件都放在go-i18n/locales下面

//文件命名zh.json、en.json、en-US.json等，可以不断的扩展支持更多的语言
/*
func (il *IL) loadDefaultTranslations(dirPath string) error {
	dir, err := os.Open(dirPath)
	if err != nil {
		return err
	}
	defer dir.Close()

	names, err := dir.Readdirnames(-1)
	if err != nil {
		return err
	}

	for _, name := range names {
		fullPath := path.Join(dirPath, name)

		fi, err := os.Stat(fullPath)
		if err != nil {
			return err
		}

		if fi.IsDir() {
			if err := il.loadTranslations(fullPath); err != nil {
				return err
			}
		} else if locale := il.matchingLocaleFromFileName(name); locale != "" {
			file, err := os.Open(fullPath)
			if err != nil {
				return err
			}
			defer file.Close()

			if err := il.loadTranslation(file, locale); err != nil {
				return err
			}
		}
	}

	return nil
}
*/

func Locale()  {
	locale1()
}