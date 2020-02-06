#include <stdlib.h>

extern void goRVExtension(char *output, size_t outputSize, char *input);
extern void goRVExtensionVersion(char *output, size_t outputSize);
extern void goRVExtensionArgs(char* output, size_t outputSize, char* input, char** argv, int argc);

#if defined(_WIN32) || defined(_WIN64)
// __declspec(dllexport) void __stdcall _RVExtension(char *output, size_t outputSize, char *input) {
//   goRVExtension(output, outputSize, input);
// }
void RVExtension(char *output, size_t outputSize, char *input) {
	goRVExtension(output, outputSize, input);
}
 
void RVExtensionVersion(char *output, size_t outputSize) {
	goRVExtensionVersion(output, outputSize);
}
 
void RVExtensionArgs(char* output, size_t outputSize, char* input, char** argv, int argc) {
	goRVExtensionArgs(output, outputSize, input, argv, argc);
}
#else
// __attribute__((visibility("default"))) void _RVExtension(char *output, size_t outputSize, char *input) {
//   goRVExtension(output, outputSize, input);
// }
__attribute__ ((visibility ("default"))) void RVExtension(char *output, size_t outputSize, char *input) {
	goRVExtension(output, outputSize, input);
}
 
__attribute__ ((visibility ("default"))) void RVExtensionVersion(char *output, size_t outputSize) {
	goRVExtensionVersion(output, outputSize);
}
 
__attribute__ ((visibility ("default"))) void RVExtensionArgs(char* output, size_t outputSize, char* input, char** argv, int argc) {
	goRVExtensionArgs(output, outputSize, input, argv, argc);
}
#endif

