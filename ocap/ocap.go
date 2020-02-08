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
	dataBus chan extData
)

type extData struct {
	FuncName string
	Args     []string
}

// setup any variables here
func init() {
	tempDir, _ = ioutil.TempDir("", "socap")
	setupLogger()
	resetCapture()
	// shouldn't need a buffer, but with a time step of 1 second
	// 1000 items is ~16m of data. for an average operation of
	// 1.5 hours if the compute can not keep up this buffer should
	// be more than enough to handle high load periods.
	// (cbf to rewrite the entire addon to allow full concurency and out of order processing)
	dataBus = make(chan extData, 1000)
}

func resetCapture() {
	entityIDs = []int{}
	entities = make(map[int]interface{})
	captureJSON = capture{}
	go handleLoop()
}

func RVExensionHandle(funcName string, args []string) string {
	dataBus <- extData{
		FuncName: funcName,
		Args:     args,
	}
	return ""
}

func handleLoop() {
	var err error = errors.New("Not Handled")

	for {
		if logger == nil {
			setupLogger()
		}
		select {
		case data, ok := <-dataBus:
			if !ok {
				// channel is busto lets exit
				logger.Print("ERR: Channel busted, exiting\n")
				return
			}
			funcName, args := data.FuncName, data.Args
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
			}
			logger.Printf("Func: %s, Args: %s\n", funcName, strings.Join(args, ","))
		}
		logFile.Sync()
	}
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
