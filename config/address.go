package config

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
)

type EmailAddress struct {
	Address string
}

func NewEmailAddressList(fileName string) ([]*EmailAddress, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("打开邮件地址文件失败,error:%v", err.Error()))
	}
	var list []*EmailAddress
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		email := &EmailAddress{record[0]}
		list = append(list, email)
	}
	return list, nil
}
