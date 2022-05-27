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

function getHtmlOrJson(url, success) {
    function makeHttpObject() {
        if ("XMLHttpRequest" in window) return new XMLHttpRequest();
        else if ("ActiveXObject" in window) return new ActiveXObject("Msxml2.XMLHTTP");
    }

    var request = makeHttpObject();
    request.open("GET", url + "?" + new Date().getMilliseconds(), true);
    request.send(null);
    request.onreadystatechange = function () {
        if (request.readyState == 4) {
            console.log(request.responseText);
            success(request.responseText)
        }
    };
}

function whenready(onload) {
    document.body.onload = onload
}


function require(file, callback) {
    // create script element
    var script = document.createElement("script");
    script.src = file;
    // monitor script loading
    // IE < 7, does not support onload
    if (callback) {
        script.onreadystatechange = function () {
            if (script.readyState === "loaded" || script.readyState === "complete") {
                // no need to be notified again
                script.onreadystatechange = null;
                // notify user
                callback();
            }
        };
        // other browsers
        script.onload = function () {
            callback();
        };
    }
    // append and execute script
    document.documentElement.firstChild.appendChild(script);
}

function makeCalo(o) {
    window.calo = { ...window.calo, ...o }
    calo.run.apply(calo)
    return window.calo
}