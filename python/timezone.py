
from datetime import datetime
from datetime import timezone

# 環境変数 TZ で制御できる
print(datetime.now())

# 常にUTCの時刻を表示
print(datetime.now(timezone.utc))
