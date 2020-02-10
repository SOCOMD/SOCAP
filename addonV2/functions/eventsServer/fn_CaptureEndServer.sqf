#include "\socap\predefined.hpp"

if(!isServer) exitWith {};

if(socap_global_captureEnabled isEqualTo false) exitWith {};

_frame = socap_global_frame;
_sideWon = str sideEmpty;
_author = getMissionConfigValue ["author", ""];
_description = "";
[":EVENT:", [_frame, "endMission", [_sideWon, _description]]] call socap_fnc_Post;
[":SAVE:", [worldName, briefingName, _author, FRAME_INTERVAL, _frame]] call socap_fnc_Post;

socap_global_captureEnabled = false;
socap_global_entity_id = 0;
socap_global_frame = 0;