#include "\ocap\script_macros.hpp"

//Scope
#define private	0
#define protected 1
#define public 2

class CfgPatches {
    class OCAP {
        name = "OCAP";
        author = "Dell, Zealot, Kurt";
        requiredAddons[] = {"A3_Functions_F",
                            "cba_main" /*TODO: Добавить WMT, SWT*/};
        requiredVersion = REQUIRED_VERSION;
        units[] = {
			"ModuleEndMissionOCAP"
		};
        weapons[] = {};
    };
};
class CfgFunctions {
    class OCAP {
        class null {
            file = "ocap\functions";
            class init {
                preInit = 1;
            };
            class startCaptureLoop {};
            class getDelay {};
            class addEventHandlers {};
            class addEventMission {};
            class eh_connected {};
            class eh_disconnected {};
            class eh_fired {};
            class eh_hit {};
            class eh_killed {};
            class exportData {};
            class extension {};
            class handleMarkers {};
			class moduleEndMission {};
        };
    };
};

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
		function = "ocap_fnc_moduleEndMission";
	};
};