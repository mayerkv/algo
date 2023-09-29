package main

import (
	"errors"
	"fmt"
	"io"

	"github.com/mayerkv/algo/pkg/sort"
)

func mergeChunks[T sort.Ordered](chunks []io.ReadSeekCloser, out io.Writer) error {
	// создаем список для поиска минимального значения
	arr := make(itemList[T], 0, len(chunks))

	// читаем по одному значению из каждого чанка
	for _, chunk := range chunks {
		_, err := chunk.Seek(0, io.SeekStart) // при первом чтении перекидываем указатель в начало
		checkError(err)
		var value T
		_, err = fmt.Fscanln(chunk, &value)
		if errors.Is(err, io.EOF) {
			break
		}
		checkError(err)
		arr = append(arr, &item[T]{value, chunk})
	}

	// пока список не пустой
	for len(arr) > 0 {
		// сортируем список и берем минимально значение
		sort.HeapSortT[itemList[T]](arr)
		minItem := arr[0]

		// записываем его в выход
		_, err := fmt.Fprintln(out, minItem.value)
		checkError(err)

		// читаем следующее значение из файла с минимальным значением
		var value T
		_, err = fmt.Fscanln(minItem.src, &value)

		// если файл закончился, то закрываем его и убираем из списка
		if errors.Is(err, io.EOF) {
			minItem.src.Close()
			arr = arr[1:]
		} else if err != nil {
			checkError(err)
		} else { // иначе сохраняем новое значение
			minItem.value = value
		}
	}

	return nil
}

type item[T sort.Ordered] struct {
	value T
	src   io.ReadSeekCloser
}

type itemList[T sort.Ordered] []*item[T]

func (l itemList[T]) Len() int {
	return len(l)
}

func (l itemList[T]) Less(i, j int) bool {
	return l[i].value < l[j].value
}

func (l itemList[T]) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
