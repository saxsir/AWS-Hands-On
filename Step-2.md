# Step-2
Step-2では「パブリックサブネット」内のEC2単体で賄っていたサービスのうちMySQLを「プライベートサブネット」に「Amazon Aurora MySQL」を利用して切り出します。

## Question
RDS(Relational Database Service)、Auroraについて調べてみましょう(10分)

## DB用セキュリティグループの作成
**サービスからEC2を選択しましょう**

![security-1](./images/security-1.png "SECURITY1")

----
**セキュリティグループタグからセキュリティグループの作成ボタンを押下**

![security-1](./images/security-1.png "SECURITY1")

----

