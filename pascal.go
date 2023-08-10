package main
import "fmt"
func main() {
var s int
fmt.Scanln(&s)
 a := make([][]int,s)
for i:= range a {
 a[i] = make([]int,i+1)
}
b:=1
a[0][0]=b
a[1][0]=b
a[1][1]=b
for i:=2;i<s;i++ {
 a[i][0]=b
 a[i][i]=b
 for j:=0;j<i-1;j++ {
  
 a[i][j+1]=a[i-1][j]+a[i-1][j+1]

 }
 }
fmt.Println(a)
for i:=0;i<s;i++ {
 for k:=0;k<s*2/2-i;k++{
 fmt.Print(" ")
 }
 for j:=0;j<=i;j++ {
 fmt.Print(a[i][j])

 fmt.Print(" ")
 }
fmt.Println()
 }
}

