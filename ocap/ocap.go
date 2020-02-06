package ocap

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	entityIDs   []int
	entities    map[int]interface{}
	captureJSON capture

	tempDir string
	logFile *os.File
	logger  *log.Logger
)

// setup any variables here
func init() {
	tempDir, _ = ioutil.TempDir("", "socap")
	setupLogger()
	resetCapture()
}

func resetCapture() {
	entityIDs = []int{}
	entities = make(map[int]interface{})
	captureJSON = capture{}
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
		err = rvSaveHandler(args)
	case ":LOG:":
	case ":START:":
		err = rvStartHandler(args)
	case ":FIRED:":
		err = rvFiredHandler(args)
	default:
		err = errors.New("Not Known")
	}

	if err != nil {
		logger.Printf("ERR: %s, Func: %s, Args: %s\n", err, funcName, strings.Join(args, ","))
		return err.Error()
	}

	logger.Printf("Func: %s, Args: %s\n", funcName, strings.Join(args, ","))
	return ""
}

func setupLogger() {
	tmpfn := filepath.Join(tempDir, "socap.log")
	logFile, err := os.OpenFile(tmpfn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	logger = log.New(logFile, "", log.LstdFlags)
}

func teardownLogger() {
	if logFile != nil {
		logFile.Close()
	}
}
