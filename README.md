# go-quine
main.go and the output of main.go is exactly equivalent.

```console
# the below command doesn't output anything
$ diff main.go <(go run main.go)
```
