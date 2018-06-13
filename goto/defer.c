#include <stdio.h>

void func() {
  FILE *fp;
  char buffer[1024];
  size_t s;

  fp = fopen("test.dat", "r");

  if (fp == NULL) {
    goto defer;
  }

  s = fread(buffer, 2, 3, fp);

  if (s == 0) {
    goto defer;
  }

  buffer[6] = '\0';
  printf("%s\n", buffer);
defer:
  if (fp != NULL) {
    fclose(fp);
  }

  printf("execute defer\n");
}

int main(void) {
  func();

  return 0;
}
