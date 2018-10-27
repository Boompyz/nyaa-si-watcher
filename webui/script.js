function request() {
    text = document.getElementById("query").value;
    xmlrequest = new XMLHttpRequest();
    xmlrequest.open("POST", "/request", true);
    xmlrequest.send(text);
    alert('Sent request for ' + text);
}