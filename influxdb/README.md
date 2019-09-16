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


## InfluxDBの基本操作

### データベース作成

```
> create database climbing
> show databases
name: databases
name
----
_internal
climbing
> use climbing
Using database climbing
```

### データベース削除

```
> drop database climbing
> show databases
name: databases
name
----
_internal
```


### retention policyの作成

```
> create retention policy "1day" on "climbing" DURATION 1d REPLICATION 1
> SHOW RETENTION POLICIES ON "climbing"
name    duration shardGroupDuration replicaN default
----    -------- ------------------ -------- -------
autogen 0s       168h0m0s           1        true
1day    24h0m0s  1h0m0s             1        false
```


## CSVからのデータ登録

参照: https://www.influxdata.com/blog/how-to-write-points-from-csv-to-influxdb/

```
Finally, I highly recommend using a Telegraf plugin to write points
to your InfluxDB database because Telegraf is written in Go.
```


## 参照

- InfluxDB key concepts
  - https://docs.influxdata.com/influxdb/v1.7/concepts/key_concepts/
  - database / measurements / tag / field / retention policy などの説明がある


## 疑問点

- Unique制約はあるのか
- replicationの仕組み
