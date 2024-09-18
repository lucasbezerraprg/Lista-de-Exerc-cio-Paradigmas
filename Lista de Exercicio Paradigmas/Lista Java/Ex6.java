class Motor {
    private int potencia;
    private String tipo;

    public Motor(int potencia, String tipo) {
        this.potencia = potencia;
        this.tipo = tipo;
    }

    public String detalhes() {
        return "Motor: " + this.potencia + " CV, Tipo: " + this.tipo;
    }
}

class Carro {
    private String marca;
    private String modelo;
    private int ano;
    private Motor motor;

    public Carro(String marca, String modelo, int ano, Motor motor) {
        this.marca = marca;
        this.modelo = modelo;
        this.ano = ano;
        this.motor = motor;
    }

    public String detalhes() {
        return "Carro: " + this.marca + " " + this.modelo + ", Ano: " + this.ano + ", " + this.motor.detalhes();
    }

    public static void main(String[] args) {
        Motor motor1 = new Motor(150, "Gasolina");
        Carro carro1 = new Carro("Toyota", "Corolla", 2020, motor1);

        Motor motor2 = new Motor(120, "Etanol");
        Carro carro2 = new Carro("Honda", "Civic", 2019, motor2);

        System.out.println(carro1.detalhes());
        System.out.println(carro2.detalhes());
    }
}
