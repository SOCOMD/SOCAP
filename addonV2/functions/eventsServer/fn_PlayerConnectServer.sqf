params ["_id", "_uid", "_name", "_jip", "_owner", "_idstr"];

_frame = socap_global_frame;
[":EVENT:",[_frame, "connected", _name]] call socap_fnc_Post;