package prototype

import "fmt"

func CloneFileANDFolder() {
  file1 := &File{name: "File1",extension:".png"}
  file2 := &File{name: "File2",extension:".mp4"}
  file3 := &File{name: "main",extension:".go"}

    folder1 := &Folder{
        children: []INode{file1},
        name:     "Folder1",
    }

    folder2 := &Folder{
        children: []INode{folder1, file2, file3},
        name:     "Folder2",
    }
    fmt.Println("\nPrinting hierarchy for Folder2")
    folder2.print("  ")

    cloneFolder := folder2.clone()
    fmt.Println("\nPrinting hierarchy for clone Folder")
    cloneFolder.print("  ")
}

