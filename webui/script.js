function request() {
    text = document.getElementById("query").value;
    xmlrequest = new XMLHttpRequest();
    xmlrequest.open("POST", "/request", true);
    xmlrequest.send(text);
    alert('Sent request for ' + text);
}
function addWatch() {
    text = document.getElementById("query").value;
    xmlrequest = new XMLHttpRequest();
    xmlrequest.open("POST", "/addwatch", true);
    xmlrequest.send(text);
    alert('Sent request for ' + text);
}