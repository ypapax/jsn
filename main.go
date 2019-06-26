package jsn

import (
	"encoding/json"

	"github.com/golang/glog"
)

func B(inp interface{}) string {
	b, err := json.MarshalIndent(&inp, "", "   ")
	if err != nil {
		glog.Error(err)
		return err.Error()
	}
	return string(b)
}
