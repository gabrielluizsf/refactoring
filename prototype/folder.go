package prototype

import "fmt"

type Folder struct{
  children []INode
  name       string
}

func (folder *Folder) print(indentation string){
  fmt.Println(indentation + folder.name);
  for _, value := range folder.children{
    value.print(indentation + indentation);
  }
}

func (folder *Folder) clone() INode{
  cloneFolder := &Folder{name: folder.name + "_clone"};
  var tempChildren []INode;
  for _, value := range folder.children{
    cp := value.clone();
    tempChildren = append(tempChildren, cp);
  }
  cloneFolder.children = tempChildren;
  return cloneFolder;
}
