#include "predefined.hpp"

class CfgPatches {
    class socap {
        name = "socap";

        requiredAddons[] = {
			"A3_Functions_F",
            "cba_main"
		};

        units[] = {};
        weapons[] = {};
    };
};

#include "configs\CfgFunctions.hpp"
#include "configs\CfgVehicles.hpp"
#include "configs\CfgExtendedEventHandlers.hpp"