package jwtAuth

import "time"

// 还有些问题 稍后更改
var Blacklist = make(map[string]int64)

func Isinblacklist(token string) bool {
	if _, ok := Blacklist[token]; !ok {
		return false
	}
	return true
}

func ClearList() {
	for {
		for k, v := range Blacklist {
			if v < time.Now().Unix() {
				delete(Blacklist, k)
			}
		}
		time.Sleep(60 * time.Second)
	}
}
