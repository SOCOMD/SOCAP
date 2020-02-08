#include "\socap\predefined.hpp"

_frame = socap_global_frame;
_sideWon = str sideEmpty;
_author = getMissionConfigValue ["author", ""];
_description = "";
[":EVENT:", [_frame, "endMission", [_sideWon, _description]]] call socap_fnc_Post;
[":SAVE:", [worldName, briefingName, _author, FRAME_INTERVAL, _frame]] call socap_fnc_Post;

missionNamespace setVariable["socap_capture_enabled", false, true];