package main

import "github.com/gin-gonic/gin"

import "crypto/rand"
import "encoding/hex"
import "fmt"
import "log"
import "syscall"
import "fmt"
import "strings"



var router *gin.Engine

func main() {
	router = gin.Default()
	initializeRoutes()
	router.Run(":5000")
}

func initializeRoutes(){
	router.GET("/api", Os_variables)
	
}

func Os_variables(c *gin.Context) {
   var uname syscall.Utsname
   if err := syscall.Uname(&uname); err != nil {
      fmt.Printf("Uname: %v", err)
   }
   fmt.Println(arrayToString(uname.Nodename))
   fmt.Println(arrayToString(uname.Release))
   fmt.Println(arrayToString(uname.Sysname))
   fmt.Println(arrayToString(uname.Version))
   fmt.Println(arrayToString(uname.Machine))
   fmt.Println(arrayToString(uname.Domainname))
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






}

