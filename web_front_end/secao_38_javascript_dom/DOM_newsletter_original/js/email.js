const form = document.querySelector('form')
const popover = document.querySelector('#pop')
const popoverRed = document.querySelector('#popred')
const emailRegex = /^[a-zA-Z0-9._%+-]+@([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}$/;


function validarEmailCompleto(email) {
  return emailRegex.test(email)
}



form.addEventListener('submit', e => {
    e.preventDefault()

    const email = form.querySelector('input').value

    form.querySelector('input').value = ''

    if(!email){
        return
    }

    if(!validarEmailCompleto(email)){
        popoverRed.togglePopover()

        setTimeout(() => {
            popoverRed.hidePopover()
        }, 2000)
        return
    }

    popover.togglePopover()

    setTimeout(() => {
        popover.hidePopover()
    }, 2000)
})