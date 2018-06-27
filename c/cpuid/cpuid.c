#include <stdio.h>
#include <string.h>


struct abcd{
    unsigned int a,b,c,d;
};

void cpuid(struct abcd* r, unsigned int eax, unsigned int ecx){
    __asm__ volatile ("cpuid"
                      :"=a"(r->a), "=b"(r->b), "=c"(r->c), "=d"(r->d)
                      : "a"(eax), "c"(ecx));
}

int main(){
    struct abcd r;
    char buf[48];

    // GenuinIntel
    cpuid(&r, 0x0, 0x0);
    printf("0x%x\n", r.a);
    memcpy(buf, &r.b, 4);
    memcpy(buf+4, &r.d, 4);
    memcpy(buf+8, &r.c, 4);
    buf[12] = '\0';
    printf("%s\n", buf);

    // Processor Name, Freq
    cpuid(&r, 0x80000002, 0x0);
    memcpy(buf, &r, 16);
    cpuid(&r, 0x80000003, 0x0);
    memcpy(buf+16, &r, 16);
    cpuid(&r, 0x80000004, 0x0);
    memcpy(buf+32, &r, 16);
    printf("%s\n", buf);

    // SIMD
    // https://software.intel.com/en-us/articles/how-to-detect-new-instruction-support-in-the-4th-generation-intel-core-processor-family
    cpuid(&r, 0x1, 0x0);
    printf("MMX:    %s\n", r.d & 1 << 23 ? "OK" : "NG");
    printf("SSE:    %s\n", r.d & 1 << 25 ? "OK" : "NG");
    printf("AVX:    %s\n", r.c & 1 << 28 ? "OK" : "NG");
    printf("FMA:    %s\n", r.c & 1 << 12 ? "OK" : "NG");
    cpuid(&r, 0x7, 0x0);
    printf("AVX2:   %s\n", r.b & 1 <<  5 ? "OK" : "NG");

    return 0;
}
