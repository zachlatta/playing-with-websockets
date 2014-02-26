var canvas = document.getElementById('canvas');
var circle, ctx;

if (canvas.getContext) {
  var ctx = canvas.getContext('2d');
  circle = {
    x: 100,
    y: 100,
    radius: 50,
    draw: function (ctx) {
      ctx.strokeStyle = '#000';
      ctx.fillStyle = '#ffff00';
      ctx.beginPath();
      ctx.arc(this.x,this.y,this.radius,0,Math.PI*2,true);
      ctx.closePath();
      ctx.stroke();
      ctx.fill();
    }
  }
}

function update (delta) {
}

function render () {
  ctx.save();
  ctx.setTransform(1, 0, 0, 1, 0, 0);
  ctx.clearRect(0, 0, canvas.width, canvas.height);
  ctx.restore();

  circle.draw(ctx);
}

function main () {
  var now = Date.now();
  var delta = now - then;

  update(delta / 1000);
  render();

  then = now;
}

var then = Date.now();
setInterval(main, 1);
