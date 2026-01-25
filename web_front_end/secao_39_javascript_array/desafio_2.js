function sum(...values){
    const sumTotal = values.reduce((acc, current) => {
        return acc += current
    }, 0)
    return sumTotal
}

function avarage(...values){
    const avarage = sum(...values) / values.length
    return avarage
}


console.log(sum(3, 5, 8, 9))
console.log(avarage(3, 5, 8, 9))
