

# 生成
print(float('NaN'))

# 真偽
print(True if float('NaN') else False)

# 算術
print(float('NaN') + 1)
print(float('NaN') * 1)

# キャスト
try:
    int(float('NaN'))
except ValueError as err:
    print(err)

print(float(float('NaN')))

# NaNであるかチェック（自分自身と比較する）
n = float('NaN')
print(True if n == n else False)


import math
print(math.isnan(n))
print(math.isnan(1))
