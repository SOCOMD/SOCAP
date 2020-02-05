package ocap

import (
	"errors"
	"fmt"
	"strings"
)

var (
	entities    map[int]interface{}
	captureJSON capture
)

// setup any variables here
func init() {
	entities = make(map[int]interface{})
	captureJSON = capture{}

	rvTest1()
}

func RVExensionHandle(funcName string, args []string) string {
	var err error = errors.New("Not Handled")

	switch funcName {
	case ":NEW:UNIT:":
		err = rvNewUnitHandler(args)
	case ":NEW:VEH:":
		err = rvNewVehicleHandler(args)
	case ":EVENT:":
		err = rvEventHandler(args)
	case ":UPDATE:UNIT:":
		err = rvUpdateUnitHandler(args)
	case ":UPDATE:VEH:":
		err = rvUpdateVehicleHandler(args)
	case ":SAVE:":
	case ":LOG:":
	case ":START:":
		err = rvStartHandler(args)
	case ":FIRED:":
		err = rvFiredHandler(args)
	default:
		err = errors.New("Not Known")
	}

	if err != nil {
		fmt.Printf("ERR: %s, Func: %s, Args: %s\n", err, funcName, strings.Join(args, ","))
		return err.Error()
	}

	return ""
}
