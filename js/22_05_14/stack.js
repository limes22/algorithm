//stack은 FILO(First In Last Out)  마지막에 push(in) 한 데이터가 제일 먼저 pop(out) 됨 
//데이터 공간블럭의 움직임을 최소화함.
var stack //data structure array
var cursor //data index cursor

function stackInit(n) {
  stack = new Array(n);
  cursor = 0;
}
const push = (inputNumber) => {
  stack[cursor] = inputNumber
  cursor = cursor + 1
}
const pop = () => {
  cursor = cursor - 1
  stack[cursor] = undefined
}
const printAll = () => {
  console.log("Print Stack Values")
  for (i=0; i < 3; i++) {
    console.log(stack[i])
  }
}
// Run Test
stackInit(3)
push(1)
push(2)
push(3)
printAll()
