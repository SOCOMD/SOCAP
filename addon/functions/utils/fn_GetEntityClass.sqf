params ["_entity"];
if (_entity isKindOf "CAManBase") exitWith {"man"};
if (_entity isKindOf "Truck_F") exitWith {"truck";}; // Should be higher than Car
if (_entity isKindOf "Wheeled_APC_F") exitWith {"apc"}; // Should be higher than Car
if (_entity isKindOf "Car") exitWith {"car"};
if (_entity isKindOf "Tank") exitWith {"tank"};
if (_entity isKindOf "StaticMortar") exitWith {"static-mortar"};
if (_entity isKindOf "StaticWeapon") exitWith {"static-weapon"};
if (_entity isKindOf "ParachuteBase") exitWith {"parachute"};
if (_entity isKindOf "Helicopter") exitWith {"heli"};
if (_entity isKindOf "Plane") exitWith {"plane"};
if (_entity isKindOf "Air") exitWith {"plane"};
if (_entity isKindOf "Ship") exitWith {"sea"};
"unknown";