package main

import (
	"os"
)

func main() {
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("file name:", fileInfo.Name())
	// fmt.Println("is folder:", fileInfo.IsDir())
	// fmt.Println("filesize:", fileInfo.Size())
	// fmt.Println("file modified at:", fileInfo.ModTime())

	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }

	// buf := make([]byte, fileInfo.Size())
	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < len(buf); i++ {
	// 	fmt.Println("Data", d, string(buf[i]))
	// }

	// data, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(data))

	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(0)

	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	// f, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// f.WriteString("hi go")
	// f.WriteString("nice language")

	// bytes := []byte("Hello golang")

	// f.Write(bytes)

	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destFile, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}

	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }

	// writer.Flush()

	// fmt.Println("written to new line successfully")

	err := os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}
}