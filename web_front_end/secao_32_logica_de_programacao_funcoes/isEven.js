function isEven(number){
    const isValid = typeof number !== 'number' || isNaN(number)
    try {
        if(isValid) throw new Error('Digite apenas números')
        return number % 2 === 0
    } catch (error) {
        return error
    }
}


console.log(isEven(2))
console.log(isEven(3))
