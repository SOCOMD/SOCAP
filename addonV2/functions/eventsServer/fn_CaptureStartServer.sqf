#include "\socap\predefined.hpp"

_author = getMissionConfigValue ["author", "Unknown"];
[":START:", [worldName, briefingName, _author , FRAME_INTERVAL]] call socap_fnc_Post;

systemChat format[":START:"];

socap_global_entity_id = 0;
socap_global_frame = 0;
socap_global_captureEnabled = true;