## 概要

InfluxDBの検証用のgoプログラム。
例によって登山アプリをベースにする。

## テーブル構成

- 登山ID
- time
- 位置情報


## ターミナルからの接続

```
% influx -precision rfc3339
```


## データベース

```
> create database climbing
> show databases
name: databases
name
----
_internal
climbing

> create retention policy "1day" on "climbing" DURATION 1d REPLICATION 1
> SHOW RETENTION POLICIES ON "climbing"
name    duration shardGroupDuration replicaN default
----    -------- ------------------ -------- -------
autogen 0s       168h0m0s           1        true
1day    24h0m0s  1h0m0s             1        false
```


## 参照

- InfluxDB key concepts
  - https://docs.influxdata.com/influxdb/v1.7/concepts/key_concepts/
  - database / measurements / tag / field / retention policy などの説明がある
