package ocap

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

var (
	frameData map[int][]json.Marshaler
)

type position struct {
	X float64
	Y float64
}

// setup any variables here
func init() {
	frameData = make(map[int][]json.Marshaler)
}

func RVExensionHandle(funcName string, args []string) string {
	var result string
	var err error = errors.New("Not Handled")
	switch funcName {
	case ":NEW:UNIT:":
		result, err = newUnitHandler(args)
	case ":NEW:VEH:":
	case ":EVENT:":
	case ":CLEAR:":
	case ":UPDATE:UNIT:":
	case ":UPDATE:VEH:":
	case ":SAVE:":
	case ":LOG:":
	case ":START:":
	case ":FIRED:":
	default:
		err = errors.New("Not Known")
	}
	if err != nil {
		fmt.Printf("ERR: %s, Func: %s, Args: %s\n", err, funcName, strings.Join(args, ","))
		return ""
	}
	return result
}
