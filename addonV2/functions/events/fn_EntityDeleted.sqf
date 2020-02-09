params ["_entity"];

if(isServer) then {
	_this call socap_fnc_EntityDeletedServer;
} else {
	_this remoteExec ["socap_fnc_EntityDeletedServer", 2];
};