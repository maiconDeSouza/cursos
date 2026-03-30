// const arr1 = [1, 2, 3, 4]
// const arr2 = new Array('abc')
// const arr3 = Array.of(5)
// const arr4 = Array.from('abc')
// const arr5 = new Array(5)
// const misto = [42, 'texto', undefined, true, null]

// arr5[2] = 'casa'

// console.log(arr1)
// console.log(arr2)
// console.log(arr3)
// console.log(arr4)
// console.log(arr5[3])
// console.log(arr5)

// arr5.forEach(item => {
//     console.log(item)
// })

// misto.forEach(item => {
//     console.log(item)
// })

const text = 'abc'
const arr = Array.from(text)

function returnArray(){
    const arr = Array.from(arguments)
    return arr
}

const resultArr = returnArray(1, 2, 'casa', true, null, 21.23, {}, [])

console.log(arr)
console.log(resultArr)

// [ 'a', 'b', 'c' ]
// [1, 2, 'casa', true, null, 21.23, {}, []]




