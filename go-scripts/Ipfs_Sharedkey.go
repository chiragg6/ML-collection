package main


    import "crypto/rand"
import "encoding/hex"
import "fmt"
import "log"
import "os"
import "io"


var path = "swarm.txt"

func main() {
    privatekey := keyGenrator()
    createFile()
    writeFile(privatekey)
    readFile()
    
}

func keyGenrator() (string){
    var sharedkey string
    key := make([]byte, 32)
    _, err := rand.Read(key)
    if err != nil {
        log.Fatalln("While trying to read random source:", err)
    }

    fmt.Println("/key/swarm/psk/1.0.0/")
    fmt.Println("/base16/")
    sharedkey = hex.EncodeToString(key)
    fmt.Print(sharedkey)
    return sharedkey
    }


func createFile() {
    
    var _, err = os.Stat(path)

    
    if os.IsNotExist(err) {
        var file, err = os.Create(path)
        if isError(err) {
            return
        }
        defer file.Close()
    }

    fmt.Println("File Created Successfully", path)
}

func writeFile(privatekey string) {
    
    var file, err = os.OpenFile(path, os.O_RDWR, 0644)
    if isError(err) {
        return
    }
    defer file.Close()

    
    _, err = file.WriteString(privatekey)
    if isError(err) {
        return
    }

    
    err = file.Sync()
    if isError(err) {
        return
    }

    fmt.Println("File Updated Successfully.")
}

func readFile() {
    
    var file, err = os.OpenFile(path, os.O_RDWR, 0644)
    if isError(err) {
        return
    }
    defer file.Close()

    
    var text = make([]byte, 1024)
    for {
        _, err = file.Read(text)

        
        if err == io.EOF {
            break
        }

        
        if err != nil && err != io.EOF {
            isError(err)
            break
        }
    }

    fmt.Println("Reading from file.")
    fmt.Println(string(text))
}

func isError(err error) bool {
    if err != nil {
        fmt.Println(err.Error())
    }

    return (err != nil)
}
