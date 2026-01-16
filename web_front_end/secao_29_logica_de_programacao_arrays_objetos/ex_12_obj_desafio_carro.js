const car = {
    modelo: 'Uno',
    ano: 2023,
    km: 10000,
    combustivel: 'Gasolina',
    litrosConsumo: 625
}

console.log(`O carro ${car.modelo} ${car.ano} fez em média ${car.km / car.litrosConsumo} km/l de ${car.combustivel}`)