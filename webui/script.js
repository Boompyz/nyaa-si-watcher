function request() {
    text = document.getElementById("query").value;
    xmlrequest = new XMLHttpRequest();
    xmlrequest.open("POST", "/request", true);
    xmlrequest.send(text);
    alert('Sent request for ' + text);
}
function add(what, name) {
    xmlrequest = new XMLHttpRequest();
    xmlrequest.onreadystatechange = function () {
        if (xmlrequest.readyState === 4 && xmlrequest.status === 200) {
            location.reload();
        } else if (xmlrequest.readyState == 4) {
            alert("error")
        }
    }
    xmlrequest.open("POST", "/add" + what, true);
    xmlrequest.send(name);
}

function remove(what, name) {
    xmlrequest = new XMLHttpRequest();
    xmlrequest.onreadystatechange = function () {
        if (xmlrequest.readyState == 4 && xmlrequest.status == 200) {
            location.reload();
        } else if (xmlrequest.readyState == 4) {
            alert("error")
        }
    }
    xmlrequest.open("POST", "/remove" + what, true);
    xmlrequest.send(name);
}
