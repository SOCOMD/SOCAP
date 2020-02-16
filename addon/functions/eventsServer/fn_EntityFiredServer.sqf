params ["_unit", "_weapon", "_muzzle", "_mode", "_ammo", "_magazine", "_projectile", "_gunner"];

if(!isServer) exitWith {};

_id = _unit getVariable["socap_entity_id", -1];
if(_id < 0) exitWith {};

_frame = socap_global_frame;

_pos = getPosATL _unit;
waitUntil {
	if(isNull _projectile) exitWith {true};
	_pos = getPosATL _projectile;
	false;
};

_posX = _pos select 0;
_posY = _pos select 1;

[":FIRED:",[_id, _frame, [_posX, _posY]]] call socap_fnc_Post;