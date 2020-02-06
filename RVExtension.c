#include <stdlib.h>

extern void goRVExtension(char *output, size_t outputSize, char *input);
extern void goRVExtensionVersion(char *output, size_t outputSize);
extern void goRVExtensionArgs(char* output, size_t outputSize, char* input, char** argv, int argc);

#if defined(_MSC_VER)
    //  Microsoft 
    #define DLL_PUBLIC __declspec(dllexport)
#else
    //  GCC
    #define DLL_PUBLIC __attribute__((visibility("default")))
#endif

DLL_PUBLIC void _RVExtension(char *output, size_t outputSize, char *input) {
  goRVExtension(output, outputSize, input);
}

DLL_PUBLIC void _RVExtensionVersion(char *output, size_t outputSize) {
  goRVExtensionVersion(output, outputSize);
}

DLL_PUBLIC void _RVExtensionArgs(char* output, size_t outputSize, char* input, char** argv, int argc) {
  goRVExtensionArgs(output, outputSize, input, argv, argc);
}