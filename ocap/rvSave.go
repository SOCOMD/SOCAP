package ocap

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
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

	fn := fmt.Sprintf("%d_socap.json", time.Now().Unix())
	err = curlSave(fn)
	if err != nil {
		return err
	}

	err = curlUpload(fn)
	if err != nil {
		return err
	}

	err = curlUpdate(fn)
	if err != nil {
		return err
	}

	teardownLogger()

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

func curlSave(fn string) error {
	tmpfn := filepath.Join(tempDir, fn)
	output, err := json.Marshal(captureJSON)
	if err != nil {
		return err
	}

	f, err := os.Create(tmpfn)
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(string(output))

	return nil
}

func curlUpload(fn string) error {
	secret := "uid10t"

	tmpfn := filepath.Join(tempDir, fn)
	f, err := os.Open(tmpfn)
	if err != nil {
		return err
	}
	defer f.Close()

	url := fmt.Sprintf(`http://127.0.0.1:8080/recieve.php?option=addFile&fileName=%s&secret=%s`, fn, secret)
	req, err := http.NewRequest("POST", url, f)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func curlUpdate(fn string) error {
	secret := "uid10t"
	worldName := captureJSON.WorldName
	missionName := captureJSON.MissionName
	missionDuration := (int64)(captureJSON.CaptureDelay * (float64)(captureJSON.EndFrame))

	url := fmt.Sprintf(`http://127.0.0.1:8080/recieve.php?option=dbInsert&secret=%s&worldName=%s&missionName=%s&missionDuration=%d&type=coop&filename=%s`, secret, worldName, missionName, missionDuration, fn)
	resp, err := http.Post(url, "", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
