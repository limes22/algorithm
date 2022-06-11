let stack;
let cursor;

const temp = () => {
  for(let i =0; i < 10; ++i) {

  }
}

const stackInit = (n) => {
  stack = new Array(n);
  cursor = 0;
}

const push = (inputNumber) => {
  if (cursor === stack.length) {
    console.log('stack is full');
    return;
  }
  stack[cursor] = inputNumber;
  cursor += 1;
  console.log(inputNumber)
}

const pop = () => {
  if (cursor === 0) {
    console.log('queue is empty');
    return;
  }
  console.log('popData:', stack[cursor-1]);
  stack[cursor] = undefined;
  cursor -= 1;
}

stackInit(3)
push(5)
push(7)
push(9)
push(10)
console.log(stack)
