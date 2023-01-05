package composite ;

import "fmt";

type Folder struct{
  components []Component
  name         string
}

func (folder *Folder) search(keyword string){
  fmt.Printf("Searching recursively for keyword %s in folder %s\n",keyword,folder.name);
  searchRecursively(folder, keyword);
}

func (folder *Folder) add(component Component){
  folder.components = append(folder.components,component);
}

func searchRecursively(folder *Folder, keyword string){
  for _, composite := range folder.components {
    composite.search(keyword);
  }
}


