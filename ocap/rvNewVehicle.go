package ocap

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type rvNewVehicle struct {
	Frame int
	ID    int
	Class string
	Name  string
}

// example line
// 0,1,"heli","CH-47I Chinook"
var rvNewVehicleRe *regexp.Regexp = regexp.MustCompile(`(\d+),(\d+),"(.*?)","(.*?)"`)

func rvNewVehicleHandler(args []string) error {
	newVehicle, err := rvNewVehicleParser(strings.Join(args, ","))
	if err != nil {
		return err
	}

	entities[newVehicle.ID] = &entityVehicle{
		entity: entity{
			ID:          newVehicle.ID,
			Type:        "vehicle",
			Name:        newVehicle.Name,
			FramesFired: []eventFire{},
		},
		Class:     newVehicle.Class,
		Positions: []eventPositionVehicle{},
	}

	return nil
}

func rvNewVehicleParser(input string) (rvNewVehicle, error) {
	match := rvNewVehicleRe.FindStringSubmatch(input)
	if len(match) < 5 {
		return rvNewVehicle{}, errors.New("Bad Input string")
	}

	// strip match string
	frame, _ := strconv.Atoi(match[1])
	id, _ := strconv.Atoi(match[2])
	return rvNewVehicle{
		Frame: frame,
		ID:    id,
		Class: match[3],
		Name:  match[4],
	}, nil
}
