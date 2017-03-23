
$(function () {

// var log = $('#log');
console.log("asda");
	function connection() {
	var conn;
    var msg = document.getElementById("msg");
    var log = document.getElementById("log");
    function appendLog(item) {
        var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
        log.appendChild(item);
        if (doScroll) {
            log.scrollTop = log.scrollHeight - log.clientHeight;
        }
    }

    document.getElementById("form").onsubmit = function () {
        if (!conn) {
            return false;
        }
        if (!msg.value) {
            return false;
        }
        conn.send(msg.value);
        msg.value = "";
        return false;
    };

    document.getElementById("clear").onsubmit = function () {
        if (!conn) {
            return false;
        }
        console.log("test");
        wtf = "privet";
        conn.send(wtf);
        return false;
    };

    if (window["WebSocket"]) {
        conn = new WebSocket("ws://localhost:47000/ws");
        conn.onopen = function (evt) {
            console.log("Connection to websocket.");
        };
        conn.onclose = function (evt) {
            var item = document.createElement("div");
            item.innerHTML = "<b>Connection closed.</b>";
            appendLog(item);
        };
        conn.onmessage = function (evt) {
            console.log(evt);
            var messages = evt.data.split('\n');
            for (var i = 0; i < messages.length; i++) {
                var item = document.createElement("div");
                item.innerText = messages[i];
                appendLog(item);
            }
        };
    } else {
        var item = document.createElement("div");
        item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
        appendLog(item);
    }
	}
	connection();


});
