var queue_origin
var front = 0;
var back = 0;

const queueInit = (n) => {
  queue_origin = new Array(n);
  front = 0;
  back = 0;
  console.log("queue_origin length: ", queue_origin.length);
}

const enqueue = (inputNumbers) => {
  // if (queue_origin[front] !== undefined) {
  //     console.log(`queue_origin is full`);
  // } else {
  //   if (front > 7) {
  //     front = 0;
  //     queue_origin[front] = inputNumbers;
  //     front = front + 1;
  //   } else {
  //     queue_origin[front] = inputNumbers;
  //     front = front + 1;
  //   }
  // }
  if (queue_origin[front] !== undefined) {
    console.log('queue_origin is full');
    return;
  }
  queue_origin[front] = inputNumbers;
  front += 1;
  if (front === queue_origin.length) {
    front = 0;
  }
  printQueue();
}

const dequeue = () => {
  // queue_origin[back] = undefined;
  // back = back + 1;
  if (queue_origin[back] === undefined) {
    console.log('queue_origin is emtpy');
    return;
  }
  const dequeueData = queue_origin[back];
  queue_origin[back] = undefined;
  back += 1;
  if (back === queue_origin.length) {
    back = 0;
  }
  console.log('dequeue data :', dequeueData);
  printQueue();
}

const printQueue = () => {
  // console.log("Print All")
  // for (i=0; i<10; i++) {
  //   console.log(queue_origin[i])
  // }
  // let queueDataStrings = '';
  // for (i=0; i < queue_origin.length; i++) {
  //   queueDataStrings = queueDataStrings.concat(queue_origin[i], ' ')
  // }
  console.log('queue_origin: ', queue_origin);
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
