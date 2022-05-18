var stack // data structure array
var front = 0;
var back = 0; // data index cursor


const stackInit = (n) => {
  stack = new Array(n);
  front = 0;
  back = stack.length - 1;
  console.log(stack.length)
}

const push = (InputNumber) => {
  if (stack[front] !== undefined){
    console.log("stack is full")
    return;
  }
  stack[front] = InputNumber;
  front += 1;
  if (front === stack.length) {
    front = 0;
  }
}

const pop = () => {
  if (stack[back] === undefined){
    console.log("stack is empty")
    return;
  }
  stack[back] = undefined;
  back -= 1;
}

const printAll = () => {
  console.log(stack)
}

// Run Test
stackInit(3)
push(1)
push(2)
push(3)
push(5)
pop()
pop()
pop()
printAll()
