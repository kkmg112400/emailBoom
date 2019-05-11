# emailBoom
go语言邮箱轰炸客户端
##### address.csv 配置轰炸邮箱列表
##### 编译说明
###### 该项目使用govendor管理第三方库;
###### 执行`go build`之前先执行`govendor sync`同步第三方库
##### Usage
    Usage of ./emailBoom:
    -account string
          发送邮件账号
    -body string
          邮件内容 (default "good luck")
    -group_num int
          地址分组数量 (default 3)
    -host string
          SMTP服务器地址
    -password string
          发送邮件密码
    -port int
          SMTP端口号 (default 465)
    -times int
          攻击次数 (default 3)
    -title string
          邮件标题 (default "hello")
  ![emailBoom.png](http://gwjyhs.com/t6/702/1557584253x2890174417.png)
