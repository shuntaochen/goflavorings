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