#include "\socap\predefined.hpp"

[] call socap_fnc_CaptureStartServer;

_frameTimer = 0.0;
_frameInterval = FRAME_INTERVAL;

["socap_dt"] call BIS_fnc_deltaTime;
while { socap_global_captureEnabled } do {
	_dt = ["socap_dt"] call BIS_fnc_deltaTime;
	if(_frameTimer > 0.0) then {
		_frameTimer = _frameTimer - _dt;
	} else {
		//Update Frame Count
		socap_global_frame = socap_global_frame + 1;

		//Update Entity Positions
		_entities = missionNamespace getVariable["socap_entities", []];
		{
			if(!(isNull _x)) then {
				[_x, socap_global_frame] call socap_fnc_EntityUpdatePositionServer;
			};
		} forEach _entities;

		_frameTimer = _frameInterval;
	};
};