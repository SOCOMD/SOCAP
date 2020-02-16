class CfgFunctions {
	class socap {
		class Events {
			file="\socap\functions\events";
			class CaptureEnd {};
			class EntityCreate {};
			class EntityKilled {};
			class EntityUpdatePosition {};
			class Post {};
		};	
		class EventsServer {
			file="\socap\functions\eventsServer";
			class CaptureEndServer {};
			class CaptureStartServer {};
			class EntityCreateServer {};
			class EntityDeletedServer {};
			class EntityFiredServer {};
			class EntityHitServer {};
			class PlayerConnectServer {};
			class PlayerDisconnectServer {};
			class PostServer {};
		};

		class Init {
			file="\socap\functions\init";
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