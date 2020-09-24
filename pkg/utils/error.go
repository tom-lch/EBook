package utils

import "errors"

func ErrorConn() error {
	return errors.New("连接失败")
}