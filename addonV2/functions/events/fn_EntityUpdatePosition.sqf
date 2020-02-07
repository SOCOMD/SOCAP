params["_entity", "_frame"];

_id = _entity getVariable["socap_entity_id", -1];
if(_id < 0) exitWith {};

_pos = getPosATL _entity;
_dir = round getDir _entity;

if(_entity isKindOf "Man") then {
	_pos resize 2;
	_inVehicle = 0;
	if(!((vehicle _entity) isEqualTo _entity)) then {
		_inVehicle = 1;
	};

	_isAlive = 0;
	if(alive _entity) then {
		if(_entity getVariable ["ACE_isUnconscious", false]) then {
			_isAlive = 1;
		};
	};

	_isPlayer = 0;
	if(isPlayer _entity) then {
		_isPlayer = 1;
	};

	[":UPDATE:UNIT:",[_id, _pos, _dir, _isAlive, _inVehicle, _isPLayer]] call socap_fnc_Post;
} else {
	_pos set [2, round(_pos select 2)];
	_isAlive = 0;
	if(alive _entity) then {
		_isAlive = 1;
	};

	_crew = [];
	{
		_crewID = _x getVariable["socap_entity_id", -1];
		if(_crewID >= 0) then {
			_crew append [_crewID];
		};
	}foreach (crew _entity);

	[":UPDATE:VEH:",[_id, _pos, _dir, _isAlive, _crew]] call socap_fnc_Post;
};