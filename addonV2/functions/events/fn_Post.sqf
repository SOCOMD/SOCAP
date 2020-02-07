params ["_func", "_args"];

_captureEnabled = missionNamespace getVariable["socap_capture_enabled", false];
if(!_captureEnabled) exitWith {};

_stack = missionNamespace getVariable["socap_stack", []];
_stack pushBack _this;
missionNamespace setVariable["socap_stack", _stack, true];