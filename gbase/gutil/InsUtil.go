package gutil

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func DumpInterface(v interface{}) string {
	jsonVal, _ := json.Marshal(v)
	return fmt.Sprintf("type:%s val:%s", reflect.TypeOf(v), jsonVal)
}
