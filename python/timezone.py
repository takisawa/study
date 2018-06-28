
from datetime import datetime
from datetime import timezone
from datetime import timedelta
from datetime import date

# 環境変数 TZ で制御できる
print(datetime.now())
print(date.today())

# 常にUTCの時刻を表示
print(datetime.now(timezone.utc))
print(datetime.utcnow())

# タイムゾーンの生成
JST = timezone(timedelta(hours=9), 'JST')

# 常にJSTの時刻を表示
print(datetime.now(JST))
print(datetime.now(JST).date())

