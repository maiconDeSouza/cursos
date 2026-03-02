export class LocalStore{
    constructor(){
        this.arrTasks = JSON.parse(localStorage.getItem('todo2.0'))
    }

    getArrTasks(){
        const arrTasks = this.arrTasks || []
        return arrTasks
    }

    setArrtasks(arrtasks=[]){
        localStorage.setItem('todo2.0', JSON.stringify(arrtasks))
    }
}