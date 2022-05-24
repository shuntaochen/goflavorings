/*

License: The Calo Open Source License
Author: Shuntao Chen,
Link: https://www.caloch.cn
Privilege is hereby granted, free of use, this is an MVVM framework that I created, it is light enough to help a lot without any reliance on the other libraries, but please do keep all these info integrate when you use it! Any disobey will cause certain law issues, 

**/

(function () {
    function doForCalo(data, scope, prefix, jsonPathPrefix) {
        for (const key in data) {
            if (Object.hasOwnProperty.call(data, key)) {
                const fieldValue = data[key];
                if (isValType(fieldValue)) {
                    const els = getElsByFieldName(scope, prefix + key)
                    els.forEach(el => {
                        SetValue(el, data[key])
                        el.dataset.jsonPath = jsonPathPrefix + "." + key
                    });
                }
                else if (isObjectType(fieldValue)) {
                    const els = getElsByFieldName(scope, key)
                    els.forEach(el => {
                        el.dataset.jsonPath = jsonPathPrefix + "." + key
                        doForCalo(fieldValue, el, "", el.dataset.jsonPath)
                    })
                }
                else if (isArrayType(fieldValue)) {
                    const els = scope.querySelectorAll("[\\@model^=" + key + "\\|]")
                    els.forEach(el => {
                        const fdPrefix = el.getAttribute("@model").split('|')[1]
                        let ci = 0
                        let lastCursor = el
                        fieldValue.forEach(val => {
                            if (ci === 0) {
                                el.dataset.jsonPath = jsonPathPrefix + "." + key + `[${ci}]`
                                if (isValType(val))
                                    SetValue(el, val)
                                else
                                    doForCalo(val, el, fdPrefix + ".", el.dataset.jsonPath)
                            }
                            else {
                                clone = el.cloneNode(true)
                                clone.setAttribute("poped", "true")
                                clone.dataset.jsonPath = jsonPathPrefix + "." + key + `[${ci}]`
                                insertAfter(clone, lastCursor)
                                lastCursor = clone
                                if (isValType(val))
                                    SetValue(clone, val)
                                else
                                    doForCalo(val, clone, fdPrefix + ".", clone.dataset.jsonPath)
                            }
                            ci++
                        });
                    })
                }
            }
        }

    }

    window.calo = window.calo || {
        model: {}
    }
    window.log = x => console.log(x)
    function removePopped(root) {
        var poped = root.querySelectorAll("[poped='true']")
        poped.forEach(p => {
            p.parentNode.removeChild(p)
        })
    }
    var root = document.querySelector("[calo]");
    calo.run = function () {
        removePopped(root)
        doForCalo(calo.model, root, "", "calo.model")
        var clicks = root.querySelectorAll("[\\@Click]")
        var changes = root.querySelectorAll("[\\@Change]")
        clicks.forEach(c => {
            c.onclick = function () {
                calo[c.getAttribute("@Click")].call(calo, c, c.value)
                calo.run.apply(calo)
            }
        })
        changes.forEach(c => {
            c.onchange = function () {
                calo[c.getAttribute("@Change")].call(calo, c, c.value)
                calo.run.apply(calo)
            }
        })

        root.querySelectorAll("[\\@Show]").forEach(el => {
            const flag = getDataValByJsonPath(el)
            el.style.display = flag.toLowerCase().toString() == "true" ? '' : 'none'
        })

        root.querySelectorAll("input").forEach(ip => {
            if (window.addEventListener) {
                ip.addEventListener('keyup', function () {
                    event.preventDefault()
                    event.stopPropagation()
                    eval(ip.dataset.jsonPath + "='" + ip.value + "'")
                    root.querySelectorAll(`[data - json - Path= '${ip.dataset.jsonPath}']`).forEach(el => {
                        SetValue(el, ip.value)
                    })

                }, false);
            } else {
                ip.attachEvent('change', function () { log(5); });
            }
        })

    }

    calo.run.apply(calo)


    function SetValue(el, val) {
        if (el.tagName === "INPUT") el.value = val
        if (["LABEL", "BUTTON", "SPAN"].indexOf(el.tagName) != -1) el.innerHTML = val

    }
    function getElsByFieldName(scope, fieldName) {
        return scope.querySelectorAll("[\\@model='" + fieldName + "'")
    }
    function isValType(obj) {
        return typeof obj === "number" || typeof obj === "string" || typeof obj === "boolean"
    }
    function isArrayType(obj) {
        return obj instanceof Array
    }
    function isObjectType(obj) {
        return !(obj instanceof Array) && typeof obj === "object"
    }
    function getDataValByJsonPath(el) {
        let jsonPath = el.dataset.jsonPath
        return eval(jsonPath)
    }

    calo.getElsByJsonPath = getElsByJsonPath
    function getElsByJsonPath(jsonPath) {
        return root.querySelectorAll(`[data - json - Path= '${jsonPath}']`)
    }

    calo.$ = $
    function $(id) {
        return document.getElementById(id)
    }
    window.calo.makePlugin = function (id, functionPlugin) {
        functionPlugin.call(calo, $(id))
        calo.run.apply(calo)
    }

    function insertAfter(newElement, targetElement) {
        var parent = targetElement.parentNode;
        if (parent.lastChild == targetElement) {
            parent.appendChild(newElement);
        }
        else {
            parent.insertBefore(newElement, targetElement.nextSibling);
        }
    }
    calo.ajax = ajax
    calo.ajaxQueue = ajaxQueue
    function ajaxQueue({ url, data, type, success, error }) {
        if (ajaxQueue.queue) {
            ajaxQueue.queue.push({ url, data, type, success, error })
            while (ajaxQueue.queue.length > 0) {
                if (!ajaxQueue.queue.pending) {
                    ajaxQueue.queue.pending = true
                    const request = ajaxQueue.queue[0]
                    ajax({
                        url: request.url, data: request.data, type: request.type, success: function (resp) {
                            request.success(resp)
                            ajaxQueue.queue.shift()
                            ajaxQueue.queue.pending = false
                        }, error: function () {
                            request.error()
                            ajaxQueue.queue.pending = false
                        }
                    })
                } else break

            }
        }
    }
    ajaxQueue.queue = []


    function ajax({ url, data, type, success, error }) {
        type = type || "get";
        data = data || {};
        let str = "";
        for (let i in data) {
            str += `${i}=${data[i]}& `;
        }
        str = str.slice(0, str.length - 1);
        if (type === "get") {
            var d = new Date();
            url = url + "?" + str + "&__qft=" + d.getTime();
        }
        let xhr = new XMLHttpRequest();
        xhr.open(type, url, true);
        if (type === "get") {
            xhr.send();
        } else if (type === "post") {
            xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
            xhr.send(str);
        }
        xhr.onload = function () {
            if (xhr.status === 200) {
                success.call(calo, JSON.parse(xhr.responseText))
                calo.run.apply(calo);
            } else {
                error && error(xhr.status);
            }
        }
    }
})()


