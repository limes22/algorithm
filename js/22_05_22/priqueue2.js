let queue;

const queueInit = (n) => {
  queue = new Array(n);
}

const enqueue = (inputNumber, priority = -1) => {
  for (let i = 0; i < queue.length; i++){
    if (!queue[i]){
      queue[i] = {inputNumber, priority};
      return;
  }
}
}

const dequeue = () => {
  let priority;
  let priIdx;
  for (let i = 0; i < queue.length; i++){
    if (!queue[i]){
      continue;
    }
    if (!priority) {
      priority = queue[i].priority
      priIdx = i
      continue;
    }
    if (priority > queue[i].priority){
      priority = queue[i].priority;
      priIdx = i
    }
  }
  if (!queue[priIdx])
  {
    console.log('queue is empty')
    return;
  }
  console.log('dequeue data:', queue[priIdx])
  queue[priIdx] = undefined
}

queueInit(3)
enqueue(6, 3)
enqueue(7, 2)
enqueue(9, 4)
dequeue()
dequeue()
dequeue()
dequeue()
console.log(queue)
