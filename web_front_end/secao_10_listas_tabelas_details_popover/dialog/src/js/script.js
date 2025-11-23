const modalPromo = document.querySelector('dialog')
const buttonShowPromo = document.querySelector('.btn-promo')
const buttonsAction = document.querySelectorAll('.div-promo-container button')

let count = 0

buttonShowPromo.addEventListener('click', e => {
    count <= 3 && modalPromo.showModal()
    count += 1
})

buttonsAction.forEach(button => {
    button.addEventListener('click', e => {
        modalPromo.close()
    })
})

document.addEventListener('mouseleave', e => {
    count <= 3 && modalPromo.showModal()
    count += 1
})