package tmpl

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"testing"
	"text/template"
)

type Person struct {
	Name string
	Age int
	Boy bool
}

func TestRenderTmplByStruct(t *testing.T)  {
	// 实例化结构体
	foo := Person{
		Name: "foo",
		Age: 18,
		Boy: false,
	}
	// 定义模版文本
	const text = `
My name is {{.Name}},
I'm {{.Age}} years old,
I'm a {{if .Boy}}boy{{else}}girl{{end}}.
`

	// 根据指定模版文本生成handler
	tmpl, err := template.New("").Parse(text)
	if err != nil {
		log.Fatalln(err)
	}

	// 模版渲染，并写入文件
	f, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		log.Fatalln(err)
	}
	if err := tmpl.Execute(f, foo); err != nil {
		log.Fatalln(err)
	}

	// 模版渲染，并赋值给变量
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, foo); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(buf.String())

	// 模版渲染，并输出到屏幕标准输出
	if err := tmpl.Execute(os.Stdout, foo); err != nil {
		log.Fatalln(err)
	}
}

func TestRenderTmplByMap(t *testing.T)  {
	tpl := `foo={{.foo}}
bar={{.bar}}
baz={{.baz}}
`

	m := make(map[string]interface{})
	m["foo"] = 1
	m["bar"] = "BAR"
	m["baz"] = true

	tmpl, err := template.New("test").Parse(tpl)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(os.Stdout, m)
	if err != nil {
		log.Fatal(err)
	}
}

func TestRenderAndUnmarshal(t *testing.T)  {
	tpl := `
metadata:
  name: {{.ip}}
  labels:
    machine.infra.tce.io/cluster: global
spec:
  initialTaints: []
  initialAnnotations: {}
`

	m := make(map[string]interface{})
	m["ip"] = "192.168.10.87"

	tmpl, err := template.New("test").Parse(tpl)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(os.Stdout, m)
	if err != nil {
		log.Fatal(err)
	}

	// 模版渲染，并赋值给变量
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, m); err != nil {
		log.Fatalln(err)
	}

	bytes := buf.Bytes()

	out := make(map[string]interface{})
	if err = yaml.Unmarshal(bytes, &out); err != nil {
		log.Fatalln(err)
	}

	t.Logf("unmarshal result: %v", out)
}