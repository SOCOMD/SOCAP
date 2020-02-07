params ["_id", "_uid", "_name", "_jip", "_owner", "_idstr"];

_frame = missionNamespace getVariable["socap_frame", 0];
[":EVENT:",[_frame, "connected", _name]] call socap_fnc_Post;