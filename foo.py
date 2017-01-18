import time
import random


def hello():
    """
    Print helloworld
    """
    print "hello, world!"

def print_odds(limit=10):
    """
    Print odds numbers < limit
    """
    for i in range(limit):
        if i%2:
            print i
        time.sleep(0.1)

def print_even(limit=10):
    """
    Print even numbers < limit
    """
    for i in range(limit):
        if i%2 == 0:
            print i
        time.sleep(0.1)

def print_rand():
    """
    return a random float
    """
    return random.random()
