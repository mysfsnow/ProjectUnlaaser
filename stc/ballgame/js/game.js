var KEYS = {
	UP: 38,
	DOWN: 40,
	W: 87,
	S: 83
}

var keyManager = {};
keyManager.pressedKeys = [];

var pingpong = {};
pingpong.ball = {
	speed: 5,
	x: 150,
	y: 100,
	directionX: 1,
	directionY: 1,
	diameter: 20,
	id: ""
}

$(function(){
	// 设置interval用于每隔30s调用一次gameloop
	keyManager.timer = setInterval(gameloop,30);

	// 标记pressedKeys数组里面某键的状态
	$(document).keydown(function(e){
		keyManager.pressedKeys[e.keyCode] = true;
	});

	$(document).keyup(function(e){
		keyManager.pressedKeys[e.keyCode] = false;
	});
});

function gameloop(){
	moveBall();
	movePaddles();
}

function moveBall(){
	var playGroundHeight = parseInt($("#playground").height());
	var playGroundWidth = parseInt($("#playground").width());
	var ball = pingpong.ball;

	if (ball.y + ball.speed * ball.directionY > playGroundHeight - ball.diameter) {
		ball.directionY = -1;
	};

	if (ball.y + ball.speed * ball.directionY < 0) {
		ball.directionY = 1;
	};

	if (ball.x + ball.speed * ball.directionX > playGroundWidth - ball.diameter) {
		ball.directionX = -1;
	};

	if (ball.x + ball.speed * ball.directionX < 0) {
		ball.directionX = 1;
	};

	ball.x += ball.speed * ball.directionX;
	ball.y += ball.speed * ball.directionY;

	// 检测左球拍
	var paddleAX = parseInt($('#paddleA').css("left")) + parseInt($('#paddleA').css("width"));
	var paddleAYBottom = parseInt($('#paddleA').css("top")) + parseInt($('#paddleA').css("height"));
	var paddleAYTop = parseInt($('#paddleA').css("top"));

	if (ball.x + ball.speed * ball.directionX < paddleAX) {
		if (ball.y + ball.speed * ball.directionY + ball.diameter/2.0 <= paddleAYBottom 
			&& ball.y + ball.speed * ball.directionY >= paddleAYTop) {
			ball.directionX = 1;
		};
	};

	// 检测右球拍
	var paddleBX = parseInt($('#paddleB').css("left"));
	var paddleBYBottom = parseInt($('#paddleB').css("top")) + parseInt($('#paddleB').css("height"));
	var paddleBYTop = parseInt($('#paddleB').css("top"));

	if (ball.x + ball.speed * ball.directionX + ball.diameter >= paddleBX) {
		if (ball.y + ball.speed * ball.directionY + ball.diameter/2.0 <= paddleBYBottom 
			&& ball.y + ball.speed * ball.directionY >= paddleBYTop) {
			ball.directionX = -1;
		};
	};

	$("#ball").css({
		"left": ball.x,
		"top": ball.y
	})
}

function movePaddles(){
	if (keyManager.pressedKeys[KEYS.UP]) {
		var top = parseInt($("#paddleB").css("top"));
		$("#paddleB").css("top", top - 5);
	};

	if (keyManager.pressedKeys[KEYS.DOWN]) {
		var top = parseInt($("#paddleB").css("top"));
		$("#paddleB").css("top", top + 5);
	};

	if (keyManager.pressedKeys[KEYS.W]) {
		var top = parseInt($("#paddleA").css("top"));
		$("#paddleA").css("top", top - 5);
	};

	if (keyManager.pressedKeys[KEYS.S]) {
		var top = parseInt($("#paddleA").css("top"));
		$("#paddleA").css("top", top + 5);
	};
}