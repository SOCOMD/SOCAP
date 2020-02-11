params ["_unit", "_id", "_uid", "_name"];

if(!isServer) exitWith {};

if(socap_global_captureEnabled isEqualTo false) exitWith {};

_frame = socap_global_frame;
[":EVENT:",[_frame, "disconnected", _name]] call socap_fnc_Post;