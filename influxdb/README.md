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


## データベース作成

```
> create database climbing

> show databases
name: databases
name
----
_internal
climbing
```


## InfluxDBの用語

|用語|内容|
|---|---|
|measurements|RDBのテーブルに該当する|
