if(!isServer) exitWith {};

addMissionEventHandler["PlayerConnected", {_this spawn socap_fnc_PlayerConnect;}];
addMissionEventHandler["HandleDisconnect", {_this spawn socap_fnc_PlayerDisconnect;}];
addMissionEventHandler["EntityKilled", {_this spawn socap_fnc_EntityKilled;}];
addMissionEventHandler["Ended", { _this call socap_fnc_CaptureEnd;}];

//["socap_processEvents", "onEachFrame", {_this call socap_fnc_ProcessEvents;}] call BIS_fnc_addStackedEventHandler;

[] spawn socap_fnc_MainLoop;