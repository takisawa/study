#include <stdio.h>

const unsigned long int_size = sizeof(int) * 8;

int main(void) {
  printf("sizeof(int)*8 = %lu\n", int_size);

  int x = -10000;

  // xの最左bitは、正数の場合に0、負数の場合に1である。
  //
  // int_size - 1 右シフトを行うと
  // mは、正数の場合に0、負数の場合に-1となる（gccでは算術シフトするため）。
  int m = x >> (int_size - 1);

  // xとmのXORを求める。
  // xが正の場合、m=0のため、nはxと同じ値。
  // xが負の場合、m=-1であり、nはxをbit反転した値。
  int n = x ^ m;

  // xが正の場合、m=0のため、rはn(x)と同じ値。
  // xが負の場合、m=-1であり、rはxの符号を反転した数となる。
  int r = n - m;

  printf("%d\n", r);

  return 0;
}
