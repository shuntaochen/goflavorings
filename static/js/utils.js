function loadScript(path) {
    var node = document.createElement('script')
    node.type = 'text/javascript'
    node.src = path + '?_=' + (new Date().getMilliseconds())
    document.body.appendChild(node)
}

function loadScripts() {
    for (let i = 0; i < arguments.length; i++) {
        const el = arguments[i];
        loadScript(el)
    }
}

function redirect(path) {
    location.href = path
}

function getHtmlOrJson(success) {
    function makeHttpObject() {
        if ("XMLHttpRequest" in window) return new XMLHttpRequest();
        else if ("ActiveXObject" in window) return new ActiveXObject("Msxml2.XMLHTTP");
    }

    var request = makeHttpObject();
    request.open("GET", "/", true);
    request.send(null);
    request.onreadystatechange = function () {
        if (request.readyState == 4)
            console.log(request.responseText);
        success(request.responseText)
    };
}
