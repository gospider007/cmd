package main

import (
	"testing"

	"github.com/gospider007/cmd"
)

func TestJs(t *testing.T) {
	script := `
	function sign(val,val2){
		return {"signval":val,"signval2":val2}
	}
    function sign2(val,val2){
		return {"sign2val":val,"sign2val2":val2}
	}
	`
	jsCli, err := cmd.NewJsClient(nil)
	if err != nil {
		t.Fatal(err)
	}
	if err = jsCli.Exec(script); err != nil {
		t.Fatal(err)
	}
	rs, err := jsCli.Call("sign", 1, 2)
	if err != nil {
		t.Fatal(err)
	}
	if rs.Get("signval").Int() != 1 || rs.Get("signval2").Int() != 2 {
		t.Fatal("sign error")
	}
	rs, err = jsCli.Call("sign2", 1, 2)
	if err != nil {
		t.Fatal(err)
	}
	if rs.Get("sign2val").Int() != 1 || rs.Get("sign2val2").Int() != 2 {
		t.Fatal("sign error")
	}
}

func TestPy(t *testing.T) {
	script := `def sign(val,val2):
	return {"val":val,"val2":val2}`
	pyCli, err := cmd.NewPyClient(nil)
	if err != nil {
		t.Fatal(err)
	}
	if err = pyCli.Exec(script); err != nil {
		t.Fatal(err)
	}
	rs, err := pyCli.Call("sign", 1, 2)
	if err != nil {
		t.Fatal(err)
	}
	if rs.Get("val").Int() != 1 || rs.Get("val2").Int() != 2 {
		t.Fatal("sign error")
	}
}
