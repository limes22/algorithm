var queue
var front = 0;
var back = 0;

const queueInit = (n) => {
  queue = new Array(n);
  front = 0;
  back = 0;
}

const enqueue = (inputNumbers) => {
  if (queue[front] !== undefined) {
      console.log(`queue is full`)
  } else {
    if (front > 7) {
      front = 0
      queue[front] = inputNumbers
      front = front + 1
    } else {
      queue[front] = inputNumbers
      front = front + 1
    }
}
}

const dequeue = () => {
  queue[back] = undefined
  back = back + 1
}

const printAll = () => {
  console.log("Print All")
  for (i=0; i<10; i++) {
    console.log(queue[i])
  }
}

queueInit(8)
enqueue(3)
enqueue(5)
enqueue(7)
enqueue(9)
enqueue(11)
enqueue(12)
enqueue(13)
enqueue(1)
dequeue()
enqueue(10)
enqueue(11)
printAll()
console.log(queue)
