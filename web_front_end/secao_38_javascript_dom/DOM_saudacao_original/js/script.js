const nameUser = 'Maicon'
const p = document.querySelector('.top-bar p')

// p.textContent += nameUser

const text = document.createTextNode(nameUser)

p.appendChild(text)