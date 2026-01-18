function calcular(v1, v2, operacao){
    try {
        if (typeof v1 !== 'number' || typeof v2 !== 'number' || isNaN(v1) || isNaN(v2)) throw new Error("operação invalida")
        
        switch (operacao) {
            case '+':
                return v1 + v2
            case '-':
                return v1 - v2
            case '*':
                return v1 * v2
            case '/':
                return v1 / v2
            default:
                throw new Error("operação invalida")
        }  
    } catch (error) {
        return error
    }
}

console.log(calcular(25, 23, '+'))