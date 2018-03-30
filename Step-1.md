# Step-1
Step-1ではVPCを作成し、WordpressがインストールされたAMIを用いてEC2インスタンスを起動します。

## Question VPCとは
VPCについて調べてみましょう(10分)

## VPCの作成
**実際にVPCを作成してみましょう。まずはサービスタブを選択しVPC管理ページを開きましょう**

![vpc-1](./images/step-1/vpc-1.png "VPC1")

----

**下にスクロールしVPCを選択します**

![vpc-2](./images/step-1/vpc-2.png "VPC2")

----

**VPCウィザードの開始を選択します**

![vpc-3](./images/step-1/vpc-3.png "VPC3")

----

**「ステップ1:VPC設定の選択」では「1個のパブリックサブネットを持つVPC」タブから選択ボタンを押下**

![vpc-4](./images/step-1/vpc-4.png "VPC4")

----

**「ステップ2:1個のパブリックサブネットを持つVPC」では以下を入力しVPCの作成ボタンを押下**

**VPC名: vpc-ユーザ名(例 vpc-user05 )**  
**アベイラビリティゾーン: ap-northeast-1d**  

![vpc-5](./images/step-1/vpc-5.png "VPC5")

----

**「VPCが正常に作成されました」ではOKボタンを押下**

![vpc-6](./images/step-1/vpc-6.png "VPC6")

----

**「VPCダッシュボード」では直下の「VPCでフィルタリング」でユーザ名を入力しフィルタリングしましょう。以下の例ではuser05でフィルタリングしています**

![vpc-7](./images/step-1/vpc-7.png "VPC7")

----

**作成したVPCの設定が正しいかVPCタブをクリックし内容を確認しましょう**

**名前: vpc-ユーザ名**  
**IPv4 CIDR: 10.0.0.0/16**  

![vpc-8](./images/step-1/vpc-8.png "VPC8")

----

**ウィザードで作成したサブネットを確認しましょう。VPC作成ウィザードでは、VPC自体と一緒に1つ目のサブネットも作成されます。**

![vpc-9](./images/step-1/vpc-9.png "VPC9")

----

**作成したサブネットのRoute Tableを確認しましょう。VPCのネットワークアドレス 10.0.0.0/16 のターゲットがlocalに、デフォルトルートの 0.0.0.0/0 のターゲットがインタネットゲートウェイ(igw-XXXX)になっており、インターネットと通信できる設定になっています。**
 
![vpc-10](./images/step-1/vpc-10.png "VPC10")

----

## サブネットの追加作成

**作成したVPCに対して追加でサブネットを加えましょう。ここではサブネットの3つ作成を行いましょう**

![subnet-1](./images/step-1/subnet-1.png "SUBNET1")

----
**1から4を以下を参考に設定しましょう。**

|-|1 名前タグ|2 VPC|3 アベイラビリティ ゾーン|4 CIDR ブロック|
|:-|:-|:-|:-|:-|
|2つ目|パブリックサブネット|自分で作成したVPCを指定|ap-northeast-1c|10.0.1.0/24|
|3つ目|プライベートサブネット|自分で作成したVPCを指定|ap-northeast-1d|10.0.2.0/24|
|4つ目|プライベートサブネット|自分で作成したVPCを指定|ap-northeast-1c|10.0.3.0/24|

![subnet-2](./images/step-1/subnet-2.png "SUBNET2")

----
**全てのサブネットを確認しましょう。4つ作成され赤枠の内容が設定通りか確認しましょう。その際にIPv4でソートすると見易いです。**

![subnet-3](./images/step-1/subnet-3.png "SUBNET3")

----
**パブリックサブネット(10.0.1.0/24)のルートテーブルを変更しましょう。パブリックサブネット(10.0.1.0/24)を選択しルートテーブルタブを選択、編集ボタンを押下**

![subnet-4](./images/step-1/subnet-4.png "SUBNET4")

----
**現在使用しているルートテーブル以外にもう一つ選択できるはずです。そちらを選択し保存ボタンにて保存しましょう。選択するとインターネットゲートウェイの設定が追加されます。**

![subnet-5](./images/step-1/subnet-5.png "SUBNET5")

----

## EC2インスタンスの作成
**サービスタブからEC2を選択**

![ec2-1](./images/step-1/ec2-1.png "EC21")

----
**Webサーバの作成を行いましょう。インスタンスの作成ボタンを押下**

![ec2-2](./images/step-1/ec2-2.png "EC22")

----
**マイAIMタブ、「1Day-AMI」の選択ボタンを押下**

![ec2-3](./images/step-1/ec2-3.png "EC23")

----

![ec2-4](./images/step-1/ec2-4.png "EC24")

----

![ec2-5](./images/step-1/ec2-5.png "EC25")

----

![ec2-6](./images/step-1/ec2-6.png "EC26")

----

![ec2-7](./images/step-1/ec2-7.png "EC27")

----

![ec2-8](./images/step-1/ec2-8.png "EC28")

----

![ec2-9](./images/step-1/ec2-9.png "EC29")

----

![ec2-10](./images/step-1/ec2-10.png "EC210")

----

![ec2-11](./images/step-1/ec2-11.png "EC211")

----

![ec2-12](./images/step-1/ec2-12.png "EC212")

----

![ec2-13](./images/step-1/ec2-13.png "EC213")

----
**作成したEC2インスタンスにて赤枠で囲った「パブリックDNS(IPv4)」の値をコピーしましょう。サーバログインのドメイン、WordpressのURLとなります。**

![ec2-14](./images/step-1/ec2-14.png "EC214")

----

## Question AMIとは
AMIについて調べてみましょう(10分)

## サーバログイン
**作成したEC2インスタンスにログインし環境の確認をしましょう。`1day-userXX.pem`は各自が作成した秘密鍵、ec2-XXXXXX.comは「パブリックDNS(IPv4)」の値です。「ec2-user@ip-10-0-0-XX」のプロンプトが返れば成功です**

```
$ chmod 600 1day-userXX.pem
$ ssh -i 1day-userXX.pem -o StrictHostKeyChecking=no ec2-user@ec2-XXXXXX.com
[ec2-user@ip-10-0-0-65 ~]$
```

**MySQLにrootユーザで接続し、所有しているデータベースの確認をしてみましょう**

**パスワードは`wordpress`**

```
$ mysql -u root -p
Enter password:

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| test               |
| wordpress          |
+--------------------+
5 rows in set (0.00 sec)
```

## Question
1.このサーバに振られたIPアドレスの確認をしましょう。  
2.このサーバのデフォルトゲートウェイを確認しましょう。

## Wordpressの初期設定
**「パブリックDNS(IPv4)」の値でブラウザを開きましょう。Wordpressのサイトが開けば作成成功です。初期設定では「日本語」を選択し続けるボタンを押下**

![wordpress-1](./images/step-1/wordpress-1.png "Wordpress1")

----
**「さあ、始めましょう」を押下**

![wordpress-2](./images/step-1/wordpress-2.png "Wordpress2")

----

![wordpress-3](./images/step-1/wordpress-3.png "Wordpress3")

----

![wordpress-4](./images/step-1/wordpress-4.png "Wordpress4")

----

![wordpress-5](./images/step-1/wordpress-5.png "Wordpress5")

----

![wordpress-6](./images/step-1/wordpress-6.png "Wordpress6")

----
**ここまでの設定が間違いないか、ユーザ、パスワードを設定しログインしましょう**

![wordpress-7](./images/step-1/wordpress-7.png "Wordpress7")

----
**管理画面にログインできれば設定完了です**

![wordpress-8](./images/step-1/wordpress-8.png "Wordpress8")

----
**「パブリックDNS(IPv4)」の値でブラウザを開きましょう。Wordpressのサイトが開けば設定完了です**

![wordpress-9](./images/step-1/wordpress-9.png "Wordpress9")

----

**ここまでのオペレーションでStep1は完了です！**
