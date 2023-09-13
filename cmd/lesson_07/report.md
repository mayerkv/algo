### Тайминги по сортировке массива из случайных чисел
```shell
SelectionSort(100): 16.834µs
HeapSort(100): 12.459µs

SelectionSort(1000): 1.09975ms
HeapSort(1000): 168.167µs

SelectionSort(10000): 61.011834ms
HeapSort(10000): 868.541µs

SelectionSort(100000): 3.413933083s
HeapSort(100000): 9.102375ms
```

| Name            |       100 |      1_000 |      10_000 |       100_000 |    1_000_000 |
|-----------------|----------:|-----------:|------------:|--------------:|-------------:|
| Bubble          |  17.333µs | 1.825166ms | 81.379916ms | 10.869929458s |      timeout |
| Insertion       |   8.167µs | 1.160583ms | 36.503167ms |  3.588174458s |      timeout |
| InsertionShift  |  20.792µs |  513.208µs | 12.762583ms |  1.230636792s |      timeout |
| InsertionBinary |   8.166µs |  367.792µs |  12.13625ms |     1.194615s |      timeout |
| Shell           |  13.042µs |  153.791µs |   776.375µs |   11.221417ms | 199.692333ms |
| Selection       |  16.834µs |  1.09975ms | 61.011834ms |  3.413933083s |      timeout |
| Heap            |  12.459µs |  168.167µs |   868.541µs |    9.102375ms |  125.12575ms |

### Бенчмарки для разных характеров данных
```shell
BenchmarkSort/Random/100/Bubble-8                 275692              4337 ns/op
BenchmarkSort/Random/100/Insertion-8              287878              4152 ns/op
BenchmarkSort/Random/100/InsertionShift-8         532851              2168 ns/op
BenchmarkSort/Random/100/InsertionBinary-8                690742              1703 ns/op
BenchmarkSort/Random/100/Shell-8                         1525213               787.6 ns/op
BenchmarkSort/Random/100/SelectionSort-8                  290318              4119 ns/op
BenchmarkSort/Random/100/HeapSort-8                       732445              1631 ns/op

BenchmarkSort/Random/1000/Bubble-8                          2595            463874 ns/op
BenchmarkSort/Random/1000/Insertion-8                       3392            352878 ns/op
BenchmarkSort/Random/1000/InsertionShift-8                  7496            159895 ns/op
BenchmarkSort/Random/1000/InsertionBinary-8                 9726            123248 ns/op
BenchmarkSort/Random/1000/Shell-8                          33972             32665 ns/op
BenchmarkSort/Random/1000/SelectionSort-8                   3312            364706 ns/op
BenchmarkSort/Random/1000/HeapSort-8                       31671             37969 ns/op

BenchmarkSort/Random/10000/Bubble-8                           22          50935646 ns/op
BenchmarkSort/Random/10000/Insertion-8                        33          34924645 ns/op
BenchmarkSort/Random/10000/InsertionShift-8                   94          12691095 ns/op
BenchmarkSort/Random/10000/InsertionBinary-8                  98          12111096 ns/op
BenchmarkSort/Random/10000/Shell-8                          1563            761198 ns/op
BenchmarkSort/Random/10000/SelectionSort-8                    33          34442337 ns/op
BenchmarkSort/Random/10000/HeapSort-8                       1705            693039 ns/op

BenchmarkSort/Random/100000/Bubble-8                           1        11113416042 ns/op
BenchmarkSort/Random/100000/Insertion-8                        1        3634008375 ns/op
BenchmarkSort/Random/100000/InsertionShift-8                   1        1245496791 ns/op
BenchmarkSort/Random/100000/InsertionBinary-8                  1        1221273875 ns/op
BenchmarkSort/Random/100000/Shell-8                          100          11347430 ns/op
BenchmarkSort/Random/100000/SelectionSort-8                    1        3439591917 ns/op
BenchmarkSort/Random/100000/HeapSort-8                       129           9284771 ns/op

BenchmarkSort/Digits/100/Bubble-8                         308163              3873 ns/op
BenchmarkSort/Digits/100/Insertion-8                      324996              3663 ns/op
BenchmarkSort/Digits/100/InsertionShift-8                 584905              2044 ns/op
BenchmarkSort/Digits/100/InsertionBinary-8                757017              1543 ns/op
BenchmarkSort/Digits/100/Shell-8                         1774785               669.6 ns/op
BenchmarkSort/Digits/100/SelectionSort-8                  276399              4324 ns/op
BenchmarkSort/Digits/100/HeapSort-8                       751316              1583 ns/op

BenchmarkSort/Digits/1000/Bubble-8                          2752            431948 ns/op
BenchmarkSort/Digits/1000/Insertion-8                       3651            328555 ns/op
BenchmarkSort/Digits/1000/InsertionShift-8                  8308            144305 ns/op
BenchmarkSort/Digits/1000/InsertionBinary-8                10000            115256 ns/op
BenchmarkSort/Digits/1000/Shell-8                         102111             10097 ns/op
BenchmarkSort/Digits/1000/SelectionSort-8                   3374            357164 ns/op
BenchmarkSort/Digits/1000/HeapSort-8                       40606             29392 ns/op

BenchmarkSort/Digits/10000/Bubble-8                           24          49069455 ns/op
BenchmarkSort/Digits/10000/Insertion-8                        37          32204625 ns/op
BenchmarkSort/Digits/10000/InsertionShift-8                  100          11426217 ns/op
BenchmarkSort/Digits/10000/InsertionBinary-8                 100          11083692 ns/op
BenchmarkSort/Digits/10000/Shell-8                          3806            308142 ns/op
BenchmarkSort/Digits/10000/SelectionSort-8                    34          34283924 ns/op
BenchmarkSort/Digits/10000/HeapSort-8                       2172            551527 ns/op

BenchmarkSort/Digits/100000/Bubble-8                           1        10715580458 ns/op
BenchmarkSort/Digits/100000/Insertion-8                        1        3256822541 ns/op
BenchmarkSort/Digits/100000/InsertionShift-8                   1        1114596125 ns/op
BenchmarkSort/Digits/100000/InsertionBinary-8                  1        1088388333 ns/op
BenchmarkSort/Digits/100000/Shell-8                          333           3582887 ns/op
BenchmarkSort/Digits/100000/SelectionSort-8                    1        3395311125 ns/op
BenchmarkSort/Digits/100000/HeapSort-8                       196           6077559 ns/op

BenchmarkSort/Sorted/100/Bubble-8                         416689              2843 ns/op
BenchmarkSort/Sorted/100/Insertion-8                     7617646               155.3 ns/op
BenchmarkSort/Sorted/100/InsertionShift-8                1407574               853.9 ns/op
BenchmarkSort/Sorted/100/InsertionBinary-8               6553741               180.6 ns/op
BenchmarkSort/Sorted/100/Shell-8                         2340004               514.4 ns/op
BenchmarkSort/Sorted/100/SelectionSort-8                  328008              3646 ns/op
BenchmarkSort/Sorted/100/HeapSort-8                       666890              1773 ns/op

BenchmarkSort/Sorted/1000/Bubble-8                          4852            247437 ns/op
BenchmarkSort/Sorted/1000/Insertion-8                     881798              1300 ns/op
BenchmarkSort/Sorted/1000/InsertionShift-8                104047             11514 ns/op
BenchmarkSort/Sorted/1000/InsertionBinary-8               667864              1622 ns/op
BenchmarkSort/Sorted/1000/Shell-8                         169585              7054 ns/op
BenchmarkSort/Sorted/1000/SelectionSort-8                   3663            326880 ns/op
BenchmarkSort/Sorted/1000/HeapSort-8                       30733             39175 ns/op

BenchmarkSort/Sorted/10000/Bubble-8                           48          24017171 ns/op
BenchmarkSort/Sorted/10000/Insertion-8                    103312             11381 ns/op
BenchmarkSort/Sorted/10000/InsertionShift-8                 7807            151672 ns/op
BenchmarkSort/Sorted/10000/InsertionBinary-8               82224             14467 ns/op
BenchmarkSort/Sorted/10000/Shell-8                         10000            100812 ns/op
BenchmarkSort/Sorted/10000/SelectionSort-8                    37          32156830 ns/op
BenchmarkSort/Sorted/10000/HeapSort-8                       2324            513663 ns/op

BenchmarkSort/Sorted/100000/Bubble-8                           1        2408639625 ns/op
BenchmarkSort/Sorted/100000/Insertion-8                    10000            109725 ns/op
BenchmarkSort/Sorted/100000/InsertionShift-8                 610           1960857 ns/op
BenchmarkSort/Sorted/100000/InsertionBinary-8               8218            141080 ns/op
BenchmarkSort/Sorted/100000/Shell-8                          892           1248423 ns/op
BenchmarkSort/Sorted/100000/SelectionSort-8                    1        3210073083 ns/op
BenchmarkSort/Sorted/100000/HeapSort-8                       196           6080711 ns/op

BenchmarkSort/Reversed/100/Bubble-8                       163431              7334 ns/op
BenchmarkSort/Reversed/100/Insertion-8                    167882              7140 ns/op
BenchmarkSort/Reversed/100/InsertionShift-8               325568              3677 ns/op
BenchmarkSort/Reversed/100/InsertionBinary-8              409015              2910 ns/op
BenchmarkSort/Reversed/100/Shell-8                       1810816               663.8 ns/op
BenchmarkSort/Reversed/100/SelectionSort-8                296198              3924 ns/op
BenchmarkSort/Reversed/100/HeapSort-8                     793658              1511 ns/op

BenchmarkSort/Reversed/1000/Bubble-8                        1622            755112 ns/op
BenchmarkSort/Reversed/1000/Insertion-8                     1629            715371 ns/op
BenchmarkSort/Reversed/1000/InsertionShift-8                4570            261555 ns/op
BenchmarkSort/Reversed/1000/InsertionBinary-8               4814            248554 ns/op
BenchmarkSort/Reversed/1000/Shell-8                       113557             10543 ns/op
BenchmarkSort/Reversed/1000/SelectionSort-8                 3496            342058 ns/op
BenchmarkSort/Reversed/1000/HeapSort-8                     37742             31756 ns/op

BenchmarkSort/Reversed/10000/Bubble-8                         15          72159467 ns/op
BenchmarkSort/Reversed/10000/Insertion-8                      16          70145305 ns/op
BenchmarkSort/Reversed/10000/InsertionShift-8                 49          24223315 ns/op
BenchmarkSort/Reversed/10000/InsertionBinary-8                49          24089332 ns/op
BenchmarkSort/Reversed/10000/Shell-8                        8130            144847 ns/op
BenchmarkSort/Reversed/10000/SelectionSort-8                  36          33202059 ns/op
BenchmarkSort/Reversed/10000/HeapSort-8                     2293            520586 ns/op

BenchmarkSort/Reversed/100000/Bubble-8                         1        7435156458 ns/op
BenchmarkSort/Reversed/100000/Insertion-8                      1        7307924000 ns/op
BenchmarkSort/Reversed/100000/InsertionShift-8                 1        2435383709 ns/op
BenchmarkSort/Reversed/100000/InsertionBinary-8                1        2420780959 ns/op
BenchmarkSort/Reversed/100000/Shell-8                        619           1932435 ns/op
BenchmarkSort/Reversed/100000/SelectionSort-8                  1        3313458000 ns/op
BenchmarkSort/Reversed/100000/HeapSort-8                     196           6087047 ns/op
```
