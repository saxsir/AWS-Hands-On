# Step-2
Step-2では「パブリックサブネット」内のEC2単体で賄っていたサービスのうちMySQLを「プライベートサブネット」に「Amazon Aurora MySQL」を利用して切り出します。

## Question
RDS(Relational Database Service)、Auroraについて調べてみましょう(10分)

## DB用セキュリティグループの作成
**サービスからEC2を選択しましょう**

![security-1](./images/step-2/security-1.png "SECURITY1")

----
**セキュリティグループタグからセキュリティグループの作成ボタンを押下**

![security-2](./images/step-2/security-2.png "SECURITY2")

----

|項目|設定値|
|:-|:-|
|セキュリティグループ名|db-ユーザ名 (例 db-user05)|
|説明|RDS for Aurora|
|VPC|作成したVPCを指定|

![security-3](./images/step-2/security-3.png "SECURITY3")

----

![security-4](./images/step-2/security-4.png "SECURITY4")

----

![security-5](./images/step-2/security-5.png "SECURITY5")

----
