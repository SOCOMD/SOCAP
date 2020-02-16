params ["_entity"];

if(!isServer) exitWith {};

if(socap_global_captureEnabled isEqualTo false) exitWith {};
socap_global_entities = socap_global_entities - [_entity];