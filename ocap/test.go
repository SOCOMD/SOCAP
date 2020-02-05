package ocap

import (
	"encoding/json"
	"fmt"
)

func rvTest1() {
	RVExensionHandle(`:START:`, []string{`"Altis","tempMissionSP","",1`})
	RVExensionHandle(`:NEW:UNIT:`, []string{`0,0,"ChambersAUS","Alpha 1-1","WEST",1`})
	RVExensionHandle(`:NEW:UNIT:`, []string{`25,1,"William Johnson","Alpha 1-2","WEST",0`})
	RVExensionHandle(`:NEW:UNIT:`, []string{`45,2,"Ruslan Pashinin","Alpha 1-1","EAST",0`})
	RVExensionHandle(`:NEW:VEH:`, []string{`0,1,"heli","CH-47I Chinook"`})
	RVExensionHandle(`:UPDATE:UNIT:`, []string{`0,[24075.3,16291],0,1,0,"ChambersAUS",1`})
	RVExensionHandle(`:UPDATE:VEH:`, []string{`1,[24077.5,16301.7,0],0,1,[]`})
	RVExensionHandle(`:UPDATE:UNIT:`, []string{`0,[24075.3,16291],4,1,0,"ChambersAUS",1`})
	RVExensionHandle(`:UPDATE:VEH:`, []string{`1,[24077.6,16301.7,0],0,1,[]`})
	RVExensionHandle(`:UPDATE:UNIT:`, []string{`0,[24075.7,16292.1],19,1,0,"ChambersAUS",1`})
	RVExensionHandle(`:UPDATE:VEH:`, []string{`1,[24077.6,16301.7,0],0,1,[]`})
	RVExensionHandle(`:UPDATE:UNIT:`, []string{`0,[24076.3,16294],11,1,0,"ChambersAUS",1`})
	RVExensionHandle(`:UPDATE:VEH:`, []string{`1,[24077.6,16301.7,0],0,1,[]`})
	RVExensionHandle(`:UPDATE:UNIT:`, []string{`0,[24076.3,16294],11,1,0,"ChambersAUS",1`})
	RVExensionHandle(`:UPDATE:VEH:`, []string{`1,[24077.6,16301.7,0],0,1,[]`})
	RVExensionHandle(`:FIRED:`, []string{`2,50,[24846.3,17054.6]`})
	RVExensionHandle(`:FIRED:`, []string{`2,50,[24818.7,17033.2]`})
	RVExensionHandle(`:FIRED:`, []string{`2,50,[24844.8,17036]`})
	RVExensionHandle(`:FIRED:`, []string{`2,50,[24869.5,17037.4]`})
	RVExensionHandle(`:EVENT:`, []string{`0,"connected","ChambersAUS"`})
	RVExensionHandle(`:EVENT:`, []string{`0,"connected","William Johnson"`})
	RVExensionHandle(`:EVENT:`, []string{`45,"hit",0,[1,"M4"],10`})
	RVExensionHandle(`:EVENT:`, []string{`50,"killed",0,[1,"M4"],15`})
	RVExensionHandle(`:EVENT:`, []string{`999,"endMission",["",""]`})
	RVExensionHandle(`:EVENT:`, []string{`999,"disconnected","ChambersAUS"`})

	for _, v := range entities {
		captureJSON.Entities = append(captureJSON.Entities, v)
	}

	output, _ := json.Marshal(captureJSON)
	fmt.Printf("%s\n", output)
}
