function gerarArrayNumerosAleatorios(quantidade, min, max) {
  const arrayAleatorio = []
  for (let i = 0; i < quantidade; i++) {
    
    const numero = Math.floor(Math.random() * (max - min + 1)) + min
    arrayAleatorio.push(numero)
  }
  return arrayAleatorio
}


const meusNumeros = gerarArrayNumerosAleatorios(10, 1, 20)
console.log(meusNumeros)

const numerosLimpos = meusNumeros.reduce((arr, current) => {
    if(!arr.includes(current)){
        arr.push(current)
        return arr
    }
    return arr
}, [])

console.log(numerosLimpos)