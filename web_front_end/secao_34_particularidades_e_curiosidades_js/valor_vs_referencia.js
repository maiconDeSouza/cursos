let n = 3

function add(n){
    n++
    console.log('valor n interno:', n)
}

add(n)

console.log('valor de n externo:', n)

const arr = [1, 2, 3]

function addArr(arr){
    arr.push(4)

    console.log('valor array interno:', arr)
}

addArr(arr)

console.log('valor array externo:', arr)

console.log((1234567890).toString(2))