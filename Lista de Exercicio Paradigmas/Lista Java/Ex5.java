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
    public static void fazerSomAnimais(Animal[] animais) {
        for (Animal animal : animais) {
            System.out.println("O animal faz: " + animal.som());
        }
    }

    public static void main(String[] args) {
        Animal[] animais = {new Cachorro(), new Gato(), new Cachorro(), new Gato()};

        fazerSomAnimais(animais);
    }
}
