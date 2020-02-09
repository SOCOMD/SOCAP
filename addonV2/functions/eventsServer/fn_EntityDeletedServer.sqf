params ["_entity"];

_entities = missionNamespace getVariable["socap_entities", []];
_entities = _entities - [_entity];
missionNamespace setVariable["socap_entities", _entities];