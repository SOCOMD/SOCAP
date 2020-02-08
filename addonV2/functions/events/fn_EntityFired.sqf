params ["_unit", "_weapon", "_muzzle", "_mode", "_ammo", "_magazine", "_projectile", "_gunner"];

if(isServer) then {
	_this call socap_fnc_EntityFiredServer;
} else {
	_this remoteExec ["socap_fnc_EntityFiredServer", 2];
};