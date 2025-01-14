package utils

import (
	"fmt"
	"html"
	"regexp"
	"strings"
)

func SensitiveInfoCollect(Url string, Content string) {
	infoMap := map[string]string{
		"Chinese Mobile Number": `[^\d]((?:(?:\+|00)86)?1(?:(?:3[\d])|(?:4[5-79])|(?:5[0-35-9])|(?:6[5-7])|(?:7[0-8])|(?:8[\d])|(?:9[189]))\d{8})[^\d]`,
		"Internal IP Address":   `[^0-9]((127\.0\.0\.1)|(10\.\d{1,3}\.\d{1,3}\.\d{1,3})|(172\.((1[6-9])|(2\d)|(3[01]))\.\d{1,3}\.\d{1,3})|(192\.168\.\d{1,3}\.\d{1,3}))`,
		"JSON Web Token":        `(eyJ[A-Za-z0-9_-]{10,}\.[A-Za-z0-9._-]{10,}|eyJ[A-Za-z0-9_\/+-]{10,}\.[A-Za-z0-9._\/+-]{10,})`,
		"accesskey/accessid":    `(?i)(access([_ ]?(key|id|secret)){1,2}[\w\s=":'\.]*?([0-9a-zA-Z]{10,64}))`,
		"password":              `(?i)(password[\w\s=":'\.]{1,25}[0-9a-zA-Z_@!-#\$]{1,64})`,
		"Password":              `(?i)([0-9a-zA-Z_@!-#\$]{1,64}[\w\s=":'\.]{1,25}password)`,
	}
	for key := range infoMap {
		reg := regexp.MustCompile(infoMap[key])
		res := reg.FindAllStringSubmatch(html.UnescapeString(Content), -1)
		if len(res) > 0 {
			fmt.Printf("->[*] [%s] %s\n", Url, key)
			mylist := []string{}
			for _, tmp := range res {
				mylist = append(mylist, tmp[1])
			}
			unDupList := removeDuplicatesString(mylist)
			fmt.Println(strings.Join(unDupList, "\n"))
		}
	}
}
