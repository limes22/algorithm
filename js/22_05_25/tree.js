let nodePool = {
  default: {
    parents: [],
    childs: []
  }
};
// 전개 연산자, 배열 append 방법, object 추가하는 방법
const addNode = (parentName, nodeName) => {
  // 본인 노드를 nodePool에 추가
  nodePool[nodeName] = {
    parents: [parentName],
    childs: [],
  };
  // 부모 node에 자신의 node정보를 childs로 등록
  nodePool[parentName].childs = [...nodePool[parentName].childs, nodeName];
};
// 재귀 함수(응용)로 해보기, for 문 반복으로
const printRelativeNode = (nodeName) => {
  console.log({nodeName, ...nodePool[nodeName]})
  if(nodePool[nodeName].childs.length) {
    for(let i=0; i<nodePool[nodeName].childs.length; i++){
      printRelativeNode(nodePool[nodeName].childs[i]);
    }
  }
}
addNode("default", "node1");
addNode("default", "node2");
addNode("node1", "node1Child1");
addNode("node1Child1", "node1ch")
addNode("node1", "node1Child2");
addNode("node2", "node2Child1")
// addNode("node1Child1", "node1ChildChild1");
console.log(printRelativeNode('default'))
