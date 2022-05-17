var queue
var front = 0;
var back = 0;

const queueInit = (n) => {
 queue = new Array(n);
 front = 0;
 back = 0;
 console.log('length:', queue.length)
}

const enqueue = (inputNumber, priority) => {
  if (queue[front] !== undefined) {
    console.log('queue is full')
    return;
  }
  queue[front] = {inputNumber, priority};
  const enqueueData = queue[front]
  front += 1;
  if (front === queue.length){
    front = 0;
  }
  console.log('enqueue data: ', enqueueData);
}

const dequeue = () => {
  if (queue[back] === undefined){
    console.log('queue is empty')
    return;
  }
  const dequeueData = queue[back]
  queue[back] = undefined;
  back += 1
  if (back === queue.length) {
    back = 0;
  }
  console.log('dequeue data: ', dequeueData);
  
}


const showAll = () => { 
    for (i =0; i < queue.length; i ++) {
      console.log('queue: ', queue[i])
    } 
}

const bubbleSort = () => {
  for (i = 0; i < queue.length; i++) {
    for (j = i+1; j < queue.length; j++) {
      if (queue[i]['priority'] > queue[j]['priority']) {
        const temp = queue[i];
        queue[i] = queue[j];
        queue[j] = temp;
      }
    }
  }
}


queueInit(4)
enqueue(6,5);
enqueue(5, 4);
enqueue(4, 3);
enqueue(3, 2);
dequeue();
enqueue(2, 1);
bubbleSort ();
showAll();
dequeue();
dequeue();
dequeue();
dequeue();
dequeue();
showAll();
