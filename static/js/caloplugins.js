function pluginLogo(el) {
    el.innerHTML = "Hello Calo"
    this.model.user = 'me'
}

function pluginTree(container) {
    var root = { id: 0, pid: null, name: 'root' }
    var nodes = [root, { id: 6, pid: null, name: 'root1' }]
    for (let index = 1; index < 5; index++) {
        nodes.push({
            id: index,
            pid: index - 1,
            name: `node${index}`
        })

    }
    nodes.push({ id: 5, pid: 2, name: 'node5' })
    console.log(nodes);
    function doForCalo(node, nodes) {

    }
    let obj = doForCalo({ id: null }, nodes)
    container.innerHTML = obj

    let els = container.getElementsByTagName('li');
    console.log(els);
    for (const el in els) {
        els[el].onclick = function (e) {
            e.preventDefault()
            e.stopPropagation()
            console.log(this.id);
            if (this.children[0])
                this.children[0].style.display = this.children[0].style.display == 'none' ? '' : 'none';
        }
    }
    container.getElementsByTagName('ul')[0].style.display = ''

}