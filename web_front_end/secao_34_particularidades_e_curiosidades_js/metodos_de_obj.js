function vender(qtd){
    if(qtd > this.qtd){
        return `Quantidade insuficiente para realizar essa venda você precisa ter em estoque mais ${qtd - this.qtd} ${this.name}s`
    }
    this.qtd -= qtd
    return `venda feita com sucesso! você vendeu ${qtd} - seu estoque atual é de ${this.qtd}`
}

function addEstoque(qtd){
    this.qtd += qtd

    return `Estoque adicionado com sucesso! seu estoque agora é de ${this.qtd}`
}

function teste1(){
    return this
}

const teste2 = () => this

const produto = {
    name: 'caneta',
    qtd: 10,
    vender,
    addEstoque,
    teste1,
    teste2
}

console.log(produto.vender(12))
console.log(produto.addEstoque(4))
console.log(produto.vender(12))
console.log(produto.teste1())
console.log(produto.teste2())

