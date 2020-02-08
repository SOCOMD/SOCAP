params ["_unit", "_id", "_uid", "_name"];

_frame = missionNamespace getVariable["socap_frame", 0];
[":EVENT:",[_frame, "disconnected", _name]] call socap_fnc_Post;