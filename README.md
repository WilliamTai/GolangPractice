# Golang Practice

This Repo. is for learing Golang, do some coding and leave some memo what I think is important.

Also try to write multi-language Doc. to improve my language skill. (I hope so.)

## Something important in Golang

1. In Golang, all about type.
2. All call by value, no others.
3. No class but struct.
4. A method is be attached on struct.
5. Simple, easy programing

## Trouble shooting

### Debugserver error in MacOS

~~~
could not launch process: debugserver or lldb-server not found: install XCode's command line tools or lldb-server
~~~

Install xcode command-line tool to fix it.

~~~
$xcode-select --install
~~~

### Error Handling

Always, always, always check the error.

be aware the difference of Panic and fatal.

Fatal means shutdown the program immediately with os.exit(1).

Panic you can do something with defer function to recover.

Ref.https://blog.golang.org/defer-panic-and-recover

Error is an interface. if we create another type which implement the error interface, we can add some field or struct to store something we want. when error throw, we can catch the information we want.

### Naming methods

Export or not Export.
Visable or not visalbe.

We don't say public or private in Golang.

A method that named beging with lower case means not export(visalbe) outside the package.

## Add some comment above the method begin with //

The comments will become description of the method, it generated automatically.

# Golang練習用リポジトリ

Go言語練習用のリポジトリ、コードを書いたりしながら自分的に大事な（忘れそうな）メモを残しておく。

（人間の）言語能力のトレーニングため、多言語でメモを書いていきます。

# Go程式語言練習

Go程式語言練習用的Repo.，寫code和寫一點筆記。（這個人記憶力不好大概很快就忘了。）

為了練習（人類用的）語言能力，文件會用多語言撰寫。