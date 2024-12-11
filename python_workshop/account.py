class BankAccount:
    def __init__(self, name, balance):
        self.name = name
        self.balance = balance

    def deposit(self, amount):
        self.balance += amount

    def withdraw(self, amount):
        self.balance -= amount

    def display_balance(self):
        print(f"{self.name}'s balance is {self.balance} euros")

account = BankAccount("Celian", 1000)
account.deposit(500)
account.withdraw(200)
account.display_balance()

class Vehicle:
    def __init__(self, max_speed):
        self.max_speed = max_speed
        self.speed = 0
    
    def speed_up(self, amount):
        if self.speed + amount > self.max_speed:
            self.speed = self.max_speed
        else:
            self.speed += amount

class Car(Vehicle):
    def __init__(self, max_speed):
        super().__init__(max_speed)

    def speed_up(self, amount):
        print("that's racing")
        return super().speed_up(amount)
    
    def drift(self):
        print(f"deja vu: delivering sushi at {self.speed} kmh")

spaceship = Vehicle(1000)
spaceship.speed_up(300)
print(spaceship.speed)
car = Car(200)
car.speed_up(30)
print(car.speed)
car.drift()
