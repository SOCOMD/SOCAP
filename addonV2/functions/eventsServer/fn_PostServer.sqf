params ["_func", "_args"];

if(!isServer) exitWith {};

if(!socap_global_captureEnabled) exitWith {};

"socap" callExtension [_func, _args];