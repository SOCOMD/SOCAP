params ["_entity"];

if(!isServer) exitWith {};

_entities = missionNamespace getVariable["socap_entities", []];
_entities = _entities - [_entity];
missionNamespace setVariable["socap_entities", _entities];