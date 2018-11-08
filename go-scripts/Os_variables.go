package main

import (
 

  
  "github.com/gin-gonic/gin"

 "fmt"
 "log"
 "os"
 "io/ioutil"
 "net/http"
 "strings"
 "syscall"
)

func main() {
  // Set the router as the default one shipped with Gin
  router := gin.Default()

  // Serve frontend static files
 

  /*Setting route group with name API*/
  api := router.Group("/api")
  {
    api.GET("/", func(c *gin.Context) {
     var uname syscall.Utsname
   c.Param(arrayToString(uname.Nodename))
   c.Param(arrayToString(uname.Release))
   c.Param(arrayToString(uname.Sysname))
   c.Param(arrayToString(uname.Version))
   c.Param(arrayToString(uname.Machine))
   c.Param(arrayToString(uname.Domainname))
})


    

    api.GET("/config", func(c *gin.Context){
        file, err := os.Open("/root/.ipfs/config")
   if err != nil {
      log.Fatal(err)
   }
   data, err := ioutil.ReadAll(file)
   if err != nil {
      log.Fatal(err)
   }
   c.JSON(http.StatusOK, gin.H{
    "code": http.StatusOK,"message": string(data),
    })
   
    })

    /*Anotehr GET Route for getting Swarm Key from IPFS Admin*/
    api.GET("/key", func(c *gin.Context){
      file, err := os.Open("/root/.ipfs/swarm.key")
   if err != nil {
      log.Fatal(err)
   }
   key, err := ioutil.ReadAll(file)
   if err != nil {
      log.Fatal(err)
   }
     c.JSON(http.StatusOK, gin.H{
    "code" : http.StatusOK,
    "message": string(key),
    })
    })
  



  // Start and run the server
  /*If required to define custom port, define below like router.RUN(":5000)*/
  /*By Default it would be running on port 3000, and have used gin web proxy for go, to read all the changes to go code automatically*/
  /*Tool for go server automation - https://github.com/codegangsta/gin*/
  router.Run()
}
}
func arrayToString(x [65]int8) string {
   var buf [65]byte
   for i, b := range x {
      buf[i] = byte(b)
   }
   str := string(buf[:])
   if i := strings.Index(str, "\x00"); i != -1 {
      str = str[:i]
   }
   return str
}