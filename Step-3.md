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

![multiaz-2](./images/step-3/multiaz-2.png "MULTIAZ2")

----

![multiaz-3-1](./images/step-3/multiaz-3-1.png "MULTIAZ3-1")

----

![multiaz-3-2](./images/step-3/multiaz-3-2.png "MULTIAZ3-2")

----

![multiaz-3-3](./images/step-3/multiaz-3-3.png "MULTIAZ3-3")

----
**約5分ほどでAuroraのレプリカは作成されます。画面中央上のリロードボタンを押して「利用可能」になるまで待ちましょう**

![multiaz-4](./images/step-3/multiaz-4.png "MULTIAZ4")

----

## Auroraクラスターの確認

![aurora-cluster-1](./images/step-3/aurora-cluster-1.png "AURORA-CLUSTER1")

----
**クラスターエンドポイント、読み込みエンドポイントの確認をしましょう**

![aurora-cluster-2](./images/step-3/aurora-cluster-2.png "AURORA-CLUSTER3")

----

## 接続確認
**EC2サーバにSSH接続し、EC2サーバからAuroraに接続してみましょう。また作成したAuroraインスタンスが意図したセグメントに配置されているかも確認しましょう。**

```
$ ssh -i 1day-userXX.pem -o StrictHostKeyChecking=no ec2-user@ec2-XXXXXX.com
[ec2-user@ip-10-0-0-65 ~]$
```

**クラスタエンドポイントを使用してAuroraに接続しましょう。読み書きの権限についても確認しましょう。**

**注意 wp-userXX-cluster.cluster-cenae7eyijpr.ap-northeast-1.rds.amazonaws.comは各自のクラスタエンドポイントに直すこと。パスワードはAurora作成時に設定した内容を指定すること**

```
$ mysql -u admin -p -hwp-user05-cluster.cluster-cenae7eyijpr.ap-northeast-1.rds.amazonaws.com

mysql> select @read_only;
+------------+
| @read_only |
+------------+
| NULL       |
+------------+
1 row in set (0.00 sec)

mysql> exit
```

**続いてネットワークセグメントの確認(クラスタエンドポイント)をしましょう**

```
$ nslookup wp-user05-cluster.cluster-cenae7eyijpr.ap-northeast-1.rds.amazonaws.com
Server:     10.0.0.2
Address:    10.0.0.2#53

Non-authoritative answer:
wp-user05-cluster.cluster-cenae7eyijpr.ap-northeast-1.rds.amazonaws.com canonical name = wp-user05.cenae7eyijpr.ap-northeast-1.rds.amazonaws.com.
Name:   wp-user05.cenae7eyijpr.ap-northeast-1.rds.amazonaws.com
Address: 10.0.2.226
```

**読み込みエンドポイントを使用してAuroraに接続しましょう。読み書きの権限についても確認しましょう。**

**注意 wp-userXX-cluster.cluster-ro-cenae7eyijpr.ap-northeast-1.rds.amazonaws.comは各自の読み込みエンドポイントに直すこと。パスワードはAurora作成時に設定した内容を指定すること**

```
$ mysql -u admin -p -hwp-userXX-cluster.cluster-ro-cenae7eyijpr.ap-northeast-1.rds.amazonaws.com

mysql> select @read_only;
+------------+
| @read_only |
+------------+
| NULL       |
+------------+
1 row in set (0.01 sec)

mysql> exit
```

**続いてネットワークセグメントの確認(読み込みエンドポイント)をしましょう**

```
$ nslookup wp-user05-cluster.cluster-ro-cenae7eyijpr.ap-northeast-1.rds.amazonaws.com
Server:     10.0.0.2
Address:    10.0.0.2#53

Non-authoritative answer:
wp-user05-cluster.cluster-ro-cenae7eyijpr.ap-northeast-1.rds.amazonaws.com  canonical name = wp-user05-slave.cenae7eyijpr.ap-northeast-1.rds.amazonaws.com.
Name:   wp-user05-slave.cenae7eyijpr.ap-northeast-1.rds.amazonaws.com
Address: 10.0.3.217
```

## 系切り替え
**冗長構成となったAuroraのクラスタ〜エンドポイントを切り替えてみましょう。**

![fail-over-1](./images/step-3/fail-over-1.png "FAIL-OVER1")

----
![fail-over-2](./images/step-3/fail-over-2.png "FAIL-OVER2")

----
![fail-over-3](./images/step-3/fail-over-3.png "FAIL-OVER3")

----
![fail-over-4](./images/step-3/fail-over-4.png "FAIL-OVER4")

----

**ネットワークセグメントの確認(クラスタエンドポイント)をしましょう**

```
$ nslookup wp-user05-cluster.cluster-cenae7eyijpr.ap-northeast-1.rds.amazonaws.com
Server:     10.0.0.2
Address:    10.0.0.2#53

Non-authoritative answer:
wp-user05-cluster.cluster-cenae7eyijpr.ap-northeast-1.rds.amazonaws.com canonical name = wp-user05-slave.cenae7eyijpr.ap-northeast-1.rds.amazonaws.com.
Name:   wp-user05-slave.cenae7eyijpr.ap-northeast-1.rds.amazonaws.com
Address: 10.0.3.217
```

**続いてネットワークセグメントの確認(読み込みエンドポイント)をしましょう**

```
$ nslookup wp-user05-cluster.cluster-ro-cenae7eyijpr.ap-northeast-1.rds.amazonaws.com
Server:     10.0.0.2
Address:    10.0.0.2#53

Non-authoritative answer:
wp-user05-cluster.cluster-ro-cenae7eyijpr.ap-northeast-1.rds.amazonaws.com  canonical name = wp-user05.cenae7eyijpr.ap-northeast-1.rds.amazonaws.com.
Name:   wp-user05.cenae7eyijpr.ap-northeast-1.rds.amazonaws.com
Address: 10.0.2.226
```

## WebサーバのAMIを作成
**ここではWebサーバの水平分散にて必要となるEC2サーバのAMIを作成しましょう。AMIを作成することでここまでEC2サーバに設定した内容など全て反映された状態(Auroraの接続情報等)で水平分散が可能になります。**

![create-ami-1](./images/step-3/create-ami-1.png "CREATE-AMI1")

----
![create-ami-2](./images/step-3/create-ami-2.png "CREATE-AMI2")

----
![create-ami-3](./images/step-3/create-ami-3.png "CREATE-AMI3")

----
![create-ami-4](./images/step-3/create-ami-4.png "CREATE-AMI4")

----
**数分で作成が完了します**

![create-ami-5](./images/step-3/create-ami-5.png "CREATE-AMI5")

----

## 2台目のEC2インスタンスの作成(水平分散準備)

![create-ec2-1](./images/step-3/create-ec2-1.png "CREATE-EC2-1")

----
![create-ec2-2](./images/step-3/create-ec2-2.png "CREATE-EC2-2")

----
![create-ec2-3](./images/step-3/create-ec2-3.png "CREATE-EC2-3")

----
![create-ec2-4](./images/step-3/create-ec2-4.png "CREATE-EC2-4")

----
![create-ec2-5](./images/step-3/create-ec2-5.png "CREATE-EC2-5")

----
![create-ec2-6](./images/step-3/create-ec2-6.png "CREATE-EC2-6")

----
![create-ec2-7](./images/step-3/create-ec2-7.png "CREATE-EC2-7")

----
![create-ec2-8](./images/step-3/create-ec2-8.png "CREATE-EC2-8")

----
![create-ec2-9](./images/step-3/create-ec2-9.png "CREATE-EC2-9")

----
![create-ec2-10](./images/step-3/create-ec2-10.png "CREATE-EC2-10")

----
![create-ec2-11](./images/step-3/create-ec2-11.png "CREATE-EC2-11")

----
**作成完了まで数分掛かります**

![create-ec2-12](./images/step-3/create-ec2-12.png "CREATE-EC2-12")

----
**パブリックDNS(IPv4)のあたいを確認しブラウザでWordpressが参照できるか確認しましょう**

![create-ec2-13](./images/step-3/create-ec2-13.png "CREATE-EC2-13")

----
**Wordpressが表示されれば成功です**

![create-ec2-14](./images/step-3/create-ec2-14.png "CREATE-EC2-14")

----

## ELB(ALB)の作成
**ここでは作成した2台のEC2インスタンスの前段にアクセスを振り分けるELB(ALB)の作成を行います**

![create-elb-1](./images/step-3/create-elb-1.png "CREATE-ELB-1")

----

![create-elb-2](./images/step-3/create-elb-2.png "CREATE-ELB-2")

----

![create-elb-3-1](./images/step-3/create-elb-3-1.png "CREATE-ELB-3-1")

----
![create-elb-3-2](./images/step-3/create-elb-3-2.png "CREATE-ELB-3-2")

----
![create-elb-3-3](./images/step-3/create-elb-3-3.png "CREATE-ELB-3-3")

----
![create-elb-4](./images/step-3/create-elb-4.png "CREATE-ELB-4")

----
![create-elb-5](./images/step-3/create-elb-5.png "CREATE-ELB-5")

----
![create-elb-6](./images/step-3/create-elb-6.png "CREATE-ELB-6")

----
![create-elb-7-1](./images/step-3/create-elb-7-1.png "CREATE-ELB-7-1")

----
![create-elb-7-2](./images/step-3/create-elb-7-2.png "CREATE-ELB-7-2")

----
![create-elb-8](./images/step-3/create-elb-8.png "CREATE-ELB-8")

----
![create-elb-9](./images/step-3/create-elb-9.png "CREATE-ELB-9")

----
![create-elb-10](./images/step-3/create-elb-10.png "CREATE-ELB-10")

----
![create-elb-11](./images/step-3/create-elb-11.png "CREATE-ELB-11")

----
