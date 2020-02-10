if(!isServer) exitWith {};

//Set Globals
socap_global_captureEnabled = false;
socap_global_entity_id = 0;
socap_global_frame = 0;

addMissionEventHandler["PlayerConnected", {_this spawn socap_fnc_PlayerConnectServer;}];
addMissionEventHandler["HandleDisconnect", {_this spawn socap_fnc_PlayerDisconnectServer;}];
addMissionEventHandler["EntityKilled", {
	params ["_unit", "_killer", "_instigator", "_useEffects"];
	_clientID = owner _unit;
	[_unit, _killer, _instigator, _useEffects, socap_global_frame] remoteExec ["socap_fnc_EntityKilled", _clientID];
}];

//["socap_processEvents", "onEachFrame", {_this call socap_fnc_ProcessEvents;}] call BIS_fnc_addStackedEventHandler;