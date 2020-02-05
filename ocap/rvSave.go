package ocap

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type rvSave struct {
	WorldName     string
	MissionName   string
	MissionAuthor string
	CaptureDelay  float64
	Frame         int
}

//REF: `"Altis","tempMissionSP","",1.1,999`
var rvSaveRe *regexp.Regexp = regexp.MustCompile(`"(.*?)","(.*?)","(.*?)",(\d+\.?\d*?),(\d+)`)

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

	sort.Ints(entityIDs)
	for _, v := range entityIDs {
		captureJSON.Entities = append(captureJSON.Entities, entities[v])
	}

	output, err := json.Marshal(captureJSON)
	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("C:\\logs\\%d_socap.json", time.Now().Unix())
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}

	f.WriteString(string(output))
	f.Close()

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
