const playlistSemana = {
  segunda: {
    dia: "Segunda-feira",
    trecho: "Segunda-feira, essa semana eu não vou trabalhar / Pois eu já tenho o meu dinheiro pra gastar.",
    artista: "Zeca Pagodinho"
  },
  terca: {
    dia: "Terça-feira",
    trecho: "Terça-feira, e eu ainda estou esperando / O seu telefone me ligar, e eu fico imaginando.",
    artista: "Art Popular"
  },
  quarta: {
    dia: "Quarta-feira",
    trecho: "Quarta-feira é dia de jogo, e eu vou pro estádio / Ver o meu time jogar, e a galera agitar.",
    artista: "Expressão Regueira"
  },
  quinta: {
    dia: "Quinta-feira",
    trecho: "Era uma quinta-feira, um dia de sol / E eu tava de bobeira, jogando futebol.",
    artista: "Charlie Brown Jr."
  },
  sexta: {
    dia: "Sexta-feira",
    trecho: "Sexta-feira, rotina e tédio / Mas eu não quero mais saber de nada.",
    artista: "Capital Inicial"
  },
  sabado: {
    dia: "Sábado",
    trecho: "Sábado à noite tudo pode mudar / Sábado à noite eu saio pra dançar.",
    artista: "Cidade Negra"
  },
  domingo: {
    dia: "Domingo",
    trecho: "Domingo, quero te encontrar / E desabafar todo o meu sofrer.",
    artista: "Só Pra Contrariar"
  }
}

const day = prompt('Digite um dia da semana em número por exemplo Domingo = 1, Segunda-Feira = 2... ')

switch (day) {
    case "1":
        alert(`${playlistSemana["domingo"].trecho} - ${playlistSemana["domingo"].artista}`)
        break
    case "2":
        alert(`${playlistSemana["segunda"].trecho} - ${playlistSemana["segunda"].artista}`)
        break
    case "3":
        alert(`${playlistSemana["terca"].trecho} - ${playlistSemana["terca"].artista}`)
        break
    case "4":
        alert(`${playlistSemana["quarta"].trecho} - ${playlistSemana["quarta"].artista}`)
        break
    case "5":
        alert(`${playlistSemana["quinta"].trecho} - ${playlistSemana["quinta"].artista}`)
        break
    case "6":
        alert(`${playlistSemana["sexta"].trecho} - ${playlistSemana["sexta"].artista}`)
        break
    case "7":
        alert(`${playlistSemana["sabado"].trecho} - ${playlistSemana["sabado"].artista}`)
        break

    default:
        alert('Digite um dia válido!')
        break;
}