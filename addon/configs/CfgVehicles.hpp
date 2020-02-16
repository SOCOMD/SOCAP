class CfgVehicles {
	class Module_F;
	class ModuleEndMission_F : Module_F {
		scope = private;
		scopeCurator = private;
	};

	class ModuleEndMissionOCAP : ModuleEndMission_F {
		scope = public;
		scopeCurator = public;
		isGlobal = 1;
		
		_generalMacro = "ModuleEndMissionOCAP";
		function = "socap_fnc_ModuleEndMission";
		curatorInfoType = "";
		curatorInfoTypeEmpty = "";
	};
};