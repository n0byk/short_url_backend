package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"os"
)

func main() {
	var buf bytes.Buffer // на месте bytes.Buffer может быть любой io.Writer, например http.ResponseWriter
	w := gzip.NewWriter(&buf)
	_, err := w.Write([]byte("http://jwt.io")) // здесь мы записываем строку. Так как buf у нас обернут в gzip, в него запишется уже сжатая строка
	if err != nil {
		panic(err)
	}
	w.Close() // обязательно закрываем io.WriteCloser, который возвращает compress/gzip

	// сжатая строка
	fmt.Println(buf.String())

	r, err := gzip.NewReader(&buf)
	if err != nil {
		panic(err)
	}
	_ = os.WriteFile("./1.txt", buf.Bytes(), 0644)
	bytes := make([]byte, 100)
	r.Read(bytes)
	// исхоная строка
	fmt.Println(string(bytes))
}
