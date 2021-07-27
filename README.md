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

## タグについて
```-n```はプレイヤー名を指定できます。プレイヤー名は、ハイスコアを更新したときに登録されます。デフォルトはguestになっています。  
```-r```はレートを指定できます。レートは掛け金のことです。１から５の間で選択でき、レートが上がるにつれて獲得スコアも高くなります。
## Demo
```go run main.go -n kimuson -r 5```で実行していきます。
すると、
```
HIGH SCORE | Name: , Score:
Welcome to go-poker, kimuson.
rating is 5
if you are ready, input 'y'
```
と表示されます。ここで'y'以外を入力すると、
```
this command is not allowed, please type again.
```
とでます。支持通り'y'を入力すると、
```
Let's start game!
Your Cards:
hands 1: Spade 10
hands 2: Heart 2
hands 3: Club 3
hands 4: Spade 9
hands 5: Diamond 2
Please input numbers you want to change in ascending order.
if you don't change cards, please input 'No'
if you want to change all, please input 'all'
For example: '1 2 3'
```
と表示されます。今回はHeart 2とDiamond 2でワンペアになっているので、それ以外を交換します。なので、
```1 3 4```を入力します。すると、
```
change hands 1, hands 3 and hands 4
hands 1: Club 9
hands 2: Heart 2
hands 3: Heart 1
hands 4: Heart 9
hands 5: Diamond 2
TWO PAIR!!
Your result: 10
Your current chip: 25
If you want to continue, type 'y'. If you want to exit, type 'n'.
```
と表示されます。ツーペアになったので、チップが25枚になりました。このまま続けることもできますが、今回はここで```n```を入力して終了に移ります。  
すると、
```
OK, we'll be loking forward to meeting you again someday.
Add HIGH SCORE: Name kimuson, Score 25
```
と表示されます。ハイスコアが登録されたことが分かります。
