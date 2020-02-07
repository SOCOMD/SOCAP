#include "\socap\predefined.hpp"

[] call socap_fnc_CaptureStart;

missionNamespace setVariable["socap_capture_enabled", true, true];
_captureEnabled = missionNamespace getVariable["socap_capture_enabled", false];
_frameTimer = 0.0;
_frameCount = 0;
_frameInterval = FRAME_INTERVAL;

["socap_dt"] call BIS_fnc_deltaTime;
while { _captureEnabled } do {
	_captureEnabled = missionNamespace getVariable["socap_capture_enabled", false];
	_dt = ["socap_dt"] call BIS_fnc_deltaTime;
	if(_frameTimer > 0.0) then {
		_frameTimer = _frameTimer - _dt;
	} else {
		//Update Frame Count
		_frameCount = _frameCount + 1;
		missionNamespace setVariable["socap_frame", _frameCount, true];

		//Update Entity Positions
		_entities = missionNamespace getVariable["socap_entities", []];
		{
			[_x, _frameCount] call socap_fnc_EntityUpdatePosition;
		} forEach _entities;

		_frameTimer = _frameInterval;
	};
};