if(!isServer) exitWith {};

//Set Globals
socap_global_captureEnabled = false;

addMissionEventHandler["PlayerConnected", {_this spawn socap_fnc_PlayerConnectServer;}];
addMissionEventHandler["HandleDisconnect", {_this spawn socap_fnc_PlayerDisconnectServer;}];
addMissionEventHandler["EntityKilled", {_this spawn socap_fnc_EntityKilledServer;}];
addMissionEventHandler["Ended", { _this call socap_fnc_CaptureEndServer;}];

//["socap_processEvents", "onEachFrame", {_this call socap_fnc_ProcessEvents;}] call BIS_fnc_addStackedEventHandler;

[] spawn socap_fnc_MainLoop;