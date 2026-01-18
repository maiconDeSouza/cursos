const arr = [1, 2, 3, 4, 5, 6, 7, 8, 9]

const arrDouble1 = Array(arr.length)

console.log(arrDouble1)
for(let i = 0; i <= arr.length -1; i++){
    arrDouble1[i] = (arr[i] * 2)
}

const arrDouble2 = arr.map(n => n * 2)

console.log(arr)
console.log(arrDouble1)
console.log(arrDouble2)

console.log([].map)