params ["_unit", "_weapon", "_muzzle", "_mode", "_ammo", "_magazine", "_projectile", "_gunner"];

_id = _unit getVariable["socap_entity_id", -1];
if(_id < 0) exitWith {};

_frame = missionNamespace getVariable["socap_frame", 0];
_pos = getPosATL _projectile;

_lastPos = [];
waitUntil {
	_pos = getPosATL _projectile;
	if (((_pos select 0) isEqualTo 0) || isNull _projectile) exitWith {true};
	_lastPos = _pos;
	false;
};

if(count _lastPos == 0) exitWith {};

_posX = _lastPos select 0;
_posY = _lastPos select 1;

[":FIRED:",[_id, _frame, [_posX, _posY]]] call socap_fnc_Post;