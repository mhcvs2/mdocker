package utils

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strings"
	"text/template"
	myLogger "mdocker/logger"
	"mdocker/logger/colors"
	"encoding/json"
	"os/exec"
	"syscall"
	"mdocker/config"
)

func PreRun() {
	//appid := os.Getenv("APPID")
	//appsecret := os.Getenv("APPSCERET")
	//if appid == "" || appsecret == "" {
	//	PrintAndExit("ss", PreRunTemplate)
	//}
}

func GetAppName() string {
	return config.Conf.AppName
}

func GetAppDescription() string {
	return config.Conf.Description
}

// Go is a basic promise implementation: it wraps calls a function in a goroutine
// and returns a channel which will later return the function's return value.
func Go(f func() error) chan error {
	ch := make(chan error)
	go func() {
		ch <- f()
	}()
	return ch
}

// IsExist returns whether a file or directory exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// CloseFile attempts to close the passed file
// or panics with the actual error
func CloseFile(f *os.File) {
	err := f.Close()
	MustCheck(err)
}

// MustCheck panics when the error is not nil
func MustCheck(err error) {
	if err != nil {
		panic(err)
	}
}

// __FILE__ returns the file name in which the function was invoked
func FILE() string {
	_, file, _, _ := runtime.Caller(1)
	return file
}

// BeeFuncMap returns a FuncMap of functions used in different templates.
func BeeFuncMap() template.FuncMap {
	return template.FuncMap{
		"trim":       strings.TrimSpace,
		"bold":       colors.Bold,
		"headline":   colors.MagentaBold,
		"foldername": colors.RedBold,
		"endline":    EndLine,
		"tmpltostr":  TmplToString,
		"fen":        fenString,
		"red":        colors.Red,
		"fenint":     fenByint,
		"getmax":     getMaxLen,             //list max len
		"getmax2":     getMaxLen2,           // map max len, both key and value
		"fent":       fenTirtle,
		"shu":        shu,
		"getmaxm":    getMaxMultiple,            //get max len between Multiple string
		"getlen":        getLen,
		"maxvalue":   getMaxValue,
		"getmax3":    getMaxLen3,
		"getmax4":    getMaxLen4,
		"blue":       colors.Blue,
		"app":        GetAppName,
		"description": GetAppDescription,
	}
}

// TmplToString parses a text template and return the result as a string.
func TmplToString(tmpl string, data interface{}) string {
	t := template.New("tmpl").Funcs(BeeFuncMap())
	template.Must(t.Parse(tmpl))

	var doc bytes.Buffer
	err := t.Execute(&doc, data)
	MustCheck(err)

	return doc.String()
}

// EndLine returns the a newline escape character
func EndLine() string {
	return "\n"
}

func Tmpl(text string, data interface{}) {
	output := colors.NewColorWriter(os.Stderr)

	t := template.New("Usage").Funcs(BeeFuncMap())
	template.Must(t.Parse(text))

	err := t.Execute(output, data)
	if err != nil {
		myLogger.Log.Error(err.Error())
	}
}

func Tmpl2(text string, data interface{}) {
	f, _ := os.Create("/tmp/mdocker")
	output := colors.NewColorWriter(f)

	t := template.New("Usage").Funcs(BeeFuncMap())
	template.Must(t.Parse(text))

	err := t.Execute(output, data)
	if err != nil {
		myLogger.Log.Error(err.Error())
	}
	binary, lookErr := exec.LookPath("less")
	if lookErr != nil {
		panic(lookErr)
	}
	args := []string{"less", "/tmp/mdocker"}
	env := os.Environ()
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

func PrintErrorAndExit(message, errorTemplate string) {
	Tmpl(fmt.Sprintf(errorTemplate, message), nil)
	os.Exit(2)
}

func fenString(str string) string {
	r := "-"
	for i:=0; i<len([]rune(str));i++ {
		r += "-"
	}
	return r
}

func fenByint(l int) string {
	r := "-"
	for i:=0; i<l;i++ {
		r += "-"
	}
	return r
}

func fenTirtle(l int, s string) string {
	l = l - len(s)
	r := "-"
	for i:=0; i<l/2;i++ {
		r += "-"
	}
	r += colors.Bold(s)
	for i:=l/2; i<l;i++ {
		r += "-"
	}
	return r
}

func shu(l int, s string) string {
	l = l - len(s)
	r := " " + s
	for i:=0; i<l;i++ {
		r += " "
	}
	r += " "
	return r
}

func getMaxLen(strs []string) int {
	max := 0
	for _, str := range strs {
		if len(str) > max { max = len(str) }
	}
	return max
}

func getMaxValue(values ...int) int {
	max := 0
	for _, value := range values {
		if value > max { max = value }
	}
	return max
}

type LenRes struct {
	X int
	Y int
	Sum int
}

func getMaxLen2(m map[string]string) LenRes {
	maxX := 0
	maxY := 0
	for k, v := range m {
		if len(k) > maxX { maxX = len(k) }
		if len(v) > maxY { maxY = len(v) }
	}
	sum := maxX + maxY
	lenres := LenRes{maxX, maxY, sum}
	return lenres
}

func getMaxLen3(m map[string]string, keyName, valueName string) LenRes {
	maxX := len(keyName)
	maxY := len(valueName)
	for k, v := range m {
		if len(k) > maxX { maxX = len(k) }
		if len(v) > maxY { maxY = len(v) }
	}
	sum := maxX + maxY
	lenres := LenRes{maxX, maxY, sum}
	return lenres
}

func getMaxLen4(m map[string][]string, keyName, valueName string) LenRes {
	maxX := len(keyName)
	maxY := len(valueName)
	for k, v := range m {
		if len(k) > maxX { maxX = len(k) }
		lv := getMaxLen(v)
		if lv > maxY { maxY = lv }
	}
	sum := maxX + maxY
	lenres := LenRes{maxX, maxY, sum}
	return lenres
}

func getMaxMultiple(s1 ...string) int {
	return getMaxLen(s1)
}

func getLen(s string) int {
	return len(s)
}

func PrintAndExit(message, template string) {
	Tmpl(template, message)
	os.Exit(0)
}

func Contains(s string, kvs []string) bool {
	for _, kv := range kvs {
		if !strings.Contains(strings.ToLower(s), strings.ToLower(strings.TrimSpace(kv))) { return false}
	}
	return true
}

func Listfilter(target []string, keywords string, conditionfun func(text string) bool) []string {
	res := []string{}
	kws := strings.Split(keywords, "&")
	for _, s := range target {
		if strings.TrimSpace(s) == "" { continue }
		if len(kws) > 0 && strings.TrimSpace(kws[0]) != "" && !Contains(s, kws){
			continue
		}
		if conditionfun(s) == true { continue }
		res = append(res, s)
	}
	return res
}

func ListToMap(ss []string, separator string) map[string]string {
	res := make(map[string]string)
	for _, s := range ss {
		kv := strings.Split(s, separator)
		if len(kv) == 2 {
			res[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
		}
	}
	return res
}

func InSlice(i string,s []string) bool {
	for _, ss := range s {
		if strings.ToLower(i) == strings.ToLower(ss) {
			return true
		}
	}
	return false
}

var SuccessTemplate = `{{ . | blue }}{{ endline }}`

var FailedTemplate = `{{ . | red }}{{ endline }}`

var PreRunTemplate = ` {{ "Lack env 'APPID' 'APPSCERET'" | headline }}
 For example:
	{{ "export APPID=test" | bold}}
	{{ "export APPSCERET=ee1d7336-73e3-452b-9e63-fdeaf2dccde6" | bold}}
`

type ListTemplateData struct {
	Tirtle string
	Data []string
}

var ListTemplate = `{{ $l := getmax .Data }}-{{fent $l .Tirtle}}--
{{ range $k,$v := .Data }}|{{ shu $l $v }}|
-{{ fenint $l}}--
{{ end }}`

type Map1TemplateData struct {
	Data map[string]string
}

var Map1Template = `{{ $m := getmax2 .Data }}-{{ fenint $m.Sum}}-----
{{ range $k,$v := .Data }}|{{ bold (shu $m.X $k) }}|{{ blue (shu $m.Y $v) }}|
-{{ fenint $m.Sum}}-----
{{ end }}`

type MapTemplateData struct {
	Tirtle string
	KeyName string
	ValueName string
	Data map[string]string
}

var MapTemplate = `{{ $m := getmax3 .Data $.KeyName $.ValueName }}--{{fent $m.Sum .Tirtle}}----
|{{ bold (shu $m.X $.KeyName) }}|{{ bold (shu $m.Y $.ValueName) }}|
-{{ fenint $m.Sum}}-----
{{ range $k,$v := .Data }}|{{ shu $m.X $k }}|{{ shu $m.Y $v }}|
-{{ fenint $m.Sum}}-----
{{ end }}`

type MapListTemplateData struct {
	KeyName string
	ValueName string
	Data map[string][]string
}

var MapListTemplate = `{{ $m := getmax4 .Data $.KeyName $.ValueName }}-{{ fenint $m.Sum}}-----
|{{ bold (shu $m.X $.KeyName) }}|{{ bold (shu $m.Y $.ValueName) }}|
-{{ fenint $m.Sum}}-----
{{ range $k,$vs := .Data }}{{ range $i,$v := $vs}}|{{ shu $m.X $k }}|{{ shu $m.Y $v }}|
-{{ fenint $m.Sum}}-----
{{ end }}{{ end }}`

func PrintJson(showByte []byte) {
	var out bytes.Buffer
	err := json.Indent(&out, showByte, "", "    ")
	if err != nil {
		myLogger.Log.Error("Invilad json stytle:")
		PrintAndExit(string(showByte), FailedTemplate)
	}
	PrintAndExit(string(out.Bytes()), SuccessTemplate)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func DeleteFrom(s1, s2 []string) []string {
	var res []string
	for _, i1 := range s1 {
		have := false
		for _, i2 := range s2 {
			if i1 == i2 {
				have = true
				break
			}
		}
		if have {continue}
		res = append(res, i1)
	}
	return res
}