function request() {
    text = document.getElementById("query").value;
    request = new XMLHttpRequest();
    request.open("POST", "/request", true);
    request.send(text);
    alert('Sent request for ' + text);
}