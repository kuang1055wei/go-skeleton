package model

import (
	"fmt"
	"strings"

	"github.com/golang-module/carbon"
)

//carbon的database类型都没有UnmarshalJSON，这里自定义实现
//2012-08-05 13:14:15
type DateTime struct {
	carbon.Carbon
}

func (c DateTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, c.ToDateTimeString())), nil
}

func (t *DateTime) UnmarshalJSON(b []byte) error {
	//去除两边的 '"'
	*t = DateTime{carbon.Parse(strings.ReplaceAll(string(b), "\"", ""))}
	return nil
}
