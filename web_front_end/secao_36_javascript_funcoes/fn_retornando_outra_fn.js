function fn(n1){
    return function fn2(n2){
        return n1 * n2
    }
}

const retornFN = fn(3)
console.log(retornFN)

console.log(retornFN(5))