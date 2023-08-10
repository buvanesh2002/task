package main
import "fmt"

func main() {
   
    s:="hi go run gogo hello"
    fmt.Println(s)
    len:=len(s)
    fmt.Println(len)
    var count int
    a :=[]byte(s)
    for i:=0;i<len;i++ {
        if a[i]=='G'||a[i]=='g'{
            if a[i+1]=='O'||a[i+1]=='o'{
                count++
            }
        }
    }
    fmt.Println(count)
}