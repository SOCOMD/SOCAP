params ["_entity"];

waitUntil {
	socap_global_captureEnabled;
};

_check = _this call socap_fnc_ValidateEntity;
if((_check select 0) == 0) exitWith {};
_class = _check select 1;

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
	_name = getText (configFile >> "CfgVehicles" >> _vehType >> "displayName");
	_entity setVariable["socap_entity_id", _id, true];
	[":NEW:VEH:",[_frame, _id, _class, _name], true] call socap_fnc_Post;
};

_entity addEventHandler["Fired", {_this spawn socap_fnc_EntityFired;}];
_entity addEventHandler["Hit", {_this spawn socap_fnc_EntityHit;}];

_entities = missionNamespace getVariable["socap_entities", []];
_entities append [_entity];
missionNamespace setVariable["socap_entities", _entities, true];