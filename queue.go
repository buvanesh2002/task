package main

import "fmt"

type queue struct {
    top,bottom,size int
    capacity int 
    arr[]int
}

func Createqueue(capacity int)(*queue){
    return &queue{
        top :0,
        bottom:-1,
        size:0,
        capacity:capacity,
        arr:make([]int,capacity),

    }
}

func(q *queue) enqueue(item int)(bool){
    if(q.size==q.capacity){
        fmt.Println("queue is full")
        return false
    }
    q.bottom=((q.bottom+1)%q.capacity)
    q.arr[q.bottom]=item
    q.size++
    return true
}

func(q *queue) dequeue()(int,bool){
    if(q.size==0){
        fmt.Println("queue is empty")
        return 0,false
    }
    item :=q.arr[q.top]
    q.top=((q.top+1)%q.capacity)
    q.size--
    return item,true
}

func main(){
    q :=Createqueue(5)
    q.enqueue(100)
    q.enqueue(200)
    q.enqueue(300)
    q.enqueue(400)
    q.enqueue(500)
    q.enqueue(1)
    fmt.Println(q.dequeue())
    fmt.Println(q.dequeue())
    fmt.Println(q.dequeue())
    fmt.Println(q.dequeue())
    fmt.Println(q.dequeue())

}