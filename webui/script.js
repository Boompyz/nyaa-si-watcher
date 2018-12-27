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
    xmlrequest.onreadystatechange = function () {
        if (xmlrequest.readyState === 4 && xmlrequest.status === 200) {
            location.reload();
        }
    }
    xmlrequest.open("POST", "/addwatch", true);
    xmlrequest.send(text);
}