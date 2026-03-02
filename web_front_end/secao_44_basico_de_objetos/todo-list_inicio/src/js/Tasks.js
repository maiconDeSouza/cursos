export class Tasks {
    constructor(arrTasks=[]){
        this.arrayTasks = arrTasks
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

