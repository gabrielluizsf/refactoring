package proxy

import "fmt";

func INITServer(){
  nginxServer   := newNginxServer();
  appStatusURL  := "/app/status";
  createUserURL := "/create/user";

  httpCode1, body1 := nginxServer.handleRequest(appStatusURL,"GET");
    printHTTPRequest(appStatusURL,body1,httpCode1);
  httpCode2, body2 := nginxServer.handleRequest(appStatusURL,"POST");
    printHTTPRequest(createUserURL,body2,httpCode2);
   httpCode3, body3 := nginxServer.handleRequest(appStatusURL,"GET");
    printHTTPRequest(appStatusURL,body3,httpCode3);
   httpCode4, body4 := nginxServer.handleRequest(appStatusURL,"GET");
    printHTTPRequest(appStatusURL,body4,httpCode4);
   httpCode5, body5 := nginxServer.handleRequest(createUserURL,"POST");   
    printHTTPRequest(createUserURL,body5,httpCode5);
   httpCode6, body6 := nginxServer.handleRequest(appStatusURL,"GET");
    printHTTPRequest(appStatusURL,body6,httpCode6);
   httpCode7, body7 := nginxServer.handleRequest(appStatusURL,"GET");
    printHTTPRequest(appStatusURL,body7,httpCode7);
}
func printHTTPRequest(appStatusURL, body string ,httpCode int){
  fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body);
}
