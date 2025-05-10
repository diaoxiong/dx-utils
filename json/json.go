package json

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"dx-utils/regex"
)

type JsonBool bool

func (b *JsonBool) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), `"`)
	*b = str == "true"
	return nil
}

type JsonInt int64

func (i *JsonInt) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), `"`)
	if regex.IsInt64(str) {
		result, err := strconv.Atoi(str)
		if err != nil {
			return errors.Wrap(err, "strconv.Atoi failed")
		}
		*i = JsonInt(result)
	}

	return nil
}
