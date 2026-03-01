import { Tasks } from './Tasks.js'
import { Render } from './Render.js'

const tasks = new Tasks()
const render = new Render()

const form = document.querySelector('.addtask')
const ul = document.querySelector('.list-taks')
const dialog = document.querySelector('dialog')
const buttonUpdate = document.querySelector('.button-up')
const buttonClose = document.querySelector('.button-close')


document.addEventListener("DOMContentLoaded", e => {
    const arrTasks = tasks.returnArrTasks()
    console.log(arrTasks)
    render.renderListTasks(arrTasks)
})

form.addEventListener('submit', e => {
    e.preventDefault()

    const name = form.querySelector('input').value.trim()

    form.querySelector('input').value = ''

    if(!name)return

    const newTask = tasks.newTasks(name)

    const arrTasks = tasks.addNewTask(newTask)

    render.renderListTasks(arrTasks)
})

ul.addEventListener('click', e => {
    if(e.target.closest('.button-done')){
        const id = e.target.closest('.item-tasks').dataset.id
        const arrTaks = tasks.updateCompleted(id)
        render.renderListTasks(arrTaks)
    }

    if(e.target.closest('.button-remove')){
        const id = e.target.closest('.item-tasks').dataset.id
        const arrTasks = tasks.removeTask(id)
        render.renderListTasks(arrTasks)
    }

    if(e.target.closest('.button-edit')){
        dialog.showModal()
        const li = e.target.closest('.item-tasks')
        const id = li.dataset.id
        const name = li.querySelector('b').textContent

        dialog.querySelector('input').dataset.id = id
        dialog.querySelector('input').value = name
    }
})

buttonClose.addEventListener('click', e => {
    dialog.close()
})

buttonUpdate.addEventListener('click', e => {
    const id = dialog.querySelector('input').dataset.id 
    const name = dialog.querySelector('input').value 

    if(!name){
        dialog.close()
        return
    }

    const arrTasks = tasks.updateName(id, name)
    render.renderListTasks(arrTasks)
    dialog.close()
})