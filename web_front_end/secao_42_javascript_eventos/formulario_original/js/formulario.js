const form = document.querySelector('#shopping-add')
let maxCharacters = 255
const msg = document.querySelector('#feedbackMessage')

form.addEventListener('submit', e => {
    e.preventDefault()
    const checked = form.querySelector('[type="checkbox"]').checked
    const title = form.querySelector('#txtTitulo').value
    const description = form.querySelector('textarea').value

    

    if(!checked || !title || !description) return

    
    msg.querySelector('h2').textContent = title
    msg.querySelector('p').textContent = description
    form.reset()
    msg.classList.toggle('show')
})

msg.addEventListener('click', e => {
    if(!e.target.classList.contains('fa-close')) return

    msg.classList.toggle('show')
})


form.querySelector('[type="checkbox"]').addEventListener('change', e => {
    form.querySelector('[type="submit"]').disabled = !form.querySelector('[type="submit"]').disabled
})

form.querySelector('textarea').addEventListener('input', e => {
    const count = form.querySelector('#contador span')
    
    if(e.currentTarget.value.length > maxCharacters) return

    const textCount = e.currentTarget.value.length

    count.textContent = maxCharacters - textCount

})