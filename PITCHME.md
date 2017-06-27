# twitter cli 作成

### 環境

- golang が動く環境

## 参考資料

- [urfave-cli](https://godoc.org/github.com/urfave/cli)
- [anaconda](https://godoc.org/github.com/ChimeraCoder/anaconda)
- [qiita](http://qiita.com/stenpel/items/01f6ef20c4691564faa5) 
- [go-pipeline](https://github.com/mattn/go-pipeline)

---

## はじめに
#### 作業するディレクトリへ
``` sh
$ cd $GOPATH/src
$ mkdir twicli
$ cd twicli
```

#### 必要なライブラリを go get !!! & install
``` sh
$ brew install fzf

$ go get github.com/ChimeraCoder/anaconda
$ go get github.com/urfave/cli
$ go get github.com/gizak/termui
```
---
## 早速ツイートしてみよう！

#### main.go
---?code=main1.go

---
#### tweet.go
---?code=tweet.go
<br>
これでtweetするための準備ができました。
早速、`go build` してtweetしてみましょう。<br>
うまく、tweetができれば成功です！次に進みましょう。

---

## User情報を出力！

次に`user.go`を作成し、user情報を取得し`termui`を使用し出力してみましょう。

---

#### main.go　に以下の文を追加
```sh
{
	Name:    "user",
	Aliases: []string{"u"},
	Usage:   "Show user's profile",
	Action:  user,
},
```
<br>
---

#### user.go
---?code=user.go

試しに `$ twicli user TigersDreamlink` を行うと<br><br>
<!--<img src=user.png width=400px>-->
---?image=./images/user.png&size=auto 90%
<br>
---
<br>
とこのようにtwitterのuser情報を取得し、出力することができます。

---

## TwitterのTimeLineを取得
次は**twitter**の**タイムライン**を取得しよう〜<br>
**twitter**の**タイムライン**はこの部分です
<br><br>

<!--<img src=timeline.png width=100px>-->
---?image=./images/timeline.png&size=auto 70% 

---
次はこれを実装していきます!

---

#### main.go
以下を`main.go`に追加します

```sh
{
	Name:    "timeline",
	Aliases: []string{"ti", "time"},
	Usage:   "Show twitter timelines of 15",
	Action:  timeline,
},
```
<br>
---
#### timeline.go
---?code=timeline.go
---

<br>
実際に `$ twicli timeline` を行うと<br><br>
---?image=./images/timeline2.png&size=auto 90%

<!--<img src=timeline2.png width=400px>-->
---
<br>
<br>
とこのようにtwitterのタイムラインを取得し、出力することができます。
<br>
---

## 検索機能をつけよう！
次はtwitterの検索機能をつけます。
検索機能は以下のようです
<br><br>
<!--<img src=search.png width=100px>-->
---?image=./images/search.png&size=auto 70% 

---
次はこれを実装していきます! 
---

#### main.go
以下を`main.go`に追加します

```sh
{
	Name:    "search",
	Aliases: []string{"s"},
	Usage:   "Search latest 15 tweets",
	Action:  search,
},
```
<br>
---

#### search.go
---?code=search.go
---
<br>
実際に `$ twicli search キングダム` を行うと<br><br>
---?image=./images/search2.png&size=auto 90%

---
<!--<img src=search2.png width=400px>-->
<br>
<br>
とこのようにtwitterの**検索機能**を使うことができます

<br>
---

## 画像付きツイートをしよう！
ここでは画像付きツイートの実装をしていきます。<br>
また、最初にした`fzf`の出番がようやく来ました！早速実装しましょう

<br>
---
以下をいつも通り`main.go`に追加します

```sh
{
	Name:    "pictweet",
	Aliases: []string{"pic", "p"},
	Usage:   "Tweet picture and sentence",
	Action:  pictweet,
},
```
<br>
---

#### pictweet.go

---?code=pictweet1.go

---
#####まず、これで画像付きツイートができるようになります、早速、適当な画像を保存してツイートしてみましょう！

<br><br>
####成功したら次に進みましょう！
---

##次にファジーサーチ機能をつけて画像を取得しツイートしよう👍
まず、画像を保存しておくためのディレクトリを作成し、そこに画像を保存し、そこから取得しツイートできるようにします。

<br>
---

####main.go
`main.go`を以下のようにします

---?code=main.go

---

####pictweet.go
`pictweet.go`も同様に変更していきます<br>

---?code=pictweet.go

---

##次にファジーサーチ用のメソッドを定義します
ファジーサーチメソッドを定義する上で必要なライブラリを`go get`します!<br>

```sh
$ go get github.com/mattn/go-pipeline
```
---

####getimage.go

---?code=getimage.go

---

###さぁ、これでファジーサーチを利用して画像付きツイートができるようになりました！<br>
<br>

---

## 最後に
このように**golang**では簡単に**CLI**を作成でき、実装できます。また、CLIを作る上でのライブラリも豊富なので、是非皆さんも自分なりの、自分にあった、もしくは便利なCLIを作ってみてください！
<br>
golang... **楽しい**<br>
もう... 最高です!!!



<!-- $ export CON_KEY_TW="Your-Twitter-Consumer-Key"
$ export CON_SECRET_KEY_TW="Your-Twitter-Consumer-Secret-Key"
$ export ACC_KEY_TW="Your-Twitter-Access-Token"
$ export ACC_SECRET_KEY_TW="Your-Twitter-Access-Token-Secret" -->
