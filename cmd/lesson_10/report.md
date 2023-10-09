### Report

```shell
Test random
insert 1000 items at 290.375µs
search 100 items at 14.583µs
remove 100 items at 25.417µs
insert 10000 items at 3.227083ms
search 1000 items at 136.625µs
remove 1000 items at 247.625µs
insert 100000 items at 36.8245ms
search 10000 items at 1.171917ms
remove 10000 items at 2.270958ms
insert 1000000 items at 379.745625ms
search 100000 items at 10.320417ms
remove 100000 items at 29.446542ms
insert 10000000 items at 9.537502166s
search 1000000 items at 218.15175ms
remove 1000000 items at 694.727167ms
insert 100000000 items at 2m44.525269583s
search 10000000 items at 7.957623s
remove 10000000 items at 14.876329084s
Test asc ordered
insert 1000 items at 3.7865ms
search 100 items at 6.791µs
remove 100 items at 500ns
insert 10000 items at 391.300167ms
search 1000 items at 431.75µs
remove 1000 items at 2.881209ms
insert 100000 items at 39.815045959s
search 10000 items at 55.659917ms
remove 10000 items at 368.39875ms
^C
```

### Random items

| op            |      10^3 |       10^4 |        10^5 |         10^6 |          10^7 |            10^8 |
|:--------------|----------:|-----------:|------------:|-------------:|--------------:|----------------:|
| Insert        | 290.375µs | 3.227083ms |   36.8245ms | 379.745625ms |  9.537502166s | 2m44.525269583s |
| Search (N/10) |  14.583µs |  136.625µs |  1.171917ms |  10.320417ms |   218.15175ms |       7.957623s |
| Remove (N/10) |  25.417µs |  247.625µs |  2.270958ms |  29.446542ms |  694.727167ms |   14.876329084s |

### Asc ordered items

| op            |     10^3 |         10^4 |          10^5 |    10^6 | 10^7 | 10^8 |
|:--------------|---------:|-------------:|--------------:|--------:|-----:|-----:|
| Insert        | 3.7865ms | 391.300167ms | 39.815045959s | timeout |    - |    - |
| Search (N/10) |  6.791µs |     431.75µs |   55.659917ms |       - |    - |    - |
| Remove (N/10) |    500ns |   2.881209ms |   368.39875ms |       - |    - |    - |

### Выводы
- в случае со вставкой по возрастанию происходит перекос дерева направо
- сложность всех операций O(N)