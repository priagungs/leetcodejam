type queue []int

func ConstructQueue() queue {
    return queue(make([]int, 0, 100))
}

func (q *queue) Push(x int) {
    *q = append(*q, x)
}

func (q *queue) Pop() int {
    res := (*q)[0]
    *q = (*q)[1:]
    return res
}

func (q queue) Size() int {
    return len(q)
}


type MyStack struct {
    queues []queue
    queueSelector int    
}


func Constructor() MyStack {
    queues := make([]queue, 2)
    queues[0] = ConstructQueue()
    queues[1] = ConstructQueue()
    return MyStack{
        queues: queues,
        queueSelector: 0,
    }
}


func (this *MyStack) Push(x int)  {
    this.queues[this.queueSelector].Push(x)
}


func (this *MyStack) Pop() int {
    newQueueSelector := (this.queueSelector+1)%2
    oldQueue := this.queues[this.queueSelector]
    newQueue := this.queues[newQueueSelector]

    for oldQueue.Size() > 1 {
        newQueue.Push(oldQueue.Pop())
    }

    res := oldQueue.Pop()

    this.queues[this.queueSelector] = oldQueue
    this.queues[newQueueSelector] = newQueue
    this.queueSelector = newQueueSelector

    return res
}


func (this *MyStack) Top() int {
    newQueueSelector := (this.queueSelector+1)%2
    oldQueue := this.queues[this.queueSelector]
    newQueue := this.queues[newQueueSelector]

    var top int
    for oldQueue.Size() > 0 {
        top = oldQueue.Pop()
        newQueue.Push(top)
    }

    this.queues[this.queueSelector] = oldQueue
    this.queues[newQueueSelector] = newQueue
    this.queueSelector = newQueueSelector

    return top
}


func (this *MyStack) Empty() bool {
    return this.queues[0].Size() == 0 && this.queues[1].Size() == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */