package ocap

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type rvUpdateVehicle struct {
	ID        int
	Position  vec3
	Direction int
	IsAlive   int
	Crew      []int
}

//REF: `1,[24077.5,16301.7,0],0,1,[]`
var rvUpdateVehicleRe *regexp.Regexp = regexp.MustCompile(`(\d+),\[(\d+\.?\d*?),(\d+\.?\d*?),(\d+\.?\d*?)\],(\d+),(\d+),\[(.*?)\]`)

func rvUpdateVehicleHandler(args []string) error {
	update, err := rvUpdateVehicleParser(strings.Join(args, ","))
	if err != nil {
		return err
	}

	if _, ok := entities[update.ID]; !ok {
		return errors.New(string(update.ID) + " Does not exist, can't update")
	}
	vehicle, ok := entities[update.ID].(*entityVehicle)
	if !ok {
		return errors.New(string(update.ID) + " Can not type cast to entityVehicle")
	}
	vehicle.Positions = append(vehicle.Positions, eventPositionVehicle{
		Position:  update.Position,
		Direction: update.Direction,
		IsAlive:   update.IsAlive,
		Crew:      update.Crew,
	})
	entities[update.ID] = vehicle
	return nil
}

func rvUpdateVehicleParser(input string) (rvUpdateVehicle, error) {
	match := rvUpdateVehicleRe.FindStringSubmatch(input)
	if len(match) < 7 {
		return rvUpdateVehicle{}, errors.New("Bad Input string")
	}

	id, _ := strconv.Atoi(match[1])
	posX, _ := strconv.ParseFloat(match[2], 64)
	posY, _ := strconv.ParseFloat(match[3], 64)
	posZ, _ := strconv.ParseFloat(match[4], 64)
	dir, _ := strconv.Atoi(match[5])
	isAlive, _ := strconv.Atoi(match[6])

	crewStr := strings.Split(match[7], ",")
	crew := []int{}
	if len(crewStr) > 0 {
		for _, v := range crewStr {
			id, _ := strconv.Atoi(v)
			crew = append(crew, id)
		}
	}

	return rvUpdateVehicle{
		ID: id,
		Position: vec3{
			X: posX,
			Y: posY,
			Z: posZ,
		},
		Direction: dir,
		IsAlive:   isAlive,
		Crew:      crew,
	}, nil
}
