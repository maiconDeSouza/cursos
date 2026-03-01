export class Tasks {
    constructor(){
        this.arrayTasks = [
            {
                id: 'f72d4b1a-8e2c-4b9a-9d3e-5f1a2b3c4d5e',
                name: "task 1",
                completed: false,
                createdAt: 1592667375012,
                updatedAt: null
            },
            {
                id: 'a1b2c3d4-e5f6-4a5b-9c8d-7e6f5a4b3c2d',
                name: "task 2",
                completed: false,
                createdAt: 1581667345723,
                updatedAt: 1592667325018
            },
            {
                id: '8c9d0e1f-2a3b-4c5d-ae6f-7a8b9c0d1e2f',
                name: "task 3",
                completed: false,
                createdAt: 1592667355018,
                updatedAt: 1593677457010
            }
        ]
    }

    returnArrTasks(){
        this.arrayTasks.sort((a, b) => b.updatedAt - a.updatedAt).sort((a, b) => b.createdAt - b.createdAt)
        return this.arrayTasks
    }

    newTasks(name=""){
        return {
            id: crypto.randomUUID(),
            name,
            completed: false,
            createdAt: new Date().getTime(),
            updatedAt: null
        }
    }

    updatedAt(id){
        const index = this.arrayTasks.findIndex(task => task.id === id)
        this.arrayTasks[index].updatedAt = new Date().getTime()
        return this.returnArrTasks()
    }

    updateName(id, name){
        const index = this.arrayTasks.findIndex(task => task.id === id)
        this.arrayTasks[index].name = name
        this.updatedAt(id)
        return this.returnArrTasks()
    }

    updateCompleted(id){
        const index = this.arrayTasks.findIndex(task => task.id === id)
        this.arrayTasks[index].completed = !this.arrayTasks[index].completed
        this.updatedAt(id)
        return this.returnArrTasks()
    }

    removeTask(id){
        const index = this.arrayTasks.findIndex(task => task.id === id)
        this.arrayTasks.splice(index, 1)
        return this.returnArrTasks()
    }

    addNewTask(task={}){
        this.arrayTasks.push(task)
        return this.returnArrTasks()
    }
}

