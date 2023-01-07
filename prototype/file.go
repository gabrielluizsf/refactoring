package prototype

import "fmt"

type File struct{
  name      string
  extension string
}

func (file *File) print(indentation string){
  fmt.Println(indentation + file.name + file.extension);
}
func (file *File) clone() INode{
  return &File{name: file.name + "_clone" + file.extension};
}
