package ocap

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type rvEvent struct {
	Fame int
	Type string
}

var rvEventConnectedRe *regexp.Regexp = regexp.MustCompile(`"(.*)"`)

type rvEventConnected struct {
	rvEvent
	Name string
}

var rvEventDisconnectedRe *regexp.Regexp = regexp.MustCompile(`"(.*)"`)

type rvEventDisconnected struct {
	rvEvent
	Name string
}

var rvEventHitRe *regexp.Regexp = regexp.MustCompile(`(\d+),\[(\d+),"(.*?)"\],(\d+)`)

type rvEventHit struct {
	rvEvent
	VictimID int
	HitterID int
	WeaponID string
	Distance int
}

var rvEventKilledRe *regexp.Regexp = regexp.MustCompile(`(\d+),\[(\d+),"(.*?)"\],(\d+)`)

type rvEventKilled struct {
	rvEvent
	VictimID int
	KillerID int
	WeaponID string
	Distance int
}

var rvEventEndMissionRe *regexp.Regexp = regexp.MustCompile(`\["(.*?)","(.*?)"\]`)

type rvEventEndMission struct {
	rvEvent
	SideWon     string
	Description string
}

var rvEventRe *regexp.Regexp = regexp.MustCompile(`(\d+),"(.*?)",(.*)`)

func rvEventHandler(args []string) error {
	event, input, err := rvEventParser(strings.Join(args, ","))
	if err != nil {
		return err
	}

	switch event.Type {
	case "connected":
		err = rvEventConnectedHandler(event, input)
	case "disconnected":
		err = rvEventDisconnectedHandler(event, input)
	case "hit":
		err = rvEventHitHandler(event, input)
	case "killed":
		err = rvEventKilledHandler(event, input)
	case "endMission":
		err = rvEventEndMissionHandler(event, input)
	}

	return nil
}

func rvEventParser(input string) (rvEvent, string, error) {
	match := rvEventRe.FindStringSubmatch(input)
	if len(match) < 4 {
		return rvEvent{}, "", errors.New("Bad Input string")
	}

	frame, _ := strconv.Atoi(match[1])
	return rvEvent{
		Fame: frame,
		Type: match[2],
	}, match[3], nil
}

////////////////////////////////////////////////////////////////////////////////

func rvEventConnectedHandler(event rvEvent, input string) error {
	connected, err := rvEventConnectedParser(event, input)
	if err != nil {
		return err
	}

	captureJSON.Events = append(captureJSON.Events, eventConnected{
		Frame: connected.Fame,
		Name:  connected.Name,
	})
	return nil
}

func rvEventConnectedParser(event rvEvent, input string) (rvEventConnected, error) {
	match := rvEventConnectedRe.FindStringSubmatch(input)
	if len(match) < 2 {
		return rvEventConnected{}, errors.New("Bad Input string")
	}

	// strip match string
	return rvEventConnected{
		rvEvent: event,
		Name:    match[1],
	}, nil
}

////////////////////////////////////////////////////////////////////////////////

func rvEventDisconnectedHandler(event rvEvent, input string) error {
	disconnected, err := rvEventDisconnectedParser(event, input)
	if err != nil {
		return err
	}

	captureJSON.Events = append(captureJSON.Events, eventDisconnected{
		Frame: disconnected.Fame,
		Name:  disconnected.Name,
	})
	return nil
}

func rvEventDisconnectedParser(event rvEvent, input string) (rvEventDisconnected, error) {
	match := rvEventConnectedRe.FindStringSubmatch(input)
	if len(match) < 2 {
		return rvEventDisconnected{}, errors.New("Bad Input string")
	}

	// strip match string
	return rvEventDisconnected{
		rvEvent: event,
		Name:    match[1],
	}, nil
}

////////////////////////////////////////////////////////////////////////////////

func rvEventHitHandler(event rvEvent, input string) error {
	hit, err := rvEventHitParser(event, input)
	if err != nil {
		return err
	}

	captureJSON.Events = append(captureJSON.Events, eventHit{
		Frame:    hit.Fame,
		VictimID: hit.VictimID,
		HitterID: hit.HitterID,
		WeaponID: hit.WeaponID,
		Distance: hit.Distance,
	})

	return nil
}

func rvEventHitParser(event rvEvent, input string) (rvEventHit, error) {
	match := rvEventHitRe.FindStringSubmatch(input)
	if len(match) < 5 {
		return rvEventHit{}, errors.New("Bad Input string")
	}

	// strip match string
	victimID, _ := strconv.Atoi(match[1])
	hitterID, _ := strconv.Atoi(match[2])
	distance, _ := strconv.Atoi(match[4])

	return rvEventHit{
		rvEvent:  event,
		VictimID: victimID,
		HitterID: hitterID,
		WeaponID: match[3],
		Distance: distance,
	}, nil
}

////////////////////////////////////////////////////////////////////////////////

func rvEventKilledHandler(event rvEvent, input string) error {
	killed, err := rvEventKilledParser(event, input)
	if err != nil {
		return err
	}

	captureJSON.Events = append(captureJSON.Events, eventKilled{
		Frame:    killed.Fame,
		VictimID: killed.VictimID,
		KillerID: killed.KillerID,
		WeaponID: killed.WeaponID,
		Distance: killed.Distance,
	})

	return nil
}

func rvEventKilledParser(event rvEvent, input string) (rvEventKilled, error) {
	match := rvEventKilledRe.FindStringSubmatch(input)
	if len(match) < 2 {
		return rvEventKilled{}, errors.New("Bad Input string")
	}

	victimID, _ := strconv.Atoi(match[1])
	hitterID, _ := strconv.Atoi(match[2])
	distance, _ := strconv.Atoi(match[4])

	// strip match string
	return rvEventKilled{
		rvEvent:  event,
		VictimID: victimID,
		KillerID: hitterID,
		WeaponID: match[3],
		Distance: distance,
	}, nil
}

////////////////////////////////////////////////////////////////////////////////

func rvEventEndMissionHandler(event rvEvent, input string) error {
	endMission, err := rvEventEndMissionParser(event, input)
	if err != nil {
		return err
	}

	captureJSON.Events = append(captureJSON.Events, eventEndMission{
		Frame:              endMission.Fame,
		SideWon:            endMission.SideWon,
		MissionDescription: endMission.Description,
	})

	captureJSON.EndFrame = endMission.Fame

	return fmt.Errorf("Unsupported Event")
}

func rvEventEndMissionParser(event rvEvent, input string) (rvEventEndMission, error) {
	match := rvEventEndMissionRe.FindStringSubmatch(input)
	if len(match) < 2 {
		return rvEventEndMission{}, errors.New("Bad Input string")
	}

	// strip match string
	return rvEventEndMission{
		rvEvent:     event,
		SideWon:     match[1],
		Description: match[2],
	}, nil
}
