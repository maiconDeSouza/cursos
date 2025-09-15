const openSearch = document.querySelector('#open-search')
const openDialog = document.querySelector('#search-dialog')
const closeDialog = document.querySelector('.close-search')

openSearch.addEventListener('click', e => {
    openDialog.showModal()
})

closeDialog.addEventListener('click', e => {
    openDialog.close()
})