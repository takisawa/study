from typing import List

def greeting(name: str) -> str:
    return 'Hello ' + name


print(greeting('taro'))


def greetings(names: List) -> str:
    return ', '.join(['Hello ' + n for n in names])

print(greetings(['taro', 'hanako']))
