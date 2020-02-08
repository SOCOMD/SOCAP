params ["_func", "_args"];

_captureEnabled = missionNamespace getVariable["socap_capture_enabled", false];
if(!_captureEnabled) exitWith {};

if(isServer) then {
	_this call socap_fnc_PostServer;
} else {
	_this remoteExec ["socap_fnc_PostServer", 2];
};

// Push events to the stack
//_stack = missionNamespace getVariable["socap_stack", []];
//_stack pushBack _this;
//missionNamespace setVariable["socap_stack", _stack, true];