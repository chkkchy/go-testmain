# go-testmain

exec test:

```
~/.g/s/g/c/go-testmain ❯❯❯ go test -v ./...
```

output:

```
?   	github.com/chkkchy/go-testmain	[no test files]
2016/11/14 19:29:22 This gets run BEFORE any tests get run!
2016/11/14 19:29:22 This gets run AFTER any tests get run!
ok  	github.com/chkkchy/go-testmain/fuga	0.006s
=== RUN   TestOne
do hogefunc1
2016/11/14 19:29:22 TestOne running
--- PASS: TestOne (0.00s)
=== RUN   TestTwo
do hogefunc2
2016/11/14 19:29:22 TestTwo running
--- PASS: TestTwo (0.00s)
PASS
ok  	github.com/chkkchy/go-testmain/hoge	0.007s
```
