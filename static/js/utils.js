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


function getQueryJson(route) {
    let arr = route.split("?")[1].split("&");
    let queryJson = {};
    for (let i of arr) {
        queryJson[i.split("=")[0]] = i.split("=")[1];
    }
    return queryJson
}

function whenready(onload) {
    document.body.onload = onload
}

//intended as a loaded scripts manager to avoid additional load of same file
window.loadedScripts = [];


function require(file, callback) {
    if (window.loadedScripts.indexOf(file) === -1) {
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
                    window.loadedScripts.push(file)
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
}

function makeCalo(o) {
    window.calo = { ...window.calo, ...o }
    calo.run.apply(calo)
    return window.calo
}


function requireAll(scripts, next) {
    let promises = [];
    scripts.filter(s => window.loadedScripts.indexOf(s) === -1).forEach(function (url) {
        var loader = new Promise(function (resolve, reject) {
            let script = document.createElement('script');
            script.src = url;
            script.async = false;
            script.onload = function () {
                window.loadedScripts.push(url)
                resolve(url);
            };
            script.onerror = function () {
                reject(url);
            };
            document.body.appendChild(script);
        });
        promises.push(loader);
    });

    return Promise.all(promises)
        .then(function () {
            console.log('all scripts loaded');
            next()
        }).catch(function (script) {
            console.log(script + ' failed to load');
        });
}

function dragger() {
}

dragger.prototype = {
    setSrc: function (el, clone) {
        this.src = el
        this.clone = clone
        el.draggable = true
        el.addEventListener('dragstart', function (e) {
            setTimeout(() => {
                console.log(this);
            }, 10)
        });

        el.addEventListener('dragend', function () {
        });
        return this
    }, setTargets: function () {
        let $g = this
        for (var i of arguments) {
            i.addEventListener('dragenter', dragEnter);
            i.addEventListener('dragover', dragOver);
            i.addEventListener('dragleave', dragLeave);
            i.addEventListener('drop', drop);
        }

        function dragEnter() {
            this.className += " holding";
        }

        function dragOver(e) {
            e.preventDefault();
        }

        function dragLeave() {
            this.className = 'empty';
        }

        function drop() {
            console.log('drop')
            this.className = 'empty';
            let el = $g.clone ? $g.src.cloneNode(true) : $g.src
            this.append(el)
        }
    },
}
