package ocap

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
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
	buf := bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(captureJSON)
	if err != nil {
		return err
	}

	err = curlUpload(fn, &buf)
	if err != nil {
		return err
	}

	err = curlUpdate(fn)
	if err != nil {
		return err
	}

	resetCapture()
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

func curlUpload(fn string, r io.Reader) error {
	secret := "uid10t"

	url, _ := url.Parse(fmt.Sprintf(`http://127.0.0.1:9000/recieve.php?option=addFile&fileName=%s&secret=%s`, fn, secret))
	req, err := http.NewRequest("POST", url.String(), r)
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

	url, _ := url.Parse(fmt.Sprintf(`http://127.0.0.1:9000/recieve.php?option=dbInsert&secret=%s&worldName=%s&missionName=%s&missionDuration=%d&type=coop&filename=%s`, secret, worldName, missionName, missionDuration, fn))
	resp, err := http.Post(url.String(), "", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
