# Step-3
Step-3ではWebサーバとアプリケーションレイヤの水平分散、データベースレイヤの冗長化を行います。具体的にはStep-2の環境にロードバランサを配置しEC2インスタンスを2つのアベイラビリティゾーンに水平分散、プライベートサブネットに配置されたデータベースAuroraを同様に2つのアベイラビリティゾーンに配置(MultiA-Z構成)し冗長構成とします。

## Question 水平分散とは
水平分散について調べてみましょう(10分)

## Question MultiA-Zとは
AWS MultiA-Zについて調べてみましょう(10分)

## AuroraをMultiA-Z構成に変更
**サービスからRDSを選択**

![multiaz-1](./images/step-3/multiaz-1.png "MULTIAZ1")

----
