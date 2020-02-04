package ocap

import "encoding/json"

var (
	frameData map[int]json.Marshaler
)

// setup any variables here
func init() {
	frameData = make(map[int]json.Marshaler)
}

func RVExensionHandle(funcName string, args []string) string {
	var result string
	var err error
	switch funcName {
	case ":NEW:UNIT:":
	case ":NEW:VEH:":
	case ":EVENT:":
	case ":CLEAR:":
	case ":UPDATE:UNIT:":
	case ":UPDATE:VEH:":
	case ":SAVE:":
	case ":LOG:":
	case ":START:":
	case ":FIRED:":
	}
	if err != nil {
		return ""
	}
	return result
}
