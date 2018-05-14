#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <sys/time.h>

const unsigned long int_size = sizeof(int) * 8;

int fast_abs(int x) {
  int m;

  // (1) x を int_size - 1 回 右シフトする
  //
  // ※ 前提として、8bitでの2進数での数値表現
  //   1 -> 00000001
  //   2 -> 00000010
  //  10 -> 00001010
  //  -1 -> 11111111 負の場合、対応する正の値をbit反転し、+1した値となる
  //  -2 -> 11111110
  // -10 -> 11110100
  //
  // int_size - 1 回 右shiftを行うと m は次の値となる
  //   - x >= 0 の場合 m =  0 （全てのbitが0）
  //   - x <  0 の場合 m = -1 （全てのbitが1）※ gccでは右シフトは算術シフト（符号ビットは維持）
  m = x >> (int_size - 1);

  // // x と m の XOR を求める。
  // // x >= 0 の場合、m の全てのbitが 0 のため、x XOR m しても x の bit は変わらない（x == n）
  // // x <  0 の場合、m の全てのbitが 1 のため、x XOR m すると n は x の bit を反転した値となる
  // n = x ^ m

  // // x >= 0 の場合、m ==  0 のため、n - m は x - 0 となり、x はそのまま
  // // x <  0 の場合、m == -1 のため、n - m は x の bit 反転 - (-1) となり、x の符号を反転した値となる
  // r = n - m;

  // return r;
  return (x ^ m) - m;
}

int slow_abs(int x) {
  if (x < 0) {
    return - x;
  }

  return x;
}

void benchmark(int (*func)(int), int n) {
  struct timeval startTime, endTime;  // 構造体宣言
  clock_t startClock, endClock;       // clock_t型変数宣言

  gettimeofday(&startTime, NULL);     // 開始時刻取得
  startClock = clock();               // 開始時刻のcpu時間取得

  int sum = 0;
  for (int i=-n; i < n; i++) {
    sum += func(i);
  }

  gettimeofday(&endTime, NULL);       // 終了時刻取得
  endClock = clock();                 // 終了時刻のcpu時間取得

  time_t diffsec = difftime(endTime.tv_sec, startTime.tv_sec);    // 秒数の差分を計算
  suseconds_t diffsub = endTime.tv_usec - startTime.tv_usec;      // マイクロ秒部分の差分を計算

  double realsec = diffsec+diffsub*1e-6;                          // 実時間を計算
  double cpusec = (endClock - startClock)/(double)CLOCKS_PER_SEC; // cpu時間を計算

  printf("cpusec: %f, realsec: %f\n", cpusec, realsec);
  printf("sum: %d\n", sum);
}


int main(int argc, char *argv[]) {
  if (argc != 2) {
    puts("Invalid arguments");
  }

  int n = atoi(argv[1]);
  puts("--- fast_abs");
  benchmark(fast_abs, n);
  puts("--- slow_abs");
  benchmark(slow_abs, n);
}
