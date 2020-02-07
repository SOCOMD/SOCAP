#include "\socap\predefined.hpp"

_author = getMissionConfigValue ["author", "Unknown"];
[":START:", [worldName, briefingName, _author , FRAME_INTERVAL]] call socap_fnc_Post;