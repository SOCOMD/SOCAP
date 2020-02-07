params ["_func", "_args"];

if(!isServer) exitWith {};

"socap" callExtension [_func, _args];

systemChat format["%1:%2", _func, _args];