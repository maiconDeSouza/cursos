const mapMCN = function (fn){
    const newArr = Array(this.length)

    for(let i = 0; i <= this.length -1; i++){
        newArr[i] = fn(this[i], i, this)
    }
    return newArr
}

Array.prototype.mapMCN = mapMCN

const arr = [1, 2, 3]


const mcnArr = arr.mapMCN(value => value * 2)

console.log(mcnArr)