package common

import (
	"errors"
	"fmt"
	"strconv"
)

func CheckAndParseIdArg(value string) (int, error) {
	taskId, err := strconv.Atoi(value)
	if err != nil {
		var msg string
		if errors.Is(err, strconv.ErrSyntax) {
			msg = fmt.Sprint("invalid value \"", value, "\" for a task id\n")
		} else {
			msg = err.Error()
		}
		return 0, errors.New(msg)
	}
	return taskId, nil
}