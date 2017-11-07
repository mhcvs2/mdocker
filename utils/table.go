package utils

import (
	"github.com/olekukonko/tablewriter"
	"os"
	"reflect"
	"strconv"
)

func getString(value reflect.Value) string {
	switch value.Kind() {
	case reflect.String:
		return value.String()
	case reflect.Int:
		return strconv.FormatInt(value.Int(),10)
	case reflect.Int32:
		return strconv.FormatInt(value.Int(),10)
	case reflect.Int64:
		return strconv.FormatInt(value.Int(),10)
	case reflect.Bool:
		return strconv.FormatBool(value.Bool())
	}
	return "not parse"
}

func ShowStructList(obs []interface{}, keys []string) string {
	var value reflect.Value
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(keys)
	for _, ob := range obs {
		showData := make([]string, len(keys))
		value = reflect.ValueOf(ob)
		if value.Kind() != reflect.Struct{
			return "Input must be Struct"
		}
		for i, key := range keys {
			field := value.FieldByName(key)
			showData[i] = getString(field);
		}
		table.Append(showData)
	}
	table.Render()
	return ""
}


