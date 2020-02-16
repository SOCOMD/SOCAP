#include "\socap\predefined.hpp"

if(!isServer) exitWith {};

if(socap_global_captureEnabled isEqualTo true) exitWith {};

_author = getMissionConfigValue ["author", "Unknown"];
[":START:", [worldName, briefingName, _author , FRAME_INTERVAL]] call socap_fnc_Post;

socap_global_entity_id = 0;
socap_global_frame = 0;
socap_global_captureEnabled = true;
socap_global_entities = [];

{
	[_x] call socap_fnc_EntityCreateServer;
} forEach (entities [[], ["Logic"], true]);

socap_global_mainLoopHandle = [] spawn socap_fnc_MainLoop;