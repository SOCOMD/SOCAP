if(isServer) then {
	[] call socap_fnc_CaptureEndServer;
} else {
	[] remoteExec ["socap_fnc_CaptureEndServer", 2];
}