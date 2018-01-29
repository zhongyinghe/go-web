1、预防跨站脚本(防xss攻击)
```
它的目的是什么?
答:就是想在页面执行它想执行的js

执行js有什么作用?
答:如session劫持.

什么是session劫持?
答:就是通过js获取用户的sessionId

如何通过js获取用户的sessionId?
答:通过执行js,获取用户的cookie,然后把该cookie发送到一个链接上进行保存.
  具体可以参考:http://www.cnblogs.com/dolphinX/p/3403027.html

如何防范xss攻击?
答:首先,过滤用户输入;其次, 设置sessionId的cookie为HttpOnly，使客户端无法获取
```
2、防CSRF攻击
```
什么是CSRF攻击?
答:攻击者盗用了你的身份，以你的名义发送恶意请求.具体过程请参考:https://www.cnblogs.com/hyddd/archive/2009/04/09/1432744.html

达成CSRF攻击的条件是什么?
答:1.登录受信任网站A，并在本地生成Cookie。
   2.在不登出A的情况下，访问危险网站B。

如何防范CSRF攻击?
答:1)在修改操作尽量使用post请求
   2)金额操作要使用短信验证码
   3)增加伪随机数的判断(hash值)
```
3、进行权限判断
```
如何进行权限判断呢?
答:用户查询、编辑、删除时要注意归属检测
```
4、防止多次提交
```
如何防止多次提交表单?
答:前端处理。点击提交按钮，让它不能再次点击
   后端处理。在表单中设置token,通过后端token判断来确定是否进行处理
```
5、过滤数据
```
如何过滤数据呢?
答:1)转义(大文本的要转义)
   2)正则匹配数据.(不要试图纠正用户的输入)(一般字段的要强制匹配)

在哪一端进行判断?
答:前端和后端都要判断
```
6、防SQL注入
```
如何防止sql注入?
答:1)对要求的数据类型进行转化或者进行匹配判断(如:strcov;针对number类型)
   2)对字符串进行转义操作(html/template中的template.HTMLEscapeString(str string))
   3)使用占位符的sql语句
   4)使用SQL注入检测工具进行检测(如:sqlmap、SQLninja)
   5)抽离出中间层
   6)匹配过滤
```
过滤函数:
```go
func main() {
    str := `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`
    re, err := regexp.Compile(str)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    aa := "存在 select"
    fmt.Println(re.MatchString(aa))  //打印出true。
}
```
//参考:https://www.golangtc.com/t/5410293b320b527a3b000179<br>
如何理解sql注入呢?<br>
答:https://www.cnblogs.com/pursuitofacm/p/6706961.html<br>
	https://www.cnblogs.com/ichunqiu/p/5749347.html<br>
