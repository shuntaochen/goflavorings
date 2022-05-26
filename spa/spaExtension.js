(function (o) {

    const root = o.rootel
    o.templateStore = {}
    o.router = [
        { '/routea': 'templateA.html' }
    ]
    for (const key in router) {
        if (Object.hasOwnProperty.call(router, key)) {
            const htmlName = router[key];
            templateStore[key] = '';

        }
    }
    o.navigate = function (route) {
        window.history.pushState(null, null, route)
        root.innerHTML = templateStore[route]
        o.run.apply(o)
    }
    window.calo = o || {
        model: {}
    }

})(window.calo)