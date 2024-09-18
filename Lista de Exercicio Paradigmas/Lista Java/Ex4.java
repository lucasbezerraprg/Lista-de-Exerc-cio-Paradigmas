abstract class Animal {
    abstract String som();
}

class Cachorro extends Animal {
    @Override
    String som() {
        return "Au Au!";
    }
}

class Gato extends Animal {
    @Override
    String som() {
        return "Miau!";
    }
}

public class Main {
    public static void fazerSom(Animal animal) {
        System.out.println("O animal faz: " + animal.som());
    }

    public static void main(String[] args) {
        Cachorro cachorro = new Cachorro();
        Gato gato = new Gato();

        fazerSom(cachorro);
        fazerSom(gato);
    }
}
