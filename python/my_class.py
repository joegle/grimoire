#!/usr/bin/env python2

class Animal():
    
    def __init__(self):
        self.name = None
        self.sound = None
        
    def makeSound(self):
        print self.sound

# Derived Class

class Dog(Animal):
    def __init__(self):
        Animal.__init__(self)
        self.name = "Dog"
        self.sound = "Woof"
        
if __name__ == "__main__":
    print "This is being run directly"
    a = Animal()
    a.makeSound()

    d = Dog()
    d.makeSound()
    
