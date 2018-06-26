int main() {
	__asm__ __volatile__ (
		"xor %eax, %eax\n\t"
		"cpuid"
	);
	return 0;
}
