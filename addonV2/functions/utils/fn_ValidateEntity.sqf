params ["_entity"];

_class = _this call socap_fnc_GetEntityClass;
if(_class isEqualTo "unknown") exitWith {
	[0, _class];
};

if(_entity isKindOf "Man") then {
	if(name _entity == "") exitWith {
		[0, _class];
	};
};

[1, _class];