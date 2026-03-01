export class Render {
    constructor(){
        this.form = document.querySelector('.addtask')
        this.ul = document.querySelector('.list-taks')
        this.template = document.querySelector('.template-task')
    }

    renderListTasks(arrTasks=[]){
        const frag = document.createDocumentFragment()
        this.ul.textContent = ""
        
        arrTasks.forEach(task => {
            const li = this.template.content.querySelector('.item-tasks').cloneNode(true)
            li.dataset.id = task.id
            li.querySelector('b').textContent = task.name

            task.completed && li.classList.add('done')

            frag.appendChild(li)
        })
        this.ul.appendChild(frag)
    }
}