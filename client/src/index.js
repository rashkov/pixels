import React from "react";
import ReactDOM from "react-dom";
import App from "./App";
import "./index.css";

// @TODO: Handle auto-reconnection using onClose event

(() => {
  var ws = new WebSocket(`ws://localhost:8080/ws`);

  ws.onopen = function() {
    ws.send(JSON.stringify({ message: "hello server!" }));
  };

  ws.onmessage = function(event) {
    var m = JSON.parse(event.data);
    console.debug("Received message", m.message);
  };

  ws.onerror = function(event) {
    console.debug(event);
  };
})();

ReactDOM.render(
  <App width="512" height="512" />,
  document.getElementById("root")
);
