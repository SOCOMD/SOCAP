#define REQUIRED_VERSION	2.0
#define LOG(_args)			[":LOG:", _args] call ocap_fnc_extension
#define BOOL(_cond)			(parseNumber (_cond))
#define SQF2JSON(_str)		([_str, """", "'"] call CBA_fnc_replace)

#define ARR2(_arg1, _arg2) [_arg1, _arg2]
#define ARR3(_arg1, _arg2, _arg3) [_arg1, _arg2, _arg3]
#define ARR4(_arg1, _arg2, _arg3, _arg4) [_arg1, _arg2, _arg3, _arg4]
#define ARR5(_arg1, _arg2, _arg3, _arg4, _arg5) [_arg1, _arg2, _arg3, _arg4, _arg5]
#define ARR6(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6) [_arg1, _arg2, _arg3, _arg4, _arg5, _arg6]