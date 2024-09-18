class Calculadora {
    public int somarDois(int a, int b) {
        return a + b;
    }

    public int somarTres(int a, int b, int c) {
        return a + b + c;
    }

    public static void main(String[] args) {
        Calculadora calc = new Calculadora();

        int resultado1 = calc.somarDois(5, 10);
        int resultado2 = calc.somarTres(5, 10, 15);

        System.out.println("Soma de dois números: " + resultado1);
        System.out.println("Soma de três números: " + resultado2);
    }
}
