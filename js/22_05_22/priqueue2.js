let queue;
const queueInit = (n) => {
  queue = new Array(n)
}

const enqueue = (inputNumber, priority = -1) => {
  for(let i=0; i<queue.length; i++) {
    if(!queue[i]) {
      queue[i] = {inputNumber, priority}
      const enqueue = queue[i]
      console.log(`enqueue:`, enqueue)
      return;
    }
  }
  console.log("queue is full")
}

const dequeue = () => {
  let priority;
  let priIdx;
  for(let i=0; i<queue.length; i++) {
    if (!queue[i]) {
      continue;
    }
    if (!priority) {
      priority = queue[i].priority
      priIdx = i;
      continue;
    }
    if (queue[i].priority <= priority) {
      priority = queue[i].priority;
      priIdx = i;
    }
    if (queue[priIdx] === undefined) {
      console.log("queue is empty");
      return;
    }
    const dequeue = queue[priIdx]
    console.log(`dequue: `,dequeue)
    queue[priIdx] = undefined;
  }
}

const showAll = () => {
  for(let i=0; i<queue.length; i++) {
    console.log(`queue : `, queue[i])
  }
}

queueInit(3)
enqueue(4,2)
enqueue(5,3)
enqueue(6,4)
showAll()
dequeue()
dequeue()
showAll()
enqueue(7,3)
enqueue(8,2)
dequeue()
dequeue()
dequeue()
dequeue()
showAll()
