params ["_unit", "_source", "_damage", "_instigator"];

if(!isServer) exitWith {};

_unitID = _unit getVariable["socap_entity_id", -1];
if(_unitID < 0) exitWith {};

_sourceID = _source getVariable["socap_entity_id", -1];
if(_sourceID < 0) exitWith {};

_frame = socap_global_frame;
_sourceWeapon = getText (configFile >> "CfgWeapons" >> currentWeapon _source >> "displayName");
_distance = round (_unit distance _source);

[":EVENT:", [_frame, "hit", _unitID, [_source, _sourceWeapon], _distance]] call socap_fnc_Post;