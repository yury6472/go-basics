/*
Данная задача поможет вам разобраться в пакете encoding/csv и path/filepath, хотя для решения может 
быть использован также пакет archive/zip (поскольку файл с заданием предоставляется именно в этом формате).

В тестовом архиве, который мы можете скачать из нашего репозитория на github.com, содержится набор папок и файлов.
 Один из этих файлов является файлом с данными в формате CSV, прочие же файлы структурированных данных не содержат.

Требуется найти и прочитать этот единственный файл со структурированными данными (это таблица 10х10, 
разделителем является запятая), а в качестве ответа необходимо указать число, 
находящееся на 5 строке и 3 позиции (индексы 4 и 2 соответственно).

*/


package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {

	fmt.Println("List by ReadDir")
	listDirByReadDir(".")
	fmt.Println()
	fmt.Println("List by Walk")
	listDirByWalk(".")
}

func listDirByWalk(path string) {
	filepath.Walk(path, func(wPath string, info os.FileInfo, err error) error {

		// Обход директории без вывода
		if wPath == path {
			return nil
		}

		// Если данный путь является директорией, то останавливаем рекурсивный обход
		// и возвращаем название папки
		/* if info.IsDir() {
			fmt.Printf("[%s]\n", wPath)
			return filepath.SkipDir
		}
		*/

		// Выводится название файла
		if wPath != path {
			//fmt.Println(wPath)
		}

		// читаем файл
		file, err := os.Open(wPath)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		reader := csv.NewReader(file)
		reader.Comma = ','

		record, _ := reader.ReadAll()
		if len(record) == 10 {
			fmt.Println(record[4][2])
		}

		return nil
	})
}

func listDirByReadDir(path string) {
	lst, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, val := range lst {
		if val.IsDir() {
			fmt.Printf("[%s]\n", val.Name())
		} else {
			fmt.Println(val.Name())
		}
	}
}
