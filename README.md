# atcoderの提出ソースコードを貼っていくだけ

## generate.go
```
go run generate.go "contestName" "contestNumber" "generateCount"
```

## [WIP] anyenv, goenv
atcoderは1.6環境のためそれに合わせる
- winだとうまくいかないので何とかする
```
goenv install 1.6.4
goenv local 1.6.4
goenv rehash
go version
```
#### ref
- https://github.com/riywo/anyenv/blob/master/README.md
- https://qiita.com/zaburo/items/8ac16133c3823c6e6ad6

## ref
### bitCount
official function: OnesCount64
