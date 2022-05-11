var showNum = function(){
  var num1 = document.getElementsByTagName("input")[0].value;
  var result = parseInt(num1)
  document.getElementsByTagName("input")[0].value=result;

  var n = result;
  document.write('★'.repeat(n)+'<br>')

  for(let j=1; j<=n; j=j+2) {
    document.write('★'.repeat((-j+n)*0.5)+'○'.repeat(j)+'★'.repeat((-j+n)*0.5)+'<br>')
  }
  
  for(let j=n-2; j>=1; j=j-2) {
    document.write('★'.repeat((-j+n)*0.5)+'○'.repeat(j)+'★'.repeat((-j+n)*0.5)+'<br>')
  }

  document.write('★'.repeat(n)+'<br>')

}

// for(let j=3; j>=1; j=j-2) {
//   document.write('★'.repeat(-0.5*j+5.5)+'○'.repeat(j)+'★'.repeat(-0.5*j+5.5)+'<br>')
// }

// for(let j=1; j<=5; j=j+2) {
//   document.write('○'.repeat(j)+'<br>')
// }

// for(let i=4; i<=5; i++){
//   document.write('★'.repeat(i)+'<br>')
// }

// for(let j=3; j>=1; j=j-2){
//   document.write('○'.repeat(j)+'<br>')
// }
