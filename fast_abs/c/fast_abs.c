#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <sys/time.h>

const unsigned long int_size = sizeof(int) * 8;


int fast_abs(int x) {
  int m, n, r;

  // xの最左bitは、正数の場合に0、負数の場合に1である。
  //
  // int_size - 1 右シフトを行うと
  // mは、正数の場合に0、負数の場合に-1となる（gccでは算術シフトするため）。
  m = x >> (int_size - 1);

  // xとmのXORを求める。
  // xが正の場合、m=0のため、nはxと同じ値。
  // xが負の場合、m=-1であり、nはxをbit反転した値。
  return (x ^ m) - m;
  n = x ^ m;

  // xが正の場合、m=0のため、rはn(x)と同じ値。
  // xが負の場合、m=-1であり、rはxの符号を反転した数となる。
  r = n - m;

  return r;
}

int slow_abs(int x) {
  if (x < 0) {
    x = -x;
  }

  return x;
}

const int n = 1000000;

int main(int argc, char *argv[]) {
  // if (argc != 2) {
  //   puts("Invalid arguments");
  // }

  // int x = atoi(argv[1]);
  //

  long int sum = 0;
  struct timeval startTime, endTime;  // 構造体宣言
  clock_t startClock, endClock;       // clock_t型変数宣言
  time_t diffsec;
  double realsec, cpusec;
  suseconds_t diffsub;

  gettimeofday(&startTime, NULL);     // 開始時刻取得
  startClock = clock();               // 開始時刻のcpu時間取得


  for (int i = -n; i < n; i++) {
    fast_abs(i);
  }

  gettimeofday(&endTime, NULL);       // 開始時刻取得
  endClock = clock();                 // 開始時刻のcpu時間取得

  diffsec = difftime(endTime.tv_sec, startTime.tv_sec);    // 秒数の差分を計算
  diffsub = endTime.tv_usec - startTime.tv_usec;      // マイクロ秒部分の差分を計算

  realsec = diffsec+diffsub*1e-6;                          // 実時間を計算
  cpusec = (endClock - startClock)/(double)CLOCKS_PER_SEC; // cpu時間を計算

  printf("%f\n", realsec);
  printf("%f\n", cpusec);

  gettimeofday(&startTime, NULL);     // 開始時刻取得
  startClock = clock();               // 開始時刻のcpu時間取得


  for (int i = -n; i < n; i++) {
    slow_abs(i);
  }

  gettimeofday(&endTime, NULL);       // 開始時刻取得
  endClock = clock();                 // 開始時刻のcpu時間取得

  diffsec = difftime(endTime.tv_sec, startTime.tv_sec);    // 秒数の差分を計算
  diffsub = endTime.tv_usec - startTime.tv_usec;      // マイクロ秒部分の差分を計算

  realsec = diffsec+diffsub*1e-6;                          // 実時間を計算
  cpusec = (endClock - startClock)/(double)CLOCKS_PER_SEC; // cpu時間を計算

  printf("%f\n", realsec);
  printf("%f\n", cpusec);


  return (int)sum;
}
