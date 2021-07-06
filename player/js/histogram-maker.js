var x1 = [];
var x2 = [];
for (var i = 0; i < 500; i ++) {
	x1[i] = Math.random();
	x2[i] = Math.random();
}

var trace1 = {
  name: 'your average saving',
  x: x1,
  type: "histogram",
  opacity: 0.5,
  marker: {
    color: "rgba(255, 100, 102, 0.7)", 
     line: {
      color:  "rgba(255, 100, 102, 1)", 
      width: 1
    }
  },  
};
var trace2 = {
  name: 'your friends\' average saving',
  x: x2,
  type: "histogram",
  opacity: 0.8,
  marker: {
    color: "rgba(100, 200, 102, 0.7)",
     line: {
      color:  "rgba(100, 200, 102, 1)", 
      width: 1
} 
 }, 
};
var data = [trace1, trace2];
var layout = {
  bargap: 0.1,
  barmode: "overlay",
  xaxis: {title: "Month"}, 
  yaxis: {title: "Energy"}
};
Plotly.newPlot('myDiv', data, layout);
