_stack = missionNamespace getVariable["socap_stack", []];
_count = ((count _stack) * 0.15) max 5;

for "_i" from 1 to _count do {
	if(count _stack != 0) then {
		_stack select 0 call socap_fnc_PostServer;
		_stack deleteAt 0;
	};
};

missionNamespace setVariable["socap_stack", _stack, true];