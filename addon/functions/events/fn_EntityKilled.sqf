params ["_unit", "_killer", "_instigator", "_useEffects", ["_frame", -1]];

if(!(local _unit)) exitWith {};

if(_frame < 0) exitWith {};

if (isNull _instigator) then {_instigator = UAVControl vehicle _killer select 0}; // UAV/UGV player operated road kill
if (isNull _instigator) then {_instigator = _killer};
if ((isNull _instigator) || (_instigator isEqualTo _unit)) then {_instigator = _unit getVariable ["ace_medical_lastDamageSource", _unit]};

_unitID = _unit getVariable["socap_entity_id", -1];
if(_unitID < 0) exitWith {};

_instigatorID = _instigator getVariable["socap_entity_id", -1];
if(_instigatorID < 0) exitWith {};

_dist = round(_instigator distance _unit);

_instigatorWeapon = "Unknown";
if (_instigator isKindOf "CAManBase") then {
	_instigatorWeapon = getText (configFile >> "CfgWeapons" >> currentWeapon _instigator >> "displayName");
};

[":EVENT:",[_frame, "killed", _unitID, [_instigatorID, _instigatorWeapon], _dist]] call socap_fnc_Post;