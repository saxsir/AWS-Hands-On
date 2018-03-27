# Step-1
Step-1ではVPCを作成し、WordpressがインストールされたAMIを用いてEC2インスタンスを起動します。

## Question VPCとは
VPCについて調べてみましょう(10分)

## VPCの作成
**実際にVPCを作成してみましょう。まずはサービスタブを選択しVPC管理ページを開きましょう**

![vpc-1](./images/vpc-1.png "VPC1")

----

**VPCを選択します**

![vpc-2](./images/vpc-2.png "VPC2")

----

**VPCウィザードの開始を選択します**

![vpc-3](./images/vpc-3.png "VPC3")

----

**「ステップ1:VPC設定の選択」では「1個のパブリックサブネットを持つVPC」タブから選択ボタンを押下**

![vpc-4](./images/vpc-4.png "VPC4")

----

**「ステップ2:1個のパブリックサブネットを持つVPC」では以下を入力しVPCの作成ボタンを押下**

**VPC名: vpc-ユーザ名(例 vpc-user05 )**  
**アベイラビリティゾーン: ap-northeast-1d**  

![vpc-5](./images/vpc-5.png "VPC5")

----

**「VPCが正常に作成されました」ではOKボタンを押下**

![vpc-6](./images/vpc-6.png "VPC6")

----

**「VPCダッシュボード」では直下の「VPCでフィルタリング」でユーザ名を入力しフィルタリングしましょう。以下の例ではuser05でフィルタリングしています**

![vpc-7](./images/vpc-7.png "VPC7")

----

**作成したVPCの設定が正しいかVPCタブをクリックし内容を確認しましょう**

**名前: vpc-ユーザ名**  
**IPv4 CIDR: 10.0.0.0/16**  

![vpc-8](./images/vpc-8.png "VPC8")

----
