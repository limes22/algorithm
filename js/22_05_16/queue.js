var queue
var front
var back

const queueInit = (n) => {
  queue = new Array(n);
  front = 0;
  back = 0;
}

const enqueue = (inputNumbers) => {
  if (front > 7) {
    front = 0; 
  
  // } else if (queue[back] !== undefined) {
  //   console.log(`queue is full`)
  } else { 
    queue[front] = inputNumbers
    front = front + 1
  }
}

const dequeue = () => {
  queue[back] = undefined
  back = back + 1
}

const printAll = () => {
  console.log("Print All")
  for (i=0; i<5; i++) {
    console.log(queue[i])
  }
}

queueInit(8)
enqueue(3)
enqueue(5)
enqueue(7)
dequeue()
printAll()
