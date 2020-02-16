params ["_func", "_args"];

if(!isServer) exitWith {};

if(socap_global_captureEnabled isEqualTo false) exitWith {};

"socap" callExtension [_func, _args];