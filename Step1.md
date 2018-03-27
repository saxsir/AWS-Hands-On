# Step1

## リージョン選択
画像右上のリージョン選択から「アジアパシフィック（東京）」を選択

![region-select](./images/region-select.png "リージョン選択")

東京が選択されていること

![region-tokyo](./images/region-tokyo.png "リージョン東京")

## Question VPCとは
VPCについて調べてみましょう(10分)

## VPCの作成
実際にVPCを作成してみましょう。まずはサービスタブを選択しVPC管理ページを開きましょう

![vpc-1](./images/vpc-1.png "VPC1")

VPCを選択します

![vpc-2](./images/vpc-2.png "VPC2")

VPCウィザードの開始を選択します

![vpc-3](./images/vpc-3.png "VPC3")

「ステップ1:VPC設定の選択」では「1個のパブリックサブネットを持つVPC」タブから選択ボタンを押下

![vpc-4](./images/vpc-4.png "VPC4")

「ステップ2:1個のパブリックサブネットを持つVPC」では以下を入力しVPCの作成ボタンを押下

VPC名: vpc-ユーザ名(例 vpc-user05 )  
アベイラビリティゾーン: ap-northeast-1d  

![vpc-5](./images/vpc-5.png "VPC5")

「VPCが正常に作成されました」ではOKボタンを押下

「VPCダッシュボード」では直下の「VPCでフィルタリング」でユーザ名を入力しフィルタリングしましょう

![vpc-6](./images/vpc-6.png "VPC6")

![vpc-7](./images/vpc-7.png "VPC7")
