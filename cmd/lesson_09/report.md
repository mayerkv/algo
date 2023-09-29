### Тайминги по сортировке массива из случайных чисел

```shell
BucketSort(100): 13.25µs
CountSort(100): 12.125µs
RadixSort(100): 13.917µs

BucketSort(1000): 99.75µs
CountSort(1000): 14.458µs
RadixSort(1000): 39.542µs

BucketSort(10000): 976.291µs
CountSort(10000): 78.667µs
RadixSort(10000): 365.625µs

BucketSort(100000): 9.352917ms
CountSort(100000): 714.917µs
RadixSort(100000): 4.222708ms

BucketSort(1000000): 84.681291ms
CountSort(1000000): 4.767083ms
RadixSort(1000000): 23.424666ms

BucketSort(10000000): 1.318384625s
CountSort(10000000): 24.549708ms
RadixSort(10000000): 121.861209ms

BucketSort(100000000): 14.863349209s
CountSort(100000000): 248.886334ms
RadixSort(100000000): 1.253712416s

BucketSort(1000000000): timeout
CountSort(1000000000): 16.610123958s
RadixSort(1000000000): 45.396958542s
```

### Сравнительная таблица времени выполнения

| Name            |      100 |      1_000 |      10_000 |       100_000 |    1_000_000 |    10_000_000 |   100_000_000 | 1_000_000_000 |
|-----------------|---------:|-----------:|------------:|--------------:|-------------:|--------------:|--------------:|--------------:|
| Bubble          | 17.333µs | 1.825166ms | 81.379916ms | 10.869929458s |      timeout |               |               |               | 
| Insertion       |  8.167µs | 1.160583ms | 36.503167ms |  3.588174458s |      timeout |               |               |               | 
| InsertionShift  | 20.792µs |  513.208µs | 12.762583ms |  1.230636792s |      timeout |               |               |               |
| InsertionBinary |  8.166µs |  367.792µs |  12.13625ms |     1.194615s |      timeout |               |               |               |
| Shell           | 13.042µs |  153.791µs |   776.375µs |   11.221417ms | 199.692333ms |               |               |               |
| Selection       | 16.834µs |  1.09975ms | 61.011834ms |  3.413933083s |      timeout |               |               |               |
| Heap            | 12.459µs |  168.167µs |   868.541µs |    9.102375ms |  125.12575ms |               |               |               |
| Quick           |  5.666µs |   39.875µs |   491.458µs |    5.502666ms |  65.843083ms |  739.014958ms |               |               |
| Merge           |     32µs |   74.916µs |   885.542µs |   10.841417ms | 108.186542ms |     1.229479s |               |               |
| Bucket          |  13.25µs |    99.75µs |   976.291µs |    9.352917ms |  84.681291ms |  1.318384625s | 14.863349209s |       timeout |
| Count           | 12.125µs |   14.458µs |    78.667µs |     714.917µs |   4.767083ms |   24.549708ms |  248.886334ms | 16.610123958s |
| Radix           | 13.917µs |   39.542µs |   365.625µs |    4.222708ms |  23.424666ms |  121.861209ms |  1.253712416s | 45.396958542s |

### Бенчмарки для разных характеров данных

```shell
BenchmarkSort/Random/100/Bubble-8                         304724              3924 ns/op
BenchmarkSort/Random/100/Insertion-8                      325359              3693 ns/op
BenchmarkSort/Random/100/InsertionShift-8                 612372              1978 ns/op
BenchmarkSort/Random/100/InsertionBinary-8                811563              1473 ns/op
BenchmarkSort/Random/100/Shell-8                         1628013               742.9 ns/op
BenchmarkSort/Random/100/SelectionSort-8                  278622              4358 ns/op
BenchmarkSort/Random/100/HeapSort-8                       759104              1584 ns/op
BenchmarkSort/Random/100/MergeSort-8                      342204              3418 ns/op
BenchmarkSort/Random/100/QuickSort-8                     1296409               930.3 ns/op
BenchmarkSort/Random/100/BucketSort-8                     360672              3274 ns/op
BenchmarkSort/Random/100/CountSort-8                      394225              2998 ns/op
BenchmarkSort/Random/100/RadixSort-8                      897142              1318 ns/op

BenchmarkSort/Random/1000/Bubble-8                          2527            465872 ns/op
BenchmarkSort/Random/1000/Insertion-8                       3350            358219 ns/op
BenchmarkSort/Random/1000/InsertionShift-8                  7447            161650 ns/op
BenchmarkSort/Random/1000/InsertionBinary-8                 9566            125218 ns/op
BenchmarkSort/Random/1000/Shell-8                          33156             33265 ns/op
BenchmarkSort/Random/1000/SelectionSort-8                   3338            366029 ns/op
BenchmarkSort/Random/1000/HeapSort-8                       35358             32797 ns/op
BenchmarkSort/Random/1000/MergeSort-8                      25352             46875 ns/op
BenchmarkSort/Random/1000/QuickSort-8                      74138             14903 ns/op
BenchmarkSort/Random/1000/BucketSort-8                     37519             31900 ns/op
BenchmarkSort/Random/1000/CountSort-8                     242538              4900 ns/op
BenchmarkSort/Random/1000/RadixSort-8                      99420             11986 ns/op

BenchmarkSort/Random/10000/Bubble-8                           24          49522984 ns/op
BenchmarkSort/Random/10000/Insertion-8                        33          34789965 ns/op
BenchmarkSort/Random/10000/InsertionShift-8                   94          12579974 ns/op
BenchmarkSort/Random/10000/InsertionBinary-8                  99          11902969 ns/op
BenchmarkSort/Random/10000/Shell-8                          1737            686347 ns/op
BenchmarkSort/Random/10000/SelectionSort-8                    34          34158027 ns/op
BenchmarkSort/Random/10000/HeapSort-8                       1758            679281 ns/op
BenchmarkSort/Random/10000/MergeSort-8                      1518            780659 ns/op
BenchmarkSort/Random/10000/QuickSort-8                      3001            389916 ns/op
BenchmarkSort/Random/10000/BucketSort-8                     4611            249719 ns/op
BenchmarkSort/Random/10000/CountSort-8                     51099             23384 ns/op
BenchmarkSort/Random/10000/RadixSort-8                     10000            113623 ns/op

BenchmarkSort/Random/100000/Bubble-8                           1        11000736959 ns/op
BenchmarkSort/Random/100000/Insertion-8                        1        3632856042 ns/op
BenchmarkSort/Random/100000/InsertionShift-8                   1        1248031209 ns/op
BenchmarkSort/Random/100000/InsertionBinary-8                  1        1213353583 ns/op
BenchmarkSort/Random/100000/Shell-8                          132           9005752 ns/op
BenchmarkSort/Random/100000/SelectionSort-8                    1        3394357208 ns/op
BenchmarkSort/Random/100000/HeapSort-8                       134           8886327 ns/op
BenchmarkSort/Random/100000/MergeSort-8                      128           9278930 ns/op
BenchmarkSort/Random/100000/QuickSort-8                      147           8094752 ns/op
BenchmarkSort/Random/100000/BucketSort-8                     409           2912725 ns/op
BenchmarkSort/Random/100000/CountSort-8                     5474            213711 ns/op
BenchmarkSort/Random/100000/RadixSort-8                     1034           1150908 ns/op

BenchmarkSort/Digits/100/Bubble-8                         307005              3645 ns/op
BenchmarkSort/Digits/100/Insertion-8                      334366              3555 ns/op
BenchmarkSort/Digits/100/InsertionShift-8                 615330              1944 ns/op
BenchmarkSort/Digits/100/InsertionBinary-8                855304              1398 ns/op
BenchmarkSort/Digits/100/Shell-8                         1995555               600.6 ns/op
BenchmarkSort/Digits/100/SelectionSort-8                  271051              4480 ns/op
BenchmarkSort/Digits/100/HeapSort-8                       783022              1525 ns/op
BenchmarkSort/Digits/100/MergeSort-8                      335016              3497 ns/op
BenchmarkSort/Digits/100/QuickSort-8                      974859              1242 ns/op
BenchmarkSort/Digits/100/BucketSort-8                     516151              2296 ns/op
BenchmarkSort/Digits/100/CountSort-8                     3489831               344.6 ns/op
BenchmarkSort/Digits/100/RadixSort-8                     2540294               472.2 ns/op

BenchmarkSort/Digits/1000/Bubble-8                          2775            427763 ns/op
BenchmarkSort/Digits/1000/Insertion-8                       3602            327658 ns/op
BenchmarkSort/Digits/1000/InsertionShift-8                  8240            142811 ns/op
BenchmarkSort/Digits/1000/InsertionBinary-8                10000            113627 ns/op
BenchmarkSort/Digits/1000/Shell-8                         112080              9742 ns/op
BenchmarkSort/Digits/1000/SelectionSort-8                   3380            353845 ns/op
BenchmarkSort/Digits/1000/HeapSort-8                       43168             26560 ns/op
BenchmarkSort/Digits/1000/MergeSort-8                      27385             43582 ns/op
BenchmarkSort/Digits/1000/QuickSort-8                      22785             52626 ns/op
BenchmarkSort/Digits/1000/BucketSort-8                     55826             21409 ns/op
BenchmarkSort/Digits/1000/CountSort-8                     420766              2835 ns/op
BenchmarkSort/Digits/1000/RadixSort-8                     277514              4294 ns/op

BenchmarkSort/Digits/10000/Bubble-8                           25          47122042 ns/op
BenchmarkSort/Digits/10000/Insertion-8                        37          31890375 ns/op
BenchmarkSort/Digits/10000/InsertionShift-8                  100          11372892 ns/op
BenchmarkSort/Digits/10000/InsertionBinary-8                 100          10981432 ns/op
BenchmarkSort/Digits/10000/Shell-8                          3915            303848 ns/op
BenchmarkSort/Digits/10000/SelectionSort-8                    34          34050887 ns/op
BenchmarkSort/Digits/10000/HeapSort-8                       2220            538795 ns/op
BenchmarkSort/Digits/10000/MergeSort-8                      1948            603573 ns/op
BenchmarkSort/Digits/10000/QuickSort-8                       297           4020393 ns/op
BenchmarkSort/Digits/10000/BucketSort-8                     5188            228069 ns/op
BenchmarkSort/Digits/10000/CountSort-8                     44505             26879 ns/op
BenchmarkSort/Digits/10000/RadixSort-8                     28950             41452 ns/op

BenchmarkSort/Digits/100000/Bubble-8                           1        10671854791 ns/op
BenchmarkSort/Digits/100000/Insertion-8                        1        3260077333 ns/op
BenchmarkSort/Digits/100000/InsertionShift-8                   1        1115688417 ns/op
BenchmarkSort/Digits/100000/InsertionBinary-8                  1        1092068583 ns/op
BenchmarkSort/Digits/100000/Shell-8                          334           3573749 ns/op
BenchmarkSort/Digits/100000/SelectionSort-8                    1        3400782125 ns/op
BenchmarkSort/Digits/100000/HeapSort-8                       196           6078264 ns/op
BenchmarkSort/Digits/100000/MergeSort-8                      163           7290275 ns/op
BenchmarkSort/Digits/100000/QuickSort-8                        3         381453792 ns/op
BenchmarkSort/Digits/100000/BucketSort-8                     498           2388566 ns/op
BenchmarkSort/Digits/100000/CountSort-8                     4326            274643 ns/op
BenchmarkSort/Digits/100000/RadixSort-8                     2845            420341 ns/op

BenchmarkSort/Sorted/100/Bubble-8                         427372              2803 ns/op
BenchmarkSort/Sorted/100/Insertion-8                    14086786                84.59 ns/op
BenchmarkSort/Sorted/100/InsertionShift-8                1528713               786.1 ns/op
BenchmarkSort/Sorted/100/InsertionBinary-8              10638568               112.4 ns/op
BenchmarkSort/Sorted/100/Shell-8                         2659491               450.8 ns/op
BenchmarkSort/Sorted/100/SelectionSort-8                  339409              3532 ns/op
BenchmarkSort/Sorted/100/HeapSort-8                       709423              1700 ns/op
BenchmarkSort/Sorted/100/MergeSort-8                      369427              3162 ns/op
BenchmarkSort/Sorted/100/QuickSort-8                      239534              4978 ns/op
BenchmarkSort/Sorted/100/BucketSort-8                     365931              3247 ns/op
BenchmarkSort/Sorted/100/CountSort-8                     2257633               533.8 ns/op
BenchmarkSort/Sorted/100/RadixSort-8                     1458667               823.1 ns/op

BenchmarkSort/Sorted/1000/Bubble-8                          4891            245059 ns/op
BenchmarkSort/Sorted/1000/Insertion-8                    1595940               762.1 ns/op
BenchmarkSort/Sorted/1000/InsertionShift-8                108174             10923 ns/op
BenchmarkSort/Sorted/1000/InsertionBinary-8              1000000              1068 ns/op
BenchmarkSort/Sorted/1000/Shell-8                         182828              6545 ns/op
BenchmarkSort/Sorted/1000/SelectionSort-8                   3687            324366 ns/op
BenchmarkSort/Sorted/1000/HeapSort-8                       32553             35740 ns/op
BenchmarkSort/Sorted/1000/MergeSort-8                      33158             36053 ns/op
BenchmarkSort/Sorted/1000/QuickSort-8                       3045            392829 ns/op
BenchmarkSort/Sorted/1000/BucketSort-8                     37323             31979 ns/op
BenchmarkSort/Sorted/1000/CountSort-8                     235908              5066 ns/op
BenchmarkSort/Sorted/1000/RadixSort-8                     111436             10692 ns/op

BenchmarkSort/Sorted/10000/Bubble-8                           49          23988219 ns/op
BenchmarkSort/Sorted/10000/Insertion-8                    152338              7841 ns/op
BenchmarkSort/Sorted/10000/InsertionShift-8                 8151            147287 ns/op
BenchmarkSort/Sorted/10000/InsertionBinary-8              108262             11053 ns/op
BenchmarkSort/Sorted/10000/Shell-8                         12345             97199 ns/op
BenchmarkSort/Sorted/10000/SelectionSort-8                    37          31964369 ns/op
BenchmarkSort/Sorted/10000/HeapSort-8                       2338            510679 ns/op
BenchmarkSort/Sorted/10000/MergeSort-8                      2666            438733 ns/op
BenchmarkSort/Sorted/10000/QuickSort-8                        30          38056082 ns/op
BenchmarkSort/Sorted/10000/BucketSort-8                     3656            320845 ns/op
BenchmarkSort/Sorted/10000/CountSort-8                     24210             49590 ns/op
BenchmarkSort/Sorted/10000/RadixSort-8                      8785            134400 ns/op

BenchmarkSort/Sorted/100000/Bubble-8                           1        2393841792 ns/op
BenchmarkSort/Sorted/100000/Insertion-8                    14830             80850 ns/op
BenchmarkSort/Sorted/100000/InsertionShift-8                 621           1921173 ns/op
BenchmarkSort/Sorted/100000/InsertionBinary-8              10000            112839 ns/op
BenchmarkSort/Sorted/100000/Shell-8                          987           1213167 ns/op
BenchmarkSort/Sorted/100000/SelectionSort-8                    1        3194996375 ns/op
BenchmarkSort/Sorted/100000/HeapSort-8                       196           6092075 ns/op
BenchmarkSort/Sorted/100000/MergeSort-8                      222           5374689 ns/op
BenchmarkSort/Sorted/100000/QuickSort-8                        1        3812958166 ns/op
BenchmarkSort/Sorted/100000/BucketSort-8                     333           3596086 ns/op
BenchmarkSort/Sorted/100000/CountSort-8                     2509            467809 ns/op
BenchmarkSort/Sorted/100000/RadixSort-8                      708           1690882 ns/op

BenchmarkSort/Reversed/100/Bubble-8                       165704              7211 ns/op
BenchmarkSort/Reversed/100/Insertion-8                    170504              7041 ns/op
BenchmarkSort/Reversed/100/InsertionShift-8               326354              3623 ns/op
BenchmarkSort/Reversed/100/InsertionBinary-8              425416              2814 ns/op
BenchmarkSort/Reversed/100/Shell-8                       1983038               602.9 ns/op
BenchmarkSort/Reversed/100/SelectionSort-8                289263              4116 ns/op
BenchmarkSort/Reversed/100/HeapSort-8                     840543              1423 ns/op
BenchmarkSort/Reversed/100/MergeSort-8                    375063              3106 ns/op
BenchmarkSort/Reversed/100/QuickSort-8                    240830              4967 ns/op
BenchmarkSort/Reversed/100/BucketSort-8                   363375              3258 ns/op
BenchmarkSort/Reversed/100/CountSort-8                   2210203               543.1 ns/op
BenchmarkSort/Reversed/100/RadixSort-8                   1458590               823.6 ns/op

BenchmarkSort/Reversed/1000/Bubble-8                        1650            724430 ns/op
BenchmarkSort/Reversed/1000/Insertion-8                     1692            706915 ns/op
BenchmarkSort/Reversed/1000/InsertionShift-8                4580            258564 ns/op
BenchmarkSort/Reversed/1000/InsertionBinary-8               4882            245000 ns/op
BenchmarkSort/Reversed/1000/Shell-8                       119662              9986 ns/op
BenchmarkSort/Reversed/1000/SelectionSort-8                 3442            350259 ns/op
BenchmarkSort/Reversed/1000/HeapSort-8                     38840             27578 ns/op
BenchmarkSort/Reversed/1000/MergeSort-8                    33360             35739 ns/op
BenchmarkSort/Reversed/1000/QuickSort-8                     3256            366842 ns/op
BenchmarkSort/Reversed/1000/BucketSort-8                   37294             32060 ns/op
BenchmarkSort/Reversed/1000/CountSort-8                   232402              5115 ns/op
BenchmarkSort/Reversed/1000/RadixSort-8                   111115             10729 ns/op

BenchmarkSort/Reversed/10000/Bubble-8                         15          72088328 ns/op
BenchmarkSort/Reversed/10000/Insertion-8                      16          69829979 ns/op
BenchmarkSort/Reversed/10000/InsertionShift-8                 49          24180685 ns/op
BenchmarkSort/Reversed/10000/InsertionBinary-8                49          23992749 ns/op
BenchmarkSort/Reversed/10000/Shell-8                        8498            141406 ns/op
BenchmarkSort/Reversed/10000/SelectionSort-8                  36          33152439 ns/op
BenchmarkSort/Reversed/10000/HeapSort-8                     2289            514802 ns/op
BenchmarkSort/Reversed/10000/MergeSort-8                    2749            434222 ns/op
BenchmarkSort/Reversed/10000/QuickSort-8                      33          35115486 ns/op
BenchmarkSort/Reversed/10000/BucketSort-8                   3609            322259 ns/op
BenchmarkSort/Reversed/10000/CountSort-8                   24564             48754 ns/op
BenchmarkSort/Reversed/10000/RadixSort-8                    8874            135702 ns/op

BenchmarkSort/Reversed/100000/Bubble-8                         1        7399247708 ns/op
BenchmarkSort/Reversed/100000/Insertion-8                      1        7255250125 ns/op
BenchmarkSort/Reversed/100000/InsertionShift-8                 1        2405139666 ns/op
BenchmarkSort/Reversed/100000/InsertionBinary-8                1        2402390750 ns/op
BenchmarkSort/Reversed/100000/Shell-8                        630           1894320 ns/op
BenchmarkSort/Reversed/100000/SelectionSort-8                  1        3295058083 ns/op
BenchmarkSort/Reversed/100000/HeapSort-8                     198           6045068 ns/op
BenchmarkSort/Reversed/100000/MergeSort-8                    223           5338595 ns/op
BenchmarkSort/Reversed/100000/QuickSort-8                      1        3501400875 ns/op
BenchmarkSort/Reversed/100000/BucketSort-8                   332           3602775 ns/op
BenchmarkSort/Reversed/100000/CountSort-8                   2486            475761 ns/op
BenchmarkSort/Reversed/100000/RadixSort-8                    699           1709522 ns/op
```
