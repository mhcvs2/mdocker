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
	"os/exec"
	"strconv"
)

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

func PrintErrorAndExit(message, errorTemplate string) {
	Tmpl(fmt.Sprintf(errorTemplate, message), nil)
	os.Exit(2)
}

func Exec_shell(s string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}

func OutIsNil(s string) bool {
	cmd := s + "|wc -l"
	out, err := Exec_shell(cmd)
	if err != nil {
		return false
	}
	if outInt, _ := strconv.Atoi(out); outInt == 0 {
		return false
	}
	return true
}