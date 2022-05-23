let queue;
let back = 0;

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
  var priority;
  var priorityIdx;
  for (let i = 0; i < queue.length; i++) {
    if (!queue[i]) {
      continue;
    }
    if (!priority) {
      priority = queue[i].priority;
      priorityIdx = i;
      continue;
    }
    if (queue[i].priority <= priority) {
      priority = queue[i].priority;
      priorityIdx = i;
    }
  }
  const dequeue = queue[priorityIdx]
  console.log(`dequue: `,dequeue)
  queue[priorityIdx] = undefined;
  if (!priorityIdx) {
    console.log("queue is empty")
    return;
  }
};



const showAll = () => {
  for (let i=0; i<queue.length; i++) {
    console.log(`queue : `, queue[i])
  }
}

const bubbleSort = () => {
  for (let i = 0; i < queue.length; i++) {
    for (let j = i + 1; j < queue.length; j++) {
      // queue[i]
      if (!queue[i] || !queue[j]) {
        continue;
      }
      if (queue[i].priority > queue[j].priority) {
        const temp = queue[i];
        queue[i] = queue[j];
        queue[j] = temp;
      }
    }
  }
}


queueInit(3)
enqueue(5,4)
enqueue(6,5)
enqueue(7,3)
dequeue()
dequeue()
dequeue()
dequeue()
dequeue()
showAll()
