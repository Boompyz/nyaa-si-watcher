function request() {
    text = document.getElementById("query").value;
    xmlrequest = new XMLHttpRequest();
    xmlrequest.open("POST", "/request", true);
    xmlrequest.send(text);
    alert('Sent request for ' + text);
}
function addWatch() {
    watch = {}
    watch.query = document.getElementById('query').value
    watch.folder = document.getElementById('folder').value
    watch.user = document.getElementById('user').value
    
    xmlrequest = new XMLHttpRequest();
    xmlrequest.onreadystatechange = function () {
        if (xmlrequest.readyState === 4 && xmlrequest.status === 200) {
            location.reload();
        } else if (xmlrequest.readyState == 4) {
            alert("error")
        }
    }
    xmlrequest.open("POST", "/addwatch", true);
    xmlrequest.send(JSON.stringify(watch));
}

function removeWatch(query, user, folder) {
    watch = {}
    watch.query = query
    watch.folder = user
    watch.user = folder

    xmlrequest = new XMLHttpRequest();
    xmlrequest.onreadystatechange = function () {
        if (xmlrequest.readyState == 4 && xmlrequest.status == 200) {
            location.reload();
        } else if (xmlrequest.readyState == 4) {
            alert("error")
        }
    }
    xmlrequest.open("POST", "/removewatch", true);
    xmlrequest.send(JSON.stringify(watch));
}
