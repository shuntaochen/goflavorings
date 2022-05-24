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
function renderChildren(node, nodes) {
    var csh = ''
    nodes.filter(n => { return n.pid == node.id }).forEach(
        c => {
            csh += `<li>${c.name}${renderChildren(c, nodes)}</li>`
        }
    )
    return csh == '' ? '' : `<ul>${csh}</ul>`
}
let tree = renderChildren({ id: null }, nodes)
const p1 = document.createElement('div')
document.body.appendChild(p1)
p1.innerHTML = tree

let els = document.getElementsByTagName('li');
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
document.getElementsByTagName('ul')[0].style.display = ''