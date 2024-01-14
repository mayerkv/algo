### Report
```shell
Splay Tree
[SEARCH RANDOM 100] random = 6.032µs ordered = 49.442µs
[SEARCH 1-10  1000] random = 50.504µs ordered = 77.354µs
[REMOVE RANDOM 100] random = 7.655µs ordered = 10.559µs
Treap
[SEARCH RANDOM 100] random = 4.227µs ordered = 4.569µs
[SEARCH 1-10  1000] random = 51.296µs ordered = 76.803µs
[REMOVE RANDOM 100] random = 22.412µs ordered = 18.795µs

Process finished with the exit code 0
```

### Random items

| Test case              |    Splay |    Treap |
|------------------------|---------:|---------:|
| Search random 100      |  6.032µs |  4.227µs |
| Search 1-10 1000 times | 50.504µs | 51.296µs |
| Remove random 100      |  7.655µs | 22.412µs |

### Asc Ordered items

| Test case              |    Splay |    Treap |
|------------------------|---------:|---------:|
| Search random 100      | 49.442µs |  4.569µs |
| Search 1-10 1000 times | 77.354µs | 76.803µs |
| Remove random 100      | 10.559µs | 18.795µs |

### Выводы
- при вставке по возрастанию в расширяющееся дерево переходит перекос налево,
который выравнивается при доступе к элементу
- вставка по возрастанию в рандомизированное дерево не влияет на производительность
- сложность поиска одинаковая
- удаление в расширяющемся дереве несколько быстрее чем в диманическом