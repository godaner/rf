package rf

import (
	"encoding/json"
	"github.com/duke-git/lancet/v2/formatter"
)

func ToJson(src any) string {
	bs, _ := json.Marshal(src)
	return string(bs)
}

func ToPretty(src any) string {
	bs, _ := formatter.Pretty(src)
	return bs
}
