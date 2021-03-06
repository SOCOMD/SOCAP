#include "predefined.hpp"

class CfgPatches {
    class socap {
        name = "socap";
		versionStr = "1.0.0";

        requiredAddons[] = {
			"A3_Functions_F",
            "cba_main"
		};

        units[] = {
			"ModuleEndMissionOCAP"
		};
		
        weapons[] = {};
    };
};

#include "configs\CfgFunctions.hpp"
#include "configs\CfgVehicles.hpp"
#include "configs\CfgExtendedEventHandlers.hpp"