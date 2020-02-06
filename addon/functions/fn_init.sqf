//Settings
ocap_nameServer = 'SOCOMD';
ocap_minPlayerCount = 1;
ocap_frameCaptureDelay = 1.5;
ocap_excludeFromRecord = ["ACE_friesAnchorBar"];

//Vars
ocap_capture = false;
ocap_captureFrameNo = 0;

// Add event missions
call ocap_fnc_addEventMission;
[":START:", [worldName, briefingName, getMissionConfigValue ["author", ""], ocap_frameCaptureDelay]] call ocap_fnc_extension;
0 spawn ocap_fnc_startCaptureLoop;