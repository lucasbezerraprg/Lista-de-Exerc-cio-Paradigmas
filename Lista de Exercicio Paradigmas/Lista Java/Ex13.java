class Matematica {
    public static int fatorial(int n) {
        if (n < 0) {
            throw new IllegalArgumentException("O fatorial não está definido para números negativos.");
        } else if (n == 0 || n == 1) {
            return 1;
        } else {
            int resultado = 1;
            for (int i = 2; i <= n; i++) {
                resultado *= i;
            }
            return resultado;
        }
    }

    public static void main(String[] args) {
        int numero = 5;
        System.out.println("O fatorial de " + numero + " é " + Matematica.fatorial(numero) + ".");
    }
}
