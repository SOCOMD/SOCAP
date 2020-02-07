class CfgFunctions {
	class socap {
		class Events {
			file="\socap\functions\events";
			class CaptureEnd {};
			class CaptureStart {};
			class EntityCreate {};
			class EntityFired {};
			class EntityHit {};
			class EntityKilled {};
			class EntityUpdatePosition {};
			class PlayerConnect {};
			class PlayerDisconnect {};
			class Post {};
			class PostServer {};
		};

		class Init {
			file="\socap\functions\init";
			class InitLocal {};
			class InitServer { preInit = 1; };
		};

		class Systems {
			file="\socap\functions\systems";
			class MainLoop {};
			class ProcessEvents {};
		};
	};
};