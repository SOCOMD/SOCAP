package ocap

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type rvNewUnit struct {
	Frame    int
	ID       int
	Name     string
	Group    string
	Side     string
	IsPlayer int
}

//REF: `25,1,"William Johnson","Alpha 1-2","WEST",0`
var rvNewUnitRe *regexp.Regexp = regexp.MustCompile(`(\d+),(\d+),"(.*?)","(.*?)","(.*?)",(\d)`)

func rvNewUnitHandler(args []string) error {
	newUnit, err := rvNewUnitParser(strings.Join(args, ","))
	if err != nil {
		return err
	}

	entityIDs = append(entityIDs, newUnit.ID)
	entities[newUnit.ID] = &entityUnit{
		entity: entity{
			ID:          newUnit.ID,
			Type:        "unit",
			Name:        newUnit.Name,
			StartFrame:  newUnit.Frame,
			FramesFired: []eventFire{},
		},
		Group:     newUnit.Group,
		Side:      newUnit.Side,
		IsPlayer:  newUnit.IsPlayer,
		Positions: []eventPositionUnit{},
	}

	return nil
}

func rvNewUnitParser(input string) (rvNewUnit, error) {
	match := rvNewUnitRe.FindStringSubmatch(input)
	if len(match) < 7 {
		return rvNewUnit{}, errors.New("Bad Input string")
	}

	// strip match string
	frame, _ := strconv.Atoi(match[1])
	id, _ := strconv.Atoi(match[2])
	isPlayer, _ := strconv.Atoi(match[6])
	return rvNewUnit{
		Frame:    frame,
		ID:       id,
		Name:     match[3],
		Group:    match[4],
		Side:     match[5],
		IsPlayer: isPlayer,
	}, nil
}
