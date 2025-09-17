const buttonOpenSearch = document.querySelector('#open-search')
const buttonOpenSearchMobile = document.querySelector('#open-search-mobile')
const dialog = document.querySelector('#search-dialog')
const buttonCloseDialog = document.querySelector('.close-search')
const menuMobile = document.querySelector('#open-menu-mobile')
const dialogMenuMobile = document.querySelector('#menu-dialog-mobile')

buttonOpenSearch.addEventListener('click', e => {
    dialog.classList.toggle('open')
})

buttonOpenSearchMobile.addEventListener('click', e => {
    dialog.classList.toggle('open')
})

buttonCloseDialog.addEventListener('click', e => {
    dialog.classList.toggle('open')
})

menuMobile.addEventListener('click', e => {
    dialogMenuMobile.classList.toggle('open')
})

dialogMenuMobile.addEventListener('click', e => {
    dialogMenuMobile.classList.toggle('open')
})