# DynamoDB調査

## 更新日

2019-08-31

## 概要

- AWSの完全マネージド型サービス
  - https://aws.amazon.com/jp/dynamodb/
- マルチリージョン、マルチマスター
- 1日に10兆件以上のリクエストを処理することができ、毎秒2,000万件を超えるリクエストをサポート
- Apache CassandraはDynamoDBより高速と言われているが、AWSにはManagedサービスががない


## 料金体系

DynamoDBは下記のキャパシティモードがある。

- オンデマンドキャパシティモード
- プロビジョニング済みキャパシティモード


## 懸念事項

- 最大Valueサイズが400 KB
  - https://docs.aws.amazon.com/ja_jp/amazondynamodb/latest/developerguide/Limits.html


## Macでの開発環境の構築


```
% docker pull amazon/dynamodb-local
% docker run -p 8000:8000 amazon/dynamodb-local

```


## TODO

- トランザクション
- dataにappendする
- appendするときにunique制約のようなことができるのか？
- Valueサイズ調査
- goからアクセスする方法
