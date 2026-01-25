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

const tr = document.querySelectorAll('tbody tr')


tr.forEach(elementTR => {
    const arrayNotes = []

    const td = elementTR.querySelectorAll('td')

    td.forEach(elementTD => {
        if(Number(elementTD.innerText)){
            arrayNotes.push(Number(elementTD.innerText))
        }
    })
    
    const [n1, n2, n3, n4] = arrayNotes
    console.log(arrayNotes)

    const avaregeTotal = avarege(n1, n2, n3, n4)

    td[td.length - 1].innerText = avaregeTotal.toFixed(2)
})