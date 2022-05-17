var queue
var front = 0;
var back = 0;

const queueInit = (n) => {
  queue = new Array(n);
  front = 0;
  back = 0;
  console.log("queue length: ", queue.length);
}

const enqueue = (inputNumbers) => {
  // if (queue[front] !== undefined) {
  //     console.log(`queue is full`);
  // } else {
  //   if (front > 7) {
  //     front = 0;
  //     queue[front] = inputNumbers;
  //     front = front + 1;
  //   } else {
  //     queue[front] = inputNumbers;
  //     front = front + 1;
  //   }
  // }
  if (queue[front] !== undefined) {
    console.log('queue is full');
    return;
  }
  queue[front] = inputNumbers;
  front += 1;
  if (front === queue.length) {
    front = 0;
  }
  printQueue();
}

const dequeue = () => {
  // queue[back] = undefined;
  // back = back + 1;
  if (queue[back] === undefined) {
    console.log('queue is emtpy');
    return;
  }
  const dequeueData = queue[back];
  queue[back] = undefined;
  back += 1;
  if (back === queue.length) {
    back = 0;
  }
  console.log('dequeue data :', dequeueData);
  printQueue();
}

const printQueue = () => {
  // console.log("Print All")
  // for (i=0; i<10; i++) {
  //   console.log(queue[i])
  // }
  // let queueDataStrings = '';
  // for (i=0; i < queue.length; i++) {
  //   queueDataStrings = queueDataStrings.concat(queue[i], ' ')
  // }
  console.log('queue: ', queue);
}

queueInit(4)
enqueue(1);
enqueue(2);
enqueue(3);
enqueue(4);
enqueue(5);
dequeue();
dequeue();
dequeue();
dequeue();
dequeue();
