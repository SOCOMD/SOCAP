package ocap

import "encoding/json"

type capture struct {
	MissionName   string        `json:"missionName"`
	MissionAuthor string        `json:"missionAuthor"`
	WorldName     string        `json:"worldName"`
	CaptureDelay  float32       `json:"captureDelay"`
	EndFrame      int           `json:"endFrame"`
	Markers       []marker      `json:"Markers"`
	Entities      []interface{} `json:"entities"`
	Events        []interface{} `json:"events"`
}

type marker struct{}

type entityUnit struct {
	ID          int                 `json:"id"`
	Type        string              `json:"type"`
	Name        string              `json:"name"`
	Group       string              `json:"group"`
	Side        string              `json:"side"`
	IsPlayer    int                 `json:"isPlayer"`
	StartFrame  int                 `json:"startFrameNum"`
	FramesFired []eventFire         `json:"framesFired"`
	Positions   []eventPositionUnit `json:"positions"`
}

type entityVehicle struct {
	ID          int                    `json:"id"`
	Class       string                 `json:"class"`
	Type        string                 `json:"type"`
	Name        string                 `json:"name"`
	StartFrame  int                    `json:"startFrameNum"`
	FramesFired []eventFire            `json:"framesFired"`
	Positions   []eventPositionVehicle `json:"positions"`
}

type eventFire struct {
	Frame    int
	Position vec2
}

func (e eventFire) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{
		e.Frame,
		e.Position,
	})
}

type eventPositionUnit struct {
	Position    vec2
	Direction   int
	IsAlive     int
	IsInVehicle int
	Name        string
	IsPlayer    int
}

func (e eventPositionUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{
		e.Position,
		e.Direction,
		e.IsAlive,
		e.IsInVehicle,
		e.Name,
		e.IsPlayer,
	})
}

type eventPositionVehicle struct {
	Position  vec3
	Direction int
	IsAlive   int
	Units     []int
}

func (e eventPositionVehicle) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{
		e.Position,
		e.Direction,
		e.IsAlive,
		e.Units,
	})
}

type eventConnected struct {
	Frame int
	Name  string
}

func (e eventConnected) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{
		e.Frame,
		"connected",
		e.Name,
	})
}

type eventKilled struct {
	Frame    int
	VictimID int
	KillerID int
	WeaponID string
	Distance int
}

func (e eventKilled) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{
		e.Frame,
		"killed",
		e.VictimID,
		[]interface{}{
			e.KillerID,
			e.WeaponID,
		},
		e.Distance,
	})
}

type eventEndMission struct {
	Frame              int
	SideWon            string
	MissionDescription string
}

func (e eventEndMission) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{
		e.Frame,
		"endMission",
		[]interface{}{
			e.SideWon,
			e.MissionDescription,
		},
	})
}

type vec2 struct {
	X float64
	Y float64
}

func (e vec2) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{
		e.X,
		e.Y,
	})
}

type vec3 struct {
	X float64
	Y float64
	Z float64
}

func (e vec3) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{
		e.X,
		e.Y,
		e.Z,
	})
}
