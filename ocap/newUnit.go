package ocap

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type unit struct {
	Event    string `json:"event"`
	Frame    int    `json:"startFrameNum"`
	Type     string `json:"type"` // injected and not included "unit" for player, "" for vehicle
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Group    string `json:"group"`
	Side     string `json:"side"`
	IsPlayer int    `json:"isPlayer"`
}

func (u unit) MarshalJSON() ([]byte, error) {
	return json.Marshal(u)
}

// example line
// 0,0,"ChambersAUS","Alpha 1-1","WEST",1
var newUnitRe *regexp.Regexp = regexp.MustCompile(`(\d+),(\d+),"(.*?)","(.*?)","(.*?)",(\d)`)

func newUnitHandler(args []string) (string, error) {
	u, err := newUnitParser(strings.Join(args, ","))
	if err != nil {
		return "", err
	}
	frameData[u.Frame] = append(frameData[u.Frame], u)
	return "", errors.New("Not implemented")
}

func newUnitParser(input string) (unit, error) {
	match := newUnitRe.FindStringSubmatch(input)
	if len(match) < 7 {
		return unit{}, errors.New("Bad Input string")
	}
	// strip match string
	frame, _ := strconv.Atoi(match[1])
	id, _ := strconv.Atoi(match[2])
	isPlayer, _ := strconv.Atoi(match[6])
	return unit{
		Event:    "created",
		Frame:    frame,
		Type:     "unit",
		ID:       id,
		Name:     match[3],
		Group:    match[4],
		Side:     match[5],
		IsPlayer: isPlayer,
	}, nil
}
