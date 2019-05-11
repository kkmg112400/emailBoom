package main

import (
	"emailBoom/config"
	"flag"
	"fmt"
	"gopkg.in/gomail.v2"
	"sync"
	"time"
)

var (
	times    = flag.Int("times", 3, "攻击次数")
	groupNum = flag.Int("group_num", 3, "地址分组数量")
	title    = flag.String("title", "hello", "邮件标题")
	body     = flag.String("body", "good luck", "邮件内容")
	account  = flag.String("account", "", "发送邮件账号")
	password = flag.String("password", "", "发送邮件密码")
	port     = flag.Int("port", 465, "SMTP端口号")
	host     = flag.String("host", "", "SMTP服务器地址")
)

func init() {
	flag.Parse()
}

func main() {
	list, err := config.NewEmailAddressList("./address.csv")
	if err != nil {
		fmt.Printf("error:%s", err.Error())
		return
	}
	if len(list) == 0 {
		fmt.Printf("error:接收邮件地址列表为空")
		return
	}
	x := make([]int, *times)
	d := gomail.NewDialer(*host, *port, *account, *password)
	if *host == "" || *account == "" || *password == "" {
		fmt.Printf("error:参数错误")
		return
	}
	addresses := make([]string, 0)
	addressList := make([][]string, 0)
	fmt.Printf("===== 开始分组\n")
	for _, address := range list {
		addresses = append(addresses, address.Address)
		if len(addresses) >= *groupNum {
			addressList = append(addressList, addresses)
			fmt.Printf("===== 第%d组,地址:%v\n", len(addressList), addresses)
			addresses = make([]string, 0)
		}
	}
	if len(addresses) > 0 {
		addressList = append(addressList, addresses)
		fmt.Printf("===== 第%d组,地址:%v\n", len(addressList), addresses)
		addresses = make([]string, 0)
	}
	fmt.Printf("===== 开始发送\n")
	startTime := time.Now().UnixNano()
	var wg sync.WaitGroup
	for i, addresses := range addressList {
		wg.Add(1)
		go func(i int, addresses []string) {
			m := gomail.NewMessage()
			m.SetHeader("From", *account)
			m.SetHeader("To", addresses...)
			m.SetHeader("Subject", *title)
			m.SetBody("text/plain", *body)
			for j := range x {
				wg.Add(1)
				go func(i, j int) {
					if err := d.DialAndSend(m); err != nil {
						fmt.Printf("error:第%d组的第%d次发送失败:%s\n", i+1, j+1, err.Error())
					} else {
						fmt.Printf("===== 第%d组的第%d次发送成功\n", i+1, j+1)
					}
					wg.Done()
				}(i, j)
			}
			wg.Done()
		}(i, addresses)
	}
	wg.Wait()
	endTime := time.Now().UnixNano()
	fmt.Printf("===== 发送完成 耗时:%dms\n", (endTime-startTime)/1e6)
}
