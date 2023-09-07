### Тайминги по сортировке массива из случайных чисел
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

| Name            |       100 |      1_000 |      10_000 |       100_000 |    1_000_000 |
|-----------------|----------:|-----------:|------------:|--------------:|-------------:|
| Bubble          |  17.333µs | 1.825166ms | 81.379916ms | 10.869929458s |      timeout |
| Insertion       |   8.167µs | 1.160583ms | 36.503167ms |  3.588174458s |      timeout |
| InsertionShift  |  20.792µs |  513.208µs | 12.762583ms |  1.230636792s |      timeout |
| InsertionBinary |   8.166µs |  367.792µs |  12.13625ms |     1.194615s |      timeout |
| Shell           |  13.042µs |  153.791µs |   776.375µs |   11.221417ms | 199.692333ms |


### Бенчмарки для разных характеров данных

```shell
BenchmarkSort/Random/100/Bubble-8                         275719              4490 ns/op
BenchmarkSort/Random/100/Insertion-8                      286521              4187 ns/op
BenchmarkSort/Random/100/InsertionShift-8                 533038              2170 ns/op
BenchmarkSort/Random/100/InsertionBinary-8                736195              1603 ns/op
BenchmarkSort/Random/100/Shell-8                         1515880               796.6 ns/op

BenchmarkSort/Random/1000/Bubble-8                          2595            458005 ns/op
BenchmarkSort/Random/1000/Insertion-8                       3387            357298 ns/op
BenchmarkSort/Random/1000/InsertionShift-8                  7300            161650 ns/op
BenchmarkSort/Random/1000/InsertionBinary-8                 9602            123858 ns/op
BenchmarkSort/Random/1000/Shell-8                          33602             33231 ns/op

BenchmarkSort/Random/10000/Bubble-8                           24          48624217 ns/op
BenchmarkSort/Random/10000/Insertion-8                        33          34781721 ns/op
BenchmarkSort/Random/10000/InsertionShift-8                   94          12589104 ns/op
BenchmarkSort/Random/10000/InsertionBinary-8                 100          11891542 ns/op
BenchmarkSort/Random/10000/Shell-8                          1592            753348 ns/op

BenchmarkSort/Random/100000/Bubble-8                           1        10856777792 ns/op
BenchmarkSort/Random/100000/Insertion-8                        1        3633154916 ns/op
BenchmarkSort/Random/100000/InsertionShift-8                   1        1263613334 ns/op
BenchmarkSort/Random/100000/InsertionBinary-8                  1        1230495625 ns/op
BenchmarkSort/Random/100000/Shell-8                          100          11301761 ns/op

BenchmarkSort/Digits/100/Bubble-8                         318501              3829 ns/op
BenchmarkSort/Digits/100/Insertion-8                      322300              3664 ns/op
BenchmarkSort/Digits/100/InsertionShift-8                 593554              2054 ns/op
BenchmarkSort/Digits/100/InsertionBinary-8                802434              1458 ns/op
BenchmarkSort/Digits/100/Shell-8                         1762642               683.2 ns/op

BenchmarkSort/Digits/1000/Bubble-8                          2809            430695 ns/op
BenchmarkSort/Digits/1000/Insertion-8                       3646            331735 ns/op
BenchmarkSort/Digits/1000/InsertionShift-8                  7809            142574 ns/op
BenchmarkSort/Digits/1000/InsertionBinary-8                10000            114496 ns/op
BenchmarkSort/Digits/1000/Shell-8                         102592             10605 ns/op

BenchmarkSort/Digits/10000/Bubble-8                           24          47489413 ns/op
BenchmarkSort/Digits/10000/Insertion-8                        36          32238900 ns/op
BenchmarkSort/Digits/10000/InsertionShift-8                  100          11626107 ns/op
BenchmarkSort/Digits/10000/InsertionBinary-8                 100          11145163 ns/op
BenchmarkSort/Digits/10000/Shell-8                          3694            315427 ns/op

BenchmarkSort/Digits/100000/Bubble-8                           1        10912837500 ns/op
BenchmarkSort/Digits/100000/Insertion-8                        1        3325103334 ns/op
BenchmarkSort/Digits/100000/InsertionShift-8                   1        1139547041 ns/op
BenchmarkSort/Digits/100000/InsertionBinary-8                  1        1110277250 ns/op
BenchmarkSort/Digits/100000/Shell-8                          318           3705318 ns/op

BenchmarkSort/Sorted/100/Bubble-8                         417800              2868 ns/op
BenchmarkSort/Sorted/100/Insertion-8                     7532934               156.9 ns/op
BenchmarkSort/Sorted/100/InsertionShift-8                1390292               869.6 ns/op
BenchmarkSort/Sorted/100/InsertionBinary-8               6679910               180.0 ns/op
BenchmarkSort/Sorted/100/Shell-8                         2345419               518.6 ns/op

BenchmarkSort/Sorted/1000/Bubble-8                          4838            247530 ns/op
BenchmarkSort/Sorted/1000/Insertion-8                     949407              1308 ns/op
BenchmarkSort/Sorted/1000/InsertionShift-8                103824             11653 ns/op
BenchmarkSort/Sorted/1000/InsertionBinary-8               707769              1628 ns/op
BenchmarkSort/Sorted/1000/Shell-8                         169738              7275 ns/op

BenchmarkSort/Sorted/10000/Bubble-8                           49          24431753 ns/op
BenchmarkSort/Sorted/10000/Insertion-8                    103306             12157 ns/op
BenchmarkSort/Sorted/10000/InsertionShift-8                 7726            153471 ns/op
BenchmarkSort/Sorted/10000/InsertionBinary-8               84379             14869 ns/op
BenchmarkSort/Sorted/10000/Shell-8                         10000            102967 ns/op

BenchmarkSort/Sorted/100000/Bubble-8                           1        2423159000 ns/op
BenchmarkSort/Sorted/100000/Insertion-8                    10000            109155 ns/op
BenchmarkSort/Sorted/100000/InsertionShift-8                 601           1966986 ns/op
BenchmarkSort/Sorted/100000/InsertionBinary-8               8388            144532 ns/op
BenchmarkSort/Sorted/100000/Shell-8                          931           1270058 ns/op

BenchmarkSort/Reversed/100/Bubble-8                       164166              7408 ns/op
BenchmarkSort/Reversed/100/Insertion-8                    167715              7229 ns/op
BenchmarkSort/Reversed/100/InsertionShift-8               319107              3748 ns/op
BenchmarkSort/Reversed/100/InsertionBinary-8              413220              2918 ns/op
BenchmarkSort/Reversed/100/Shell-8                       1801124               658.0 ns/op

BenchmarkSort/Reversed/1000/Bubble-8                        1646            734505 ns/op
BenchmarkSort/Reversed/1000/Insertion-8                     1681            716539 ns/op
BenchmarkSort/Reversed/1000/InsertionShift-8                4593            262525 ns/op
BenchmarkSort/Reversed/1000/InsertionBinary-8               4695            250015 ns/op
BenchmarkSort/Reversed/1000/Shell-8                       110881             10642 ns/op

BenchmarkSort/Reversed/10000/Bubble-8                         15          73793303 ns/op
BenchmarkSort/Reversed/10000/Insertion-8                      16          70816526 ns/op
BenchmarkSort/Reversed/10000/InsertionShift-8                 49          24705418 ns/op
BenchmarkSort/Reversed/10000/InsertionBinary-8                48          24477017 ns/op
BenchmarkSort/Reversed/10000/Shell-8                        7987            147985 ns/op

BenchmarkSort/Reversed/100000/Bubble-8                         1        7519433333 ns/op
BenchmarkSort/Reversed/100000/Insertion-8                      1        7388163250 ns/op
BenchmarkSort/Reversed/100000/InsertionShift-8                 1        2457891417 ns/op
BenchmarkSort/Reversed/100000/InsertionBinary-8                1        2436670792 ns/op
BenchmarkSort/Reversed/100000/Shell-8                        619           1957478 ns/op
```