package ocap

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

var (
	frameData map[int][]json.Marshaler
)

// setup any variables here
func init() {
	frameData = make(map[int][]json.Marshaler)

	record := capture{
		MissionName:   "VTNRBC_37_Hard_day_v1",
		MissionAuthor: "[TF]Morkontar",
		WorldName:     "zargabad",
		CaptureDelay:  1.23,
		EndFrame:      1515,
		Markers:       []marker{},
		Entities: []interface{}{
			entityUnit{
				ID:         0,
				Type:       "unit",
				Name:       "[TF]Morkontar",
				Group:      "Alpha 1-1",
				Side:       "WEST",
				IsPlayer:   1,
				StartFrame: 0,
				FramesFired: []eventFire{
					eventFire{
						Frame: 673,
						Position: vec2{
							X: 4284.31,
							Y: 4383.04,
						},
					},
				},
				Positions: []eventPositionUnit{
					eventPositionUnit{
						Position: vec2{
							X: 4840.78,
							Y: 9134.32,
						},
						Direction:   357,
						IsAlive:     1,
						IsInVehicle: 0,
						Name:        "[TF]Morkontar",
						IsPlayer:    1,
					},
				},
			},
			entityVehicle{
				ID:          41,
				Class:       "car",
				Type:        "vehicle",
				Name:        "M1151 M2HB",
				StartFrame:  0,
				FramesFired: []eventFire{},
				Positions: []eventPositionVehicle{
					eventPositionVehicle{
						Position: vec3{
							X: 4841.89,
							Y: 6143.05,
							Z: 0,
						},
						Direction: 270,
						IsAlive:   1,
						Units:     []int{},
					},
				},
			},
		},
		Events: []interface{}{
			eventConnected{
				Frame: 0,
				Name:  "[StB]Forester",
			},
			eventKilled{
				Frame:    640,
				VictimID: 10,
				KillerID: 10,
				WeaponID: "M249 SAW",
				Distance: 0,
			},
		},
	}

	content, _ := json.MarshalIndent(record, "", "  ")
	fmt.Printf("%s\n", content)
}

func RVExensionHandle(funcName string, args []string) string {
	var result string
	var err error = errors.New("Not Handled")
	switch funcName {
	case ":NEW:UNIT:":
		result, err = newUnitHandler(args)
	case ":NEW:VEH:":
	case ":EVENT:":
	case ":CLEAR:":
	case ":UPDATE:UNIT:":
	case ":UPDATE:VEH:":
	case ":SAVE:":
	case ":LOG:":
	case ":START:":
	case ":FIRED:":
	default:
		err = errors.New("Not Known")
	}
	if err != nil {
		fmt.Printf("ERR: %s, Func: %s, Args: %s\n", err, funcName, strings.Join(args, ","))
		return ""
	}
	return result
}
