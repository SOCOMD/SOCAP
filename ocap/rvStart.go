package ocap

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type rvStart struct {
	WorldName     string
	MissionName   string
	MissionAuthor string
	CaptureDelay  float64
}

// example line
// "Altis","tempMissionSP","",1
var rvStartRe *regexp.Regexp = regexp.MustCompile(`"(.*?)","(.*?)","(.*?)",(\d+)`)

func rvStartHandler(args []string) error {
	start, err := rvStartParser(strings.Join(args, ","))
	if err != nil {
		return err
	}

	resetCapture()

	captureJSON.WorldName = start.WorldName
	captureJSON.MissionName = start.MissionName
	captureJSON.MissionAuthor = start.MissionAuthor
	captureJSON.CaptureDelay = start.CaptureDelay

	return nil
}

func rvStartParser(input string) (rvStart, error) {
	match := rvStartRe.FindStringSubmatch(input)
	if len(match) < 5 {
		return rvStart{}, errors.New("Bad Input string")
	}

	// strip match string
	captureInterval, _ := strconv.ParseFloat(match[4], 64)
	return rvStart{
		WorldName:     match[1],
		MissionName:   match[2],
		MissionAuthor: match[3],
		CaptureDelay:  captureInterval,
	}, nil
}
