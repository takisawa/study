from pytz import timezone
from datetime import datetime

utc_now = datetime.now(timezone('UTC'))
print(utc_now)

jst = timezone('Asia/Tokyo')
print(jst)

print(datetime.now(jst))
