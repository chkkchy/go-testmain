# go-testmain

```
go test ./...
```

> ?   	github.com/chkkchy/go-testmain	[no test files]
2016/11/14 18:36:42 This gets run BEFORE any tests get run!
do hogefunc1
2016/11/14 18:36:42 TestOne running
do hogefunc2
2016/11/14 18:36:42 TestTwo running
--- FAIL: TestTwo (0.00s)
FAIL
2016/11/14 18:36:42 This gets run AFTER any tests get run!
FAIL	github.com/chkkchy/go-testmain/hoge	0.007s
