public class Carro {
    private String marca;
    private String modelo;
    private int ano;
    private int velocidade;

    public Carro(String marca, String modelo, int ano) {
        this.marca = marca;
        this.modelo = modelo;
        this.ano = ano;
        this.velocidade = 0; 
    }

    public String detalhes() {
        return "Marca: " + this.marca + ", Modelo: " + this.modelo + ", Ano: " + this.ano;
    }

    public void acelerar(int incremento) {
        this.velocidade += incremento;
        System.out.println("O carro acelerou para " + this.velocidade + " km/h.");
    }

    public void frear(int decremento) {
        this.velocidade -= decremento;
        if (this.velocidade < 0) {
            this.velocidade = 0; 
        }
        System.out.println("O carro freou para " + this.velocidade + " km/h.");
    }

    public void exibirVelocidade() {
        System.out.println("A velocidade atual é " + this.velocidade + " km/h.");
    }

    public static void main(String[] args) {
        Carro carro1 = new Carro("Toyota", "Corolla", 2020);
        Carro carro2 = new Carro("Honda", "Civic", 2019);
        Carro carro3 = new Carro("Ford", "Fiesta", 2021);

        System.out.println(carro1.detalhes());
        System.out.println(carro2.detalhes());
        System.out.println(carro3.detalhes());

        carro1.acelerar(50);
        carro1.exibirVelocidade();
        carro1.frear(20);
        carro1.exibirVelocidade();
    }
}
