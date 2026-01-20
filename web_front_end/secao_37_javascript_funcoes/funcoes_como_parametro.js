function mapMCN(cb, thisArg){
    const newArr = Array(this.length)
    if (typeof cb !== 'function') {
        throw new TypeError(cb + ' is not a function');
    }

    for(let i = 0; i <= this.length -1; i++){
        if(i in this){
            newArr[i] = cb.call(thisArg, this[i], i, this)
        }
    }

    return newArr
}

Array.prototype.mapMCN = mapMCN


const arr = [1, 2, 3]

const newArr = arr.mapMCN(value => value * 2)

console.log(arr)
console.log(newArr)


