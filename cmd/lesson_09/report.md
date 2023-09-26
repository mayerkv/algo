### Тайминги по сортировке массива из случайных чисел

```shell
BucketSort(100): 22.583µs

BucketSort(1000): 95.792µs

BucketSort(10000): 657.125µs

BucketSort(100000): 6.517458ms

BucketSort(1000000): 80.555417ms

BucketSort(10000000): 1.307292833s
```

### Сравнительная таблица времени выполнения

| Name            |      100 |      1_000 |      10_000 |       100_000 |    1_000_000 |   10_000_000 |
|-----------------|---------:|-----------:|------------:|--------------:|-------------:|-------------:|
| Bubble          | 17.333µs | 1.825166ms | 81.379916ms | 10.869929458s |      timeout |              | 
| Insertion       |  8.167µs | 1.160583ms | 36.503167ms |  3.588174458s |      timeout |              | 
| InsertionShift  | 20.792µs |  513.208µs | 12.762583ms |  1.230636792s |      timeout |              |
| InsertionBinary |  8.166µs |  367.792µs |  12.13625ms |     1.194615s |      timeout |              |
| Shell           | 13.042µs |  153.791µs |   776.375µs |   11.221417ms | 199.692333ms |              |
| Selection       | 16.834µs |  1.09975ms | 61.011834ms |  3.413933083s |      timeout |              |
| Heap            | 12.459µs |  168.167µs |   868.541µs |    9.102375ms |  125.12575ms |              |
| Quick           |  5.666µs |   39.875µs |   491.458µs |    5.502666ms |  65.843083ms | 739.014958ms |
| Merge           |     32µs |   74.916µs |   885.542µs |   10.841417ms | 108.186542ms |    1.229479s |
| Bucket          | 22.583µs |   95.792µs |   657.125µs |    6.517458ms |  80.555417ms | 1.307292833s |

### Бенчмарки для разных характеров данных

```shell
BenchmarkSort/Random/100/Bubble-8                        252483              4691 ns/op
BenchmarkSort/Random/100/Insertion-8                     266713              4558 ns/op
BenchmarkSort/Random/100/InsertionShift-8                474283              2477 ns/op
BenchmarkSort/Random/100/InsertionBinary-8                590215              1902 ns/op
BenchmarkSort/Random/100/Shell-8                         1326380               905.6 ns/op
BenchmarkSort/Random/100/SelectionSort-8                  245227              4689 ns/op
BenchmarkSort/Random/100/HeapSort-8                       613874              1910 ns/op
BenchmarkSort/Random/100/MergeSort-8                      257365              4357 ns/op
BenchmarkSort/Random/100/QuickSort-8                      945619              1268 ns/op

BenchmarkSort/Random/1000/Bubble-8                          2265            529008 ns/op
BenchmarkSort/Random/1000/Insertion-8                       2955            406963 ns/op
BenchmarkSort/Random/1000/InsertionShift-8                  6400            186033 ns/op
BenchmarkSort/Random/1000/InsertionBinary-8                 8407            142385 ns/op
BenchmarkSort/Random/1000/Shell-8                          30535             33936 ns/op
BenchmarkSort/Random/1000/SelectionSort-8                   2912            421052 ns/op
BenchmarkSort/Random/1000/HeapSort-8                       27237             42208 ns/op
BenchmarkSort/Random/1000/MergeSort-8                      20736             57999 ns/op
BenchmarkSort/Random/1000/QuickSort-8                      59995             17649 ns/op

BenchmarkSort/Random/10000/Bubble-8                           20          57941885 ns/op
BenchmarkSort/Random/10000/Insertion-8                        28          39882150 ns/op
BenchmarkSort/Random/10000/InsertionShift-8                   82          14824700 ns/op
BenchmarkSort/Random/10000/InsertionBinary-8                  86          14013541 ns/op
BenchmarkSort/Random/10000/Shell-8                          1328            889073 ns/op
BenchmarkSort/Random/10000/SelectionSort-8                    28          40553494 ns/op
BenchmarkSort/Random/10000/HeapSort-8                       1467            815755 ns/op
BenchmarkSort/Random/10000/MergeSort-8                      1215            974791 ns/op
BenchmarkSort/Random/10000/QuickSort-8                      2300            512885 ns/op

BenchmarkSort/Random/100000/Bubble-8                           1        12722349125 ns/op
BenchmarkSort/Random/100000/Insertion-8                        1        4161260500 ns/op
BenchmarkSort/Random/100000/InsertionShift-8                   1        1489903667 ns/op
BenchmarkSort/Random/100000/InsertionBinary-8                  1        1456918083 ns/op
BenchmarkSort/Random/100000/Shell-8                           87          13387904 ns/op
BenchmarkSort/Random/100000/SelectionSort-8                    1        4069291042 ns/op
BenchmarkSort/Random/100000/HeapSort-8                       100          11140387 ns/op
BenchmarkSort/Random/100000/MergeSort-8                       88          12654548 ns/op
BenchmarkSort/Random/100000/QuickSort-8                      184           6405863 ns/op

BenchmarkSort/Digits/100/Bubble-8                         262225              4454 ns/op
BenchmarkSort/Digits/100/Insertion-8                      277437              4247 ns/op
BenchmarkSort/Digits/100/InsertionShift-8                 482937              2469 ns/op
BenchmarkSort/Digits/100/InsertionBinary-8                635806              1828 ns/op
BenchmarkSort/Digits/100/Shell-8                         1478373               805.9 ns/op
BenchmarkSort/Digits/100/SelectionSort-8                  227769              5262 ns/op
BenchmarkSort/Digits/100/HeapSort-8                       616951              1956 ns/op
BenchmarkSort/Digits/100/MergeSort-8                      260646              4462 ns/op
BenchmarkSort/Digits/100/QuickSort-8                      756244              1595 ns/op

BenchmarkSort/Digits/1000/Bubble-8                          2311            511136 ns/op
BenchmarkSort/Digits/1000/Insertion-8                       3030            386879 ns/op
BenchmarkSort/Digits/1000/InsertionShift-8                  6859            170387 ns/op
BenchmarkSort/Digits/1000/InsertionBinary-8                 8640            137545 ns/op
BenchmarkSort/Digits/1000/Shell-8                          90722             11578 ns/op
BenchmarkSort/Digits/1000/SelectionSort-8                   2828            423275 ns/op
BenchmarkSort/Digits/1000/HeapSort-8                       33176             35100 ns/op
BenchmarkSort/Digits/1000/MergeSort-8                      17094             60325 ns/op
BenchmarkSort/Digits/1000/QuickSort-8                      22470             53311 ns/op

BenchmarkSort/Digits/10000/Bubble-8                           24          47288795 ns/op
BenchmarkSort/Digits/10000/Insertion-8                        37          32131295 ns/op
BenchmarkSort/Digits/10000/InsertionShift-8                  100          11412267 ns/op
BenchmarkSort/Digits/10000/InsertionBinary-8                 100          11014759 ns/op
BenchmarkSort/Digits/10000/Shell-8                          3816            310415 ns/op
BenchmarkSort/Digits/10000/SelectionSort-8                    34          34097183 ns/op
BenchmarkSort/Digits/10000/HeapSort-8                       2164            548282 ns/op
BenchmarkSort/Digits/10000/MergeSort-8                      1956            605283 ns/op
BenchmarkSort/Digits/10000/QuickSort-8                       296           4036527 ns/op

BenchmarkSort/Digits/100000/Bubble-8                           1        10705869792 ns/op
BenchmarkSort/Digits/100000/Insertion-8                        1        3289263167 ns/op
BenchmarkSort/Digits/100000/InsertionShift-8                   1        1119734333 ns/op
BenchmarkSort/Digits/100000/InsertionBinary-8                  1        1088970250 ns/op
BenchmarkSort/Digits/100000/Shell-8                          330           3682813 ns/op
BenchmarkSort/Digits/100000/SelectionSort-8                    1        3396718708 ns/op
BenchmarkSort/Digits/100000/HeapSort-8                       195           6093532 ns/op
BenchmarkSort/Digits/100000/MergeSort-8                      163           7310030 ns/op
BenchmarkSort/Digits/100000/QuickSort-8                        3         382188917 ns/op

BenchmarkSort/Sorted/100/Bubble-8                         418368              2811 ns/op
BenchmarkSort/Sorted/100/Insertion-8                     7587067               157.8 ns/op
BenchmarkSort/Sorted/100/InsertionShift-8                1406169               854.0 ns/op
BenchmarkSort/Sorted/100/InsertionBinary-8               6514872               181.9 ns/op
BenchmarkSort/Sorted/100/Shell-8                         2330835               515.2 ns/op
BenchmarkSort/Sorted/100/SelectionSort-8                  330610              3621 ns/op
BenchmarkSort/Sorted/100/HeapSort-8                       675300              1775 ns/op
BenchmarkSort/Sorted/100/MergeSort-8                      370312              3215 ns/op
BenchmarkSort/Sorted/100/QuickSort-8                      235424              5070 ns/op

BenchmarkSort/Sorted/1000/Bubble-8                          4878            245846 ns/op
BenchmarkSort/Sorted/1000/Insertion-8                     909704              1315 ns/op
BenchmarkSort/Sorted/1000/InsertionShift-8                104353             11478 ns/op
BenchmarkSort/Sorted/1000/InsertionBinary-8               714876              1639 ns/op
BenchmarkSort/Sorted/1000/Shell-8                         169746              7046 ns/op
BenchmarkSort/Sorted/1000/SelectionSort-8                   3690            326095 ns/op
BenchmarkSort/Sorted/1000/HeapSort-8                       31209             37787 ns/op
BenchmarkSort/Sorted/1000/MergeSort-8                      32755             36426 ns/op
BenchmarkSort/Sorted/1000/QuickSort-8                       3028            395468 ns/op

BenchmarkSort/Sorted/10000/Bubble-8                           49          24025255 ns/op
BenchmarkSort/Sorted/10000/Insertion-8                    102206             11605 ns/op
BenchmarkSort/Sorted/10000/InsertionShift-8                 7930            150694 ns/op
BenchmarkSort/Sorted/10000/InsertionBinary-8               81194             14686 ns/op
BenchmarkSort/Sorted/10000/Shell-8                         10000            100403 ns/op
BenchmarkSort/Sorted/10000/SelectionSort-8                    37          32014023 ns/op
BenchmarkSort/Sorted/10000/HeapSort-8                       2338            513033 ns/op
BenchmarkSort/Sorted/10000/MergeSort-8                      2704            439892 ns/op
BenchmarkSort/Sorted/10000/QuickSort-8                        31          38113634 ns/op

BenchmarkSort/Sorted/100000/Bubble-8                           1        2395711292 ns/op
BenchmarkSort/Sorted/100000/Insertion-8                     9848            117271 ns/op
BenchmarkSort/Sorted/100000/InsertionShift-8                 609           1962918 ns/op
BenchmarkSort/Sorted/100000/InsertionBinary-8               7869            149005 ns/op
BenchmarkSort/Sorted/100000/Shell-8                          955           1253591 ns/op
BenchmarkSort/Sorted/100000/SelectionSort-8                    1        3201100625 ns/op
BenchmarkSort/Sorted/100000/HeapSort-8                       196           6082881 ns/op
BenchmarkSort/Sorted/100000/MergeSort-8                      198           6213598 ns/op
BenchmarkSort/Sorted/100000/QuickSort-8                        1        3834014250 ns/op

BenchmarkSort/Reversed/100/Bubble-8                       163998              7312 ns/op
BenchmarkSort/Reversed/100/Insertion-8                    167548              7151 ns/op
BenchmarkSort/Reversed/100/InsertionShift-8               319842              3715 ns/op
BenchmarkSort/Reversed/100/InsertionBinary-8              412230              2915 ns/op
BenchmarkSort/Reversed/100/Shell-8                       1806769               658.6 ns/op
BenchmarkSort/Reversed/100/SelectionSort-8                307263              3891 ns/op
BenchmarkSort/Reversed/100/HeapSort-8                     800775              1484 ns/op
BenchmarkSort/Reversed/100/MergeSort-8                    373761              3159 ns/op
BenchmarkSort/Reversed/100/QuickSort-8                    237444              5057 ns/op

BenchmarkSort/Reversed/1000/Bubble-8                        1646            728575 ns/op
BenchmarkSort/Reversed/1000/Insertion-8                     1689            710370 ns/op
BenchmarkSort/Reversed/1000/InsertionShift-8                4621            260180 ns/op
BenchmarkSort/Reversed/1000/InsertionBinary-8               4861            245997 ns/op
BenchmarkSort/Reversed/1000/Shell-8                       113697             10483 ns/op
BenchmarkSort/Reversed/1000/SelectionSort-8                 3528            340684 ns/op
BenchmarkSort/Reversed/1000/HeapSort-8                     36982             31076 ns/op
BenchmarkSort/Reversed/1000/MergeSort-8                    33108             36214 ns/op
BenchmarkSort/Reversed/1000/QuickSort-8                     3238            368959 ns/op

BenchmarkSort/Reversed/10000/Bubble-8                         15          72180936 ns/op
BenchmarkSort/Reversed/10000/Insertion-8                      16          70039698 ns/op
BenchmarkSort/Reversed/10000/InsertionShift-8                 49          24235581 ns/op
BenchmarkSort/Reversed/10000/InsertionBinary-8                49          24095300 ns/op
BenchmarkSort/Reversed/10000/Shell-8                        8307            145005 ns/op
BenchmarkSort/Reversed/10000/SelectionSort-8                  36          33072215 ns/op
BenchmarkSort/Reversed/10000/HeapSort-8                     2282            518206 ns/op
BenchmarkSort/Reversed/10000/MergeSort-8                    2727            433358 ns/op
BenchmarkSort/Reversed/10000/QuickSort-8                      33          35173582 ns/op

BenchmarkSort/Reversed/100000/Bubble-8                         1        7404523583 ns/op
BenchmarkSort/Reversed/100000/Insertion-8                      1        7269989542 ns/op
BenchmarkSort/Reversed/100000/InsertionShift-8                 1        2413946750 ns/op
BenchmarkSort/Reversed/100000/InsertionBinary-8                1        2410971000 ns/op
BenchmarkSort/Reversed/100000/Shell-8                        615           1944191 ns/op
BenchmarkSort/Reversed/100000/SelectionSort-8                  1        3305952375 ns/op
BenchmarkSort/Reversed/100000/HeapSort-8                     195           6103145 ns/op
BenchmarkSort/Reversed/100000/MergeSort-8                    222           5407128 ns/op
BenchmarkSort/Reversed/100000/QuickSort-8                      1        3513162167 ns/op
```
