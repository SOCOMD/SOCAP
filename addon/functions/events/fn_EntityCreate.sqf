params ["_entity"];

if(!(local _entity)) exitWith {};

if((!(isNil "ASORGS_Clone")) && _entity isEqualTo ASORGS_Clone) exitWith {};

if((!(isNil "ASORVS_Clone")) && _entity isEqualTo ASORVS_Clone) exitWith {};

_this remoteExec ["socap_fnc_EntityCreateServer", 2];