const alunos = [
  {
    nome: 'Ana Silva',
    notas: [10, 8, 9, 7] 
  },
  {
    nome: 'Bruno Costa',
    notas: [7, 9, 6, 8]
  },
  {
    nome: 'Carla Dias',
    notas: [5, 4, 6, 3]
  },
  {
    nome: 'Daniel Oliveira',
    notas: [9, 10, 10, 9]
  }
]

const tbody = document.querySelector('#alunos tbody')
const fragment = document.createDocumentFragment()

alunos.forEach(aluno => {
    const tr = document.createElement('tr')
    const tdName = document.createElement('td')
    const tdMedia = document.createElement('td')
    const media = aluno.notas.reduce((acc, current) => acc + current, 0) / aluno.notas.length
    tdName.textContent = aluno.nome
    tdMedia.textContent = media.toFixed(2)
    tr.appendChild(tdName)
    tr.appendChild(tdMedia)
    fragment.appendChild(tr)
})

tbody.appendChild(fragment)