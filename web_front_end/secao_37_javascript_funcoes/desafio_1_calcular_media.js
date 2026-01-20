function media(...numbers) { 
    if (numbers.length === 0) return 0

    let total = 0
    
    for (const num of numbers) {
        if (typeof num !== 'number' || Number.isNaN(num)) {
            throw new TypeError('Só é aceito números') 
        }
        total += num
    }

    return (total / numbers.length).toFixed(2)
}


try {
    console.log(media())        
    console.log(media(1, 4, 5))    
    console.log(media(1, "a"))
} catch (error) {
    console.error("Erro detectado:", error.message)
}
