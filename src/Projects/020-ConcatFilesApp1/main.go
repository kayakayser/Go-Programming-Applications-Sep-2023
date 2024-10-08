/*
------------------------------------------------------------------------------------------------------------------------

	Sınıf Çalışması: Komut satırından aşağıdaki gibi çalışan programı yazınız
	     ./csd_cf <file1> <file2> ... <fileN> <dest_path>

	 - Program file1, file2, file3, ..., fileN ile belirtilen dosyaları birleştirerek dest ile belirtilen dosyaya
	 ekleyecektir

	 - Program olmayan kaynak dosyalar için uygun mesajı verecek, akışına devam edecektir

	 - Hedef dosya varsa üzerine yazılacaktır (overwrite)

	 ~/src/Projects/020-ConcatFilesApp1 projesini inceleyiniz

------------------------------------------------------------------------------------------------------------------------
*/
package main

import (
	"020-ConcatFilesApp/csd/err"
	"fmt"
	"io"
	"os"
)

func concatFiles(fd *os.File) {
	defer func() { _ = fd.Close() }()
	n := len(os.Args) - 1

	for i := 2; i < n; i++ {
		fs, e := os.Open(os.Args[i])

		if e != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Can not open file:%s\n", os.Args[i])
			continue
		}
		_, e = io.Copy(fd, fs)

		if e != nil {
			err.ExitFailureError("Copy Problem", e)
		}
	}
}

func main() {
	if len(os.Args) < 3 {
		err.ExitFailure("usage: csd_cf <file1> <file2> ... <fileN> <dest_path>")
	}

	fd, e := os.OpenFile(os.Args[len(os.Args)-1], os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)

	if e != nil {
		err.ExitFailureError("OpenFile", e)
	}

	concatFiles(fd)
}
