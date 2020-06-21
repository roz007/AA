package helpers

import (
	"regexp"

	"github.com/sirupsen/logrus"
)

func Extractit(htmlresponse string, pattern string, cleaner string, elem int) string {
	cleanerObj, err := regexp.Compile(cleaner)
	if err != nil {
		logrus.Error(err)
	}
	regexobj, err := regexp.Compile(pattern)
	if err != nil {
		logrus.Error(err)
	}
	value := regexobj.FindAllString(htmlresponse, -1)[elem]
	clean_value := cleanerObj.ReplaceAllString(value, "")
	return clean_value

}
