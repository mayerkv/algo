package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/mayerkv/algo/pkg/sort"
)

func splitChunks[T sort.Ordered](in io.Reader, chunkSize int, sorter sort.Sorter[T]) ([]io.ReadSeekCloser, error) {
	// создаем список для хранения чанков
	out := make([]io.ReadSeekCloser, 0, 10)

	// пока можно читать источник
	hasMore := true
	for hasMore {
		// создаем чанк
		chunk := make([]T, 0, chunkSize)
		// пока не заполнился чанк
		for len(chunk) < chunkSize {
			// читаем источник
			var value T
			_, err := fmt.Fscanln(in, &value)
			// если источник закончился завершаем чтение
			if errors.Is(err, io.EOF) {
				hasMore = false
				break
			}
			checkError(err)
			// кладем значение в чанк
			chunk = append(chunk, value)
		}
		// если чанк пустой, то завершаем
		if len(chunk) == 0 {
			break
		}
		// сортируем чанк
		sorter(chunk)
		// создаем файл
		f := openFile(fmt.Sprintf("tmp.%d.txt", len(out)), os.O_CREATE|os.O_TRUNC|os.O_RDWR)
		// пишем чанк в него
		writer := bufio.NewWriter(f)
		for _, el := range chunk {
			_, err := fmt.Fprintln(writer, el)
			checkError(err)
		}
		err := writer.Flush()
		checkError(err)
		// кладем файл в список
		out = append(out, f)
	}
	return out, nil
}
