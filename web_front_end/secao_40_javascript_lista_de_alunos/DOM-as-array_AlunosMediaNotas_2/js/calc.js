function sum() {
    // const numbers = Array.from(arguments)
    const numbers = [...arguments]
    return numbers.reduce(function (sum, atual) {
        return sum + atual
    }, 0)
}
function avarege() {
    return sum(...arguments) / arguments.length
}

const alunos = [
        { nome: "Daniel", n1: 10, n2: 3, n3: 7.5, n4: 3 },
        { nome: "Maria", n1: 10, n2: 9, n3: 3, n4: 9.5 },
        { nome: "João", n1: 10, n2: 4.5, n3: 1, n4: 3.5 },
        { nome: "Joana", n1: 1, n2: 3, n3: 9, n4: 1.5 },
        { nome: "José", n1: 10, n2: 4.5, n3: 7, n4: 3 },
        { nome: "Arnaldo", n1: 10, n2: 4.5, n3: 7, n4: 3 },
        { nome: "Lucas", n1: 4.5, n2: 9, n3: 8, n4: 3 },
        { nome: "Luana", n1: 3, n2: 7, n3: 9, n4: 3 },
        { nome: "Beatriz", n1: 10, n2: 4, n3: 7, n4: 9 },
        { nome: "Sergio", n1: 4.5, n2: 9.5, n3: 10, n4: 2 }
]

function createElement(el){
    return document.createElement(el)
}

const tbody = document.querySelector('tbody')
const frag = document.createDocumentFragment()

alunos.forEach(aluno => {
    const tr = createElement('tr')
    const tdNome = createElement('td')
    const td1 = createElement('td')
    const td2 = createElement('td')
    const td3 = createElement('td')
    const td4 = createElement('td')
    const tdMedia = createElement('td')

    tdNome.textContent = aluno['nome']
    td1.textContent = aluno['n1']
    td2.textContent = aluno['n2']
    td3.textContent = aluno['n3']
    td4.textContent = aluno['n4']
    tdMedia.textContent = (avarege(aluno['n1'], aluno['n2'], aluno['n3'], aluno['n4'])).toFixed(2)

    tr.appendChild(tdNome)
    tr.appendChild(td1)
    tr.appendChild(td2)
    tr.appendChild(td3)
    tr.appendChild(td4)
    tr.appendChild(tdMedia)

    frag.appendChild(tr)
})

tbody.appendChild(frag)