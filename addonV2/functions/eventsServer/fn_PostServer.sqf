params ["_func", "_args"];

if(!isServer) exitWith {};

"socap" callExtension [_func, _args];