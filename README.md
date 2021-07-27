# go-poker -コマンドライン上で動くマシンポーカー-
[![Build Status](https://travis-ci.com/kimuson13/go-poker.svg?branch=main)](https://travis-ci.com/kimuson13/go-poker)  
## 概要
Go製のマシンポーカーです。  
基本的なルールは[ここ](https://www.xn--lck0a5auxk.jp/videopoker.html)を参照しています。  
暇つぶしにどうぞ！

## 使い方
```
$ git clone https://github.com/kimuson13/go-poker.git
$ cd cmd/poker
$ go build main.go [オプション]
```
もしくは、
```
$ git clone https://github.com/kimuson13/go-poker.git
$ cd cmd/poker
$ go run main.go [オプション]
```
また、
```
go get -u github.com/kimuson13/go-poker/cmd/poker
```
でも獲得可能です。(たぶん)

## タグについて
```-n```はプレイヤー名を指定できます。プレイヤー名は、ハイスコアを更新したときに登録されます。デフォルトはguestになっています。  
```-r```はレートを指定できます。レートは掛け金のことです。１から５の間で選択でき、レートが上がるにつれて獲得スコアも高くなります。
## Demo
```go run main.go -n kimuson -r 5```で実行していきます。
