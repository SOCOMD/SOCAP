params ["_unit", "_id", "_uid", "_name"];

if(!isServer) exitWith {};

_frame = socap_global_frame;
[":EVENT:",[_frame, "disconnected", _name]] call socap_fnc_Post;