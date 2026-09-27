package main

import (
	"fmt"
	"storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	file, err := st.Upload("example.txt", []byte("Hello, World!"))
	if err != nil {
		fmt.Println("Error uploading file:", err)
		return
	}
	restoredFile, errS := st.GetByUUID(file.UUID)
	if errS != nil {
		fmt.Println("Error retrieving file:", errS)
		return
	}
	fmt.Println("storage: ", restoredFile)
}
