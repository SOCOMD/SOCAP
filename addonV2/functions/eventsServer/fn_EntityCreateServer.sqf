private _getClass = {
	if (_this isKindOf "Truck_F") exitWith {"truck"}; // Should be higher than Car
	if (_this isKindOf "Wheeled_APC_F") exitWith {"apc"}; // Should be higher than Car
	if (_this isKindOf "Car") exitWith {"car"};
	if (_this isKindOf "Tank") exitWith {"tank"};
	if (_this isKindOf "StaticMortar") exitWith {"static-mortar"};
	if (_this isKindOf "StaticWeapon") exitWith {"static-weapon"};
	if (_this isKindOf "ParachuteBase") exitWith {"parachute"};
	if (_this isKindOf "Helicopter") exitWith {"heli"};
	if (_this isKindOf "Plane") exitWith {"plane"};
	if (_this isKindOf "Air") exitWith {"plane"};
	if (_this isKindOf "Ship") exitWith {"sea"};
	"unknown"
};

params ["_entity"];

_valid = 1;
if(_entity isKindOf "Man") then {
	if(name _entity == "") exitWith {
		_valid = 0;
	};
};

if(_valid == 0) exitWith {};

waitUntil {
	missionNamespace getVariable["socap_capture_enabled", false];
};


_id = socap_global_entity_id;
socap_global_entity_id = socap_global_entity_id + 1;
_frame = socap_global_frame;

if(_entity isKindOf "Man") then {
	_name = name _entity;
	_groupID = groupID (group _entity);
	_side = str side _entity;

	_isPlayer = 0;
	if(isPlayer _entity) then {
		_isPlayer = 1;
	};

	_entity setVariable["socap_entity_id", _id, true];
	[":NEW:UNIT:",[_frame, _id, _name, _groupID, _side, _isPlayer], true] call socap_fnc_Post;
} else {
	_vehType = typeOf _entity;
	_class = _vehType call _getClass;
	_name = getText (configFile >> "CfgVehicles" >> _vehType >> "displayName");

	_entity setVariable["socap_entity_id", _id, true];
	[":NEW:VEH:",[_frame, _id, _class, _name], true] call socap_fnc_Post;
};

_entity addEventHandler["Fired", {_this remoteExec ["socap_fnc_EntityFiredServer", 2];}];
_entity addEventHandler["Hit", {_this remoteExec ["socap_fnc_EntityHitServer", 2];}];

_entities = missionNamespace getVariable["socap_entities", []];
_entities append [_entity];
missionNamespace setVariable["socap_entities", _entities, true];