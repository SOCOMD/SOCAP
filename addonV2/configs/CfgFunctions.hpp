class CfgFunctions {
	class socap {
		class Events {
			file="\socap\functions\events";
			class CaptureEnd {};
			class EntityCreate {};
			class EntityFired {};
			class EntityHit {};
			class Post {};
		};	
		class EventsServer {
			file="\socap\functions\eventsServer";
			class CaptureStartServer {};
			class CaptureEndServer {};
			class EntityCreateServer {};
			class EntityFiredServer {};
			class EntityHitServer {};
			class EntityKilledServer {};
			class EntityUpdatePositionServer {};
			class PlayerConnectServer {};
			class PlayerDisconnectServer {};
			class PostServer {};
		};

		class Init {
			file="\socap\functions\init";
			class InitLocal {};
			class InitServer { preInit = 1; };
		};

		class Modules {
			file="\socap\functions\modules";
			class ModuleEndMission {};
		};

		class Systems {
			file="\socap\functions\systems";
			class MainLoop {};
			class ProcessEvents {};
		};

		class Utils {
			file="\socap\functions\utils";
			class GetEntityClass {};
			class ValidateEntity {};
		};
	};
};