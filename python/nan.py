

# 生成
print(float('NaN'))

# 真偽
print(True if float('NaN') else False)

# 算術
print(float('NaN') + 1)
print(float('NaN') * 1)

# キャスト
# int(float('NaN'))
# Traceback (most recent call last):
#   File "<stdin>", line 1, in <module>
# ValueError: cannot convert float NaN to integer

print(float(float('NaN')))

# NaNであるかチェック（自分自身と比較する）
n = float('NaN')
print(True if n == n else False)
