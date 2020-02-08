params ["_unit", "_source", "_damage", "_instigator"];

if(isServer) then {
	_this call socap_fnc_EntityHitServer;
} else {
	_this remoteExec ["socap_fnc_EntityHitServer", 2];
};