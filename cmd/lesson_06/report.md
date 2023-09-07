Тайминги по сортировке массива из случайных чисел
```shell
Bubble(100): 17.333µs
Insertion(100): 8.167µs
InsertionShift(100): 20.792µs
InsertionBinary(100): 8.166µs
Shell(100): 13.042µs

Bubble(1000): 1.825166ms
Insertion(1000): 1.160583ms
InsertionShift(1000): 513.208µs
InsertionBinary(1000): 367.792µs
Shell(1000): 153.791µs

Bubble(10000): 81.379916ms
Insertion(10000): 36.503167ms
InsertionShift(10000): 12.762583ms
InsertionBinary(10000): 12.13625ms
Shell(10000): 776.375µs

Bubble(100000): 10.869929458s
Insertion(100000): 3.588174458s
InsertionShift(100000): 1.230636792s
InsertionBinary(100000): 1.194615s
Shell(100000): 11.221417ms

Bubble(1000000): timeout
Insertion(1000000): timeout
InsertionShift(1000000): timeout
InsertionBinary(1000000): timeout
Shell(1000000): 199.692333ms

```

Бенчмарки для разных характеров данных
```shell
go test ./pkg/sort/... -bench ^BenchmarkSort
goos: darwin
goarch: arm64
pkg: github.com/mayerkv/algo/pkg/sort

BenchmarkSort/Random/100/Bubble-8                         275258              4254 ns/op
BenchmarkSort/Random/100/Insertion-8                      284262              4179 ns/op
BenchmarkSort/Random/100/InsertionShift-8                 555378              2139 ns/op
BenchmarkSort/Random/100/InsertionBinary-8                750210              1551 ns/op
BenchmarkSort/Random/100/Shell-8                         1529072               787.0 ns/op

BenchmarkSort/Random/1000/Bubble-8                          2584            458981 ns/op
BenchmarkSort/Random/1000/Insertion-8                       3367            352233 ns/op
BenchmarkSort/Random/1000/InsertionShift-8                  7515            159746 ns/op
BenchmarkSort/Random/1000/InsertionBinary-8                 9885            122582 ns/op
BenchmarkSort/Random/1000/Shell-8                          34579             32782 ns/op

BenchmarkSort/Random/10000/Bubble-8                           24          48611892 ns/op
BenchmarkSort/Random/10000/Insertion-8                        33          34621989 ns/op
BenchmarkSort/Random/10000/InsertionShift-8                   94          12546672 ns/op
BenchmarkSort/Random/10000/InsertionBinary-8                  99          11945530 ns/op
BenchmarkSort/Random/10000/Shell-8                          1608            744443 ns/op

BenchmarkSort/Random/100000/Bubble-8                           1        10884383084 ns/op
BenchmarkSort/Random/100000/Insertion-8                        1        3579053458 ns/op
BenchmarkSort/Random/100000/InsertionShift-8                   1        1230848291 ns/op
BenchmarkSort/Random/100000/InsertionBinary-8                  1        1205538084 ns/op
BenchmarkSort/Random/100000/Shell-8                          100          11204097 ns/op

BenchmarkSort/Sorted/100/Bubble-8                         424886              2801 ns/op
BenchmarkSort/Sorted/100/Insertion-8                     7895181               152.3 ns/op
BenchmarkSort/Sorted/100/InsertionShift-8                1408500               852.3 ns/op
BenchmarkSort/Sorted/100/InsertionBinary-8               6852363               175.7 ns/op
BenchmarkSort/Sorted/100/Shell-8                         2364781               509.1 ns/op

BenchmarkSort/Sorted/1000/Bubble-8                          4933            244224 ns/op
BenchmarkSort/Sorted/1000/Insertion-8                     925794              1268 ns/op
BenchmarkSort/Sorted/1000/InsertionShift-8                104740             11441 ns/op
BenchmarkSort/Sorted/1000/InsertionBinary-8               741886              1578 ns/op
BenchmarkSort/Sorted/1000/Shell-8                         171148              6982 ns/op

BenchmarkSort/Sorted/10000/Bubble-8                           49          23929531 ns/op
BenchmarkSort/Sorted/10000/Insertion-8                    108520             11034 ns/op
BenchmarkSort/Sorted/10000/InsertionShift-8                 7976            149942 ns/op
BenchmarkSort/Sorted/10000/InsertionBinary-8               84601             14179 ns/op
BenchmarkSort/Sorted/10000/Shell-8                         12067             99326 ns/op

BenchmarkSort/Sorted/100000/Bubble-8                           1        2390007750 ns/op
BenchmarkSort/Sorted/100000/Insertion-8                    10000            108747 ns/op
BenchmarkSort/Sorted/100000/InsertionShift-8                 614           1952114 ns/op
BenchmarkSort/Sorted/100000/InsertionBinary-8               8262            139022 ns/op
BenchmarkSort/Sorted/100000/Shell-8                          960           1240714 ns/op

BenchmarkSort/Reversed/100/Bubble-8                       164203              7241 ns/op
BenchmarkSort/Reversed/100/Insertion-8                    170524              7045 ns/op
BenchmarkSort/Reversed/100/InsertionShift-8               325664              3673 ns/op
BenchmarkSort/Reversed/100/InsertionBinary-8              424934              2802 ns/op
BenchmarkSort/Reversed/100/Shell-8                       1824944               657.3 ns/op

BenchmarkSort/Reversed/1000/Bubble-8                        1647            722796 ns/op
BenchmarkSort/Reversed/1000/Insertion-8                     1707            706469 ns/op
BenchmarkSort/Reversed/1000/InsertionShift-8                4658            258997 ns/op
BenchmarkSort/Reversed/1000/InsertionBinary-8               4873            246012 ns/op
BenchmarkSort/Reversed/1000/Shell-8                       114548             10436 ns/op

BenchmarkSort/Reversed/10000/Bubble-8                         15          71668286 ns/op
BenchmarkSort/Reversed/10000/Insertion-8                      16          69347641 ns/op
BenchmarkSort/Reversed/10000/InsertionShift-8                 49          24180876 ns/op
BenchmarkSort/Reversed/10000/InsertionBinary-8                49          23984111 ns/op
BenchmarkSort/Reversed/10000/Shell-8                        8344            143978 ns/op

BenchmarkSort/Reversed/100000/Bubble-8                         1        7349833916 ns/op
BenchmarkSort/Reversed/100000/Insertion-8                      1        7214965625 ns/op
BenchmarkSort/Reversed/100000/InsertionShift-8                 1        2402344250 ns/op
BenchmarkSort/Reversed/100000/InsertionBinary-8                1        2402072875 ns/op
BenchmarkSort/Reversed/100000/Shell-8                        619           1935286 ns/op

PASS
ok      github.com/mayerkv/algo/pkg/sort        106.713s

```