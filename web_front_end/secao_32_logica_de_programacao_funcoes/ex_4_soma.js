function sum(...numbers){
    let sum = 0

    try {
        for (let number of numbers){
            if(typeof number !== 'number' || isNaN(number)) throw new Error(`${number} não é um número`);
            sum += number
        }

        return sum
    } catch (error) {
        return error
    }
}

console.log(sum(5, 5, 25))
console.log(sum(5, 'casa', 25))
console.log(sum(5, NaN, 25))

