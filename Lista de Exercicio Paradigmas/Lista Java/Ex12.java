class Produto {
    private String nome;
    private double preco;

    public Produto(String nome, double preco) {
        this.nome = nome;
        this.preco = preco;
    }

    public Produto add(Produto outro) {
        if (outro instanceof Produto) {
            String novoNome = this.nome + " + " + outro.nome;
            double novoPreco = this.preco + outro.preco;
            return new Produto(novoNome, novoPreco);
        }
        return null;
    }

    @Override
    public String toString() {
        return "Produto(nome='" + nome + "', preco=" + String.format("%.2f", preco) + ")";
    }

    public static void main(String[] args) {
        Produto produto1 = new Produto("Produto A", 20.50);
        Produto produto2 = new Produto("Produto B", 15.75);

        Produto produtoSoma = produto1.add(produto2);
        System.out.println(produtoSoma);
    }
}
