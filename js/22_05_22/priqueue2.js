let queue;

const queueInit = (n) => {
  queue = new Array(n);
  back = 0;
}

const enqueue = (inputNumber, priority = -1) => {
  for (let i=0; i<queue.length; i++) {
    if (!queue[i]) {
    queue[i] = {inputNumber, priority}
    const enqueuedata = queue[i]
    console.log(`enqueuedata:`, enqueuedata)
    return;
    }
  }
  console.log('queue is full')
}

const dequeue = () => {
  let priority;
  let priorityIdx;
  for (let i = 0; i < queue.length; i++) {
    if (!queue[i]) {
      continue;
    }
    if (!priority) {
      priority = queue[i].priority;
      priorityIdx = i;
      console.log(`priorityIdx(1)`, priorityIdx)
      continue;
    }
    if (queue[i].priority <= priority) {
      priority = queue[i].priority;
      priorityIdx = i;
      console.log(`priorityIdx(2)`, priorityIdx)
    }
  }
  console.log(`priorityIdx(3)`, priorityIdx)
  const dequeue = queue[priorityIdx]
  console.log(`dequue: `,dequeue)
    if (queue[priorityIdx] === undefined) {
    console.log(`queue is empty`)
  }
  queue[priorityIdx] = undefined;
};



const showAll = () => {
  for (let i=0; i<queue.length; i++) {
    console.log(`queue : `, queue[i])
  }
}



queueInit(3)
enqueue(5, 2)
enqueue(6, 3)
enqueue(8, 1)
dequeue()
dequeue()
dequeue()
showAll()
