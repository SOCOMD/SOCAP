params ["_id", "_uid", "_name", "_jip", "_owner", "_idstr"];

if(!isServer) exitWith {};

if(socap_global_captureEnabled isEqualTo false) exitWith {};

_frame = socap_global_frame;
[":EVENT:",[_frame, "connected", _name]] call socap_fnc_Post;