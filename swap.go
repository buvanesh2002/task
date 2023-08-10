package main
import "fmt"
func main(){
	str:="ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	len:=len(str)
	b:=[]byte(str)
	var temp byte

	for i:=2;i<len;i+=4{
		temp=b[i+1]
		b[i+1]=b[i]
        b[i]=temp
	}
	str=string(b)
	fmt.Println(str)
}