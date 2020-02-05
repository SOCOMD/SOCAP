package ocap

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type rvSave struct {
	WorldName     string
	MissionName   string
	MissionAuthor string
	CaptureDelay  float64
	Frame         int
}

var rvSaveRe *regexp.Regexp = regexp.MustCompile(`"(.*?)","(.*?)","(.*?)",(\d+\.?\d*?),(\d)`)

func rvSaveHandler(args []string) error {
	save, err := rvSaveParser(strings.Join(args, ","))
	if err != nil {
		return err
	}

	captureJSON.WorldName = save.WorldName
	captureJSON.MissionName = save.MissionName
	captureJSON.MissionAuthor = save.MissionAuthor
	captureJSON.CaptureDelay = save.CaptureDelay
	captureJSON.EndFrame = save.Frame

	for _, v := range entities {
		captureJSON.Entities = append(captureJSON.Entities, v)
	}

	return nil
}

func rvSaveParser(input string) (rvSave, error) {
	match := rvSaveRe.FindStringSubmatch(input)
	if len(match) < 6 {
		return rvSave{}, errors.New("Bad Input string")
	}

	captureInterval, _ := strconv.ParseFloat(match[4], 64)
	frame, _ := strconv.Atoi(match[5])
	return rvSave{
		WorldName:     match[1],
		MissionName:   match[2],
		MissionAuthor: match[3],
		CaptureDelay:  captureInterval,
		Frame:         frame,
	}, nil
}
