addMissionEventHandler["HandleDisconnect", {
	_this call ocap_fnc_eh_disconnected;
}];

addMissionEventHandler["PlayerConnected", {
	_this call ocap_fnc_eh_connected;
}];

addMissionEventHandler ["EntityKilled", {
	_this call ocap_fnc_eh_killed;
}];

// Add event saving markers
["INIT"] call ocap_fnc_handleMarkers;

ocap_eh_ended = addMissionEventHandler ["Ended", {
	_description = "";
	
	{
		_config = _x >> "CfgDebriefing";
		{
			_description = getText(_config >> _this >> _x);
			if !(_description isEqualTo "") exitWith {};
		} forEach ["subtitle", "title", "description"];

		if !(_description isEqualTo "") exitWith {};

	} forEach [missionConfigFile, configFile];

	[sideEmpty, _description] call ocap_fnc_exportData;
}];