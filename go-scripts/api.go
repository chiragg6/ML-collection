
package main

import "github.com/gin-gonic/gin"


import "fmt"



import "os/exec"

import "os"
import "log"
import "io/ioutil"



var router *gin.Engine

func main() {
   
   router = gin.Default()
   initializeRoutes()
   router.Run(":5000")
}

func initializeRoutes(){
   
   router.GET("/api/ipfs", IPFS_config)
   router.GET("/api/swarmKey", Swarm_key)
   router.GET("/api/swarmPeers", Swarm_peers)
   
}



func IPFS_config(c *gin.Context) {
   file, err := os.Open("/root/.ipfs/config")
   if err != nil {
      log.Fatal(err)
   }
   data, err := ioutil.ReadAll(file)
   if err != nil {
      log.Fatal(err)
   }


   fmt.Printf("Data as string: %s\n", data)
   fmt.Println("Number of bytes read:", len(data))
}

func Swarm_key(c *gin.Context){
    file, err := os.Open("/root/.ipfs/swarm.key")
   if err != nil {
      log.Fatal(err)
   }
   data, err := ioutil.ReadAll(file)
   if err != nil {
      log.Fatal(err)
   }


   fmt.Printf("Data as string: %s\n", data)
   fmt.Println("Number of bytes read:", len(data))
}


func Swarm_peers(c *gin.Context){
     out, err := exec.Command("ipfs", "swarm", "peers").Output()
        if err != nil {
                fmt.Printf("%s", err)
        }

        fmt.Println("Command Successfully Executed")

        output := string(out[:])

        fmt.Println(output)
}