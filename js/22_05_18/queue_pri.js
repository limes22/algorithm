let queue;
let front = 0;
let back = 0;

const queueInit = (n) => {
  queue = new Array(n);
  front = 0;
  back = 0;
  console.log("length:", queue.length);
};

const enqueue = (inputNumber, priority = -1) => {
  if (queue[front] !== undefined) {
    console.log("queue is full");
    return;
  }
  queue[front] = { inputNumber, priority };
  bubbleSort();
  const enqueueData = queue[front];
  front += 1;
  if (front === queue.length) {
    front = 0;
  }
  // if (queue.priority === 'unknown') {
  // queue[front] = queue.splice(front +1, 1, {inputNumber, priority}) }
  console.log("enqueue data: ", enqueueData);
};

const dequeue = () => {
  if (queue[back] === undefined) {
    console.log("queue is empty");
    return;
  }
  const dequeueData = queue[back];
  queue[back] = undefined;
  back += 1;
  if (back === queue.length) {
    back = 0;
  }
  console.log("dequeue data: ", dequeueData);
};

const showAll = () => {
  for (let i = 0; i < queue.length; i++) {
    console.log("queue", [i], " : ", queue[i]);
  }
};

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
};

queueInit(3);
enqueue(1, 5);
enqueue(1, 3);
enqueue(1, 1);
dequeue();
dequeue();
enqueue(1, 1);
dequeue();
showAll();
