const ul = document.querySelector('#todo-list')
const form = document.querySelector('#todo-add')



let arrayTasks = []

function saveLocalStore(){
    localStorage.setItem('tasksmcn', JSON.stringify(arrayTasks))
}

function getLocalStore(){
    const arrayTasks = JSON.parse(localStorage.getItem('tasksmcn')) || []
    return arrayTasks
}

function generateUUID(){
    return crypto.randomUUID()
}

function CreateNewTasks(task){
    return {
        id: generateUUID(),
        task,
        done: false,
    }
}

function addTaskInArrayTasks(objTask){
    arrayTasks.push(objTask)
    saveLocalStore()
}

function removeTaskInArrayTasks(id){
    const indexRemove = arrayTasks.findIndex(task => task.id === id)
    arrayTasks.splice(indexRemove, 1)
    saveLocalStore()
}

function checkedTaskInArrayTasks(id){
    const indexChecked = arrayTasks.findIndex(task => task.id === id)
    arrayTasks[indexChecked].done = !arrayTasks[indexChecked].done
    saveLocalStore()
}

function updateTaskInArrayTaks(id, newText){
    const indexNewText = arrayTasks.findIndex(task => task.id === id)
    arrayTasks[indexNewText].task = newText
    saveLocalStore()
}

function renderTasks(){
    const frag = document.createDocumentFragment()
    ul.replaceChildren()
    arrayTasks.forEach(task => {
        const li = document.querySelector('.todo-item').cloneNode(true)
        li.dataset.id = task.id
        li.querySelector('.task-name').textContent = task.task
        task.done && li.querySelector('.button-check i').classList.remove('none')
        frag.appendChild(li)
    })
    ul.appendChild(frag)
}

document.addEventListener('DOMContentLoaded', e => {
    arrayTasks = [...getLocalStore()] 
    renderTasks()
})

form.addEventListener('submit', e => {
    e.preventDefault()
    const task = form.querySelector('#item-input').value
    form.reset()

    const objTask = CreateNewTasks(task)

    addTaskInArrayTasks(objTask)
    
    const li = document.querySelector('.todo-item').cloneNode(true)
    li.dataset.id = objTask.id
    li.querySelector('.task-name').textContent = objTask.task
    li.querySelector('.button-check i').classList.add('none')

    ul.appendChild(li)
})

ul.addEventListener('click', e => {
    if(e.target.classList.contains('fa-trash-alt')){
        const id = e.target.parentElement.dataset.id
        e.target.parentElement.remove()
        removeTaskInArrayTasks(id)
    }

    
    if(e.target.closest('.button-check')){
        const checkBtn = e.target.closest('.button-check')
        const li = checkBtn.closest('.todo-item')
        const id = li.dataset.id
        checkBtn.querySelector('.fa-check').classList.toggle('none')
        checkedTaskInArrayTasks(id)
    }

    if(e.target.classList.contains('fa-edit')){
        const edit = e.target.parentElement.querySelector('.editContainer')
        edit.style.display = "block"

        edit.addEventListener('click', e => {
            if(e.target.classList.contains('cancelButton')){
                edit.style.display = "none"
            }

            if(e.target.classList.contains('editButton')){ 
                const id = e.currentTarget.parentElement.dataset.id
                const newText = e.currentTarget.querySelector('.editInput').value

                if(!newText)return

                e.currentTarget.parentElement.querySelector('.task-name').textContent = newText
                updateTaskInArrayTaks(id, newText)
                edit.style.display = "none"
            }
        })
    }
})



