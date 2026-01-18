function processData(objData){
    const errorNameMessage = 'O nome precisa ser uma string ou ter pelo menos três caracteres'
    const errorAgeMessage = 'A idade precisa ser um number e não pode ser negativo'
    try {
        if (typeof objData.name !== 'string' || objData.name.length <= 3) throw new Error(errorNameMessage)
        
        if (typeof objData.age !== 'number' || objData.age.length < 0) throw new Error(errorAgeMessage)
        
        return 'sucess'
        
    } catch (error) {
        return error.message
    }
}

const listaParaTestes = [
  { name: "Alice", age: 28 },           
  { name: "", age: 30 },            
  { name: "Carla", age: "25" },        
  { name: 123, age: 40 },               
  { name: "Diego" },                      
  { name: "Elena", age: 18 },         
  { name: null, age: null }             
]

for (item of listaParaTestes){
    console.log(processData(item))
}