package proxy

type Application struct{};

func (a *Application) handleRequest(url, method string) (int, string){
  if url == "/app/status" && method == "GET" {
    return 200, "OK";
  }else if url == "/create/user" && method == "POST" {
    return 201, "User Created";
  }else{
  return 404, "NOT OK";
  }
}
