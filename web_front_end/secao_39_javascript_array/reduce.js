const names = ['Daniel', 'Maria', 'João', 'Joana']

const obj = names.reduce((acc, current) => {
    const primaryLetter = current[0]
    if(primaryLetter in acc){
        acc[primaryLetter]++
        return acc
    }
    acc[primaryLetter] = 1
    return acc
}, {})

console.log(obj)