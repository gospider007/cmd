# 功能概要
* 执行命令行命令
* 没有内存泄露的执行cmd 命令
* js 解析器，配合node ,调用js代码，使用管道
* python 解析器，配合python ,调用python代码，使用管道

## 执行js 代码示例
~~~go
package main

import (
	"log"

	"gitee.com/baixudong/cmd"
)
func TestJs() {
	script := `
	function sign(val,val2){
		return {"signval":val,"signval2":val2}
	}
    function sign2(val,val2){
		return {"sign2val":val,"sign2val2":val2}
	}
	`
	jsCli, err := cmd.NewJsClient(nil, cmd.JsClientOption{
		Script: script,
		Names:  []string{"sign", "sign2"},
	})
	if err != nil {
		log.Fatal(err)
	}
	rs, err := jsCli.Call("sign", 1, 2)
	if err != nil {
		log.Fatal(err)
	}
	if rs.Get("signval").Int() != 1 || rs.Get("signval2").Int() != 2 {
		log.Fatal("sign error")
	}
	rs, err = jsCli.Call("sign2", 1, 2)
	if err != nil {
		log.Fatal(err)
	}
	if rs.Get("sign2val").Int() != 1 || rs.Get("sign2val2").Int() != 2 {
		log.Fatal("sign error")
	}
}
~~~
## 执行python代码示例
~~~go
func TestPy() {
	script := `def sign(val,val2):
	return {"val":val,"val2":val2}`
	pyCli, err := cmd.NewPyClient(nil, cmd.PyClientOption{
		Script: script,
		Names:  []string{"sign"},
	})
	if err != nil {
		log.Fatal(err)
	}
	rs, err := pyCli.Call("sign", 1, 2)
	if err != nil {
		log.Fatal(err)
	}
	if rs.Get("val").Int() != 1 || rs.Get("val2").Int() != 2 {
		log.Fatal("sign error")
	}
}
~~~