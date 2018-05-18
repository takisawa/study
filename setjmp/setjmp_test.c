#include <stdio.h>
#include <setjmp.h>

jmp_buf env;

void func1(void) {
	if (setjmp(env) == 0) {
		puts("setjmp == 0");
	} else {
		puts("setjmp != 0");
	}
}

void func2(void) {
	puts("before longjmp");

	longjmp(env, 1);

	puts("after longjmp");
}

int main(void) {
	func1();
	func2();
}
