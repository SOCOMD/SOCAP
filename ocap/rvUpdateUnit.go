package ocap

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type rvUpdateUnit struct {
	ID          int
	Position    vec2
	Direction   int
	IsAlive     int
	IsInVehicle int
	Name        string
	IsPlayer    int
}

// example line
// 0,[24075.3,16291],0,1,0,"ChambersAUS",1
var rvUpdateUnitRe *regexp.Regexp = regexp.MustCompile(`(\d+),\[(\d+\.?\d*?),(\d+\.?\d*?)\],(\d+),(\d+),(\d+),"(.*?)",(\d)`)

func rvUpdateUnitHandler(args []string) error {
	update, err := rvUpdateUnitParser(strings.Join(args, ","))
	if err != nil {
		return err
	}

	unit := entities[update.ID].(*entityUnit)
	unit.Positions = append(unit.Positions, eventPositionUnit{
		Position:    update.Position,
		Direction:   update.Direction,
		IsAlive:     update.IsAlive,
		IsInVehicle: update.IsInVehicle,
		Name:        update.Name,
		IsPlayer:    update.IsPlayer,
	})
	entities[update.ID] = unit

	return nil
}

func rvUpdateUnitParser(input string) (rvUpdateUnit, error) {
	match := rvUpdateUnitRe.FindStringSubmatch(input)
	if len(match) < 9 {
		return rvUpdateUnit{}, errors.New("Bad Input string")
	}

	// strip match string
	id, _ := strconv.Atoi(match[1])
	posX, _ := strconv.ParseFloat(match[2], 64)
	posY, _ := strconv.ParseFloat(match[3], 64)
	dir, _ := strconv.Atoi(match[4])
	isAlive, _ := strconv.Atoi(match[5])
	isInVehicle, _ := strconv.Atoi(match[6])
	isPlayer, _ := strconv.Atoi(match[8])

	return rvUpdateUnit{
		ID: id,
		Position: vec2{
			X: posX,
			Y: posY,
		},
		Direction:   dir,
		IsAlive:     isAlive,
		IsInVehicle: isInVehicle,
		Name:        match[7],
		IsPlayer:    isPlayer,
	}, nil
}