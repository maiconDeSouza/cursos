const messagesIMC = {
    imc1: 'Muito abaixo do peso',
    imc2: 'Abaixo do peso',
    imc3: 'Peso normal',
    imc4: 'Acima do peso',
    imc5: 'Obesidade Grau I',
    imc6: 'Obesidade Grau II',
    imc7:  'Obesidade Grau III',
    error: {
        err1: 'Erro na hora de setar os parametros. É obrigatório preencher os dois e apenas com números',
    } 
}

function returnText(imc, text){
    return `Seu imc é de ${imc.toFixed(2)} e você está ${text}`
}

function CalculateBMI(weight, height){
    const imc = weight / (height ** 2)

    if(typeof imc !== 'number' || isNaN(imc) || !isFinite(imc) || imc <= 0) throw new Error(messagesIMC.error.err1)

    if(imc > 40) return returnText(imc, messagesIMC.imc7)

    if(imc > 34.9) return returnText(imc, messagesIMC.imc6)
    
    if(imc > 29.9) return returnText(imc, messagesIMC.imc5)
    
    if(imc > 24.9) return returnText(imc, messagesIMC.imc4)
    
    if(imc > 18.5) return returnText(imc, messagesIMC.imc3)
    
    if(imc > 16.9) return returnText(imc, messagesIMC.imc2)
    
    if(imc <= 16.9) return returnText(imc, messagesIMC.imc1)
}

/*
Muito abaixo do peso 16 a 16,9 kg/m2

Abaixo do peso 17 a 18,4 kg/m2

Peso normal 18,5 a 24,9 kg/m2

Acima do peso 25 a 29,9 kg/m2

Obesidade Grau I 30 a 34,9 kg/m2

Obesidade Grau II 35 a 40 kg/m2

Obesidade Grau III maior que 40 kg/m2
*/

try {
    console.log(CalculateBMI(120, 0))
} catch (error) {
    console.error(error.message)
}