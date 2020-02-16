#include "\socap\predefined.hpp"

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
		{
			if(!(isNull _x)) then {
				_clientID = owner _x;
				[_x, socap_global_frame] remoteExec ["socap_fnc_EntityUpdatePosition", _clientID];
			};
		} forEach socap_global_entities;

		_frameTimer = _frameInterval;
	};
};