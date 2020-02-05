package ocap

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type rvFired struct {
	ID       int
	Frame    int
	Position vec2
}

//REF: `2,50,[24869.5,17037.4]`
var rvFiredRe *regexp.Regexp = regexp.MustCompile(`(\d+),(\d+),\[(\d+\.?\d*?),(\d+\.?\d*?)\]`)

func rvFiredHandler(args []string) error {
	event, err := rvFiredParser(strings.Join(args, ","))
	if err != nil {
		return err
	}

	entity := entities[event.ID].(entityFired)
	if entity == nil {
		return fmt.Errorf("Failed to cast")
	}

	entity.addFiredEvent(eventFire{
		Frame:    event.Frame,
		Position: event.Position,
	})

	return nil
}

func rvFiredParser(input string) (rvFired, error) {
	match := rvFiredRe.FindStringSubmatch(input)
	if len(match) < 5 {
		return rvFired{}, errors.New("Bad Input string")
	}

	// strip match string
	id, _ := strconv.Atoi(match[1])
	frame, _ := strconv.Atoi(match[2])
	posX, _ := strconv.ParseFloat(match[3], 64)
	posY, _ := strconv.ParseFloat(match[4], 64)
	return rvFired{
		ID:    id,
		Frame: frame,
		Position: vec2{
			X: posX,
			Y: posY,
		},
	}, nil
}
