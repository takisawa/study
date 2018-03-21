import numpy as np
import matplotlib.pyplot as plt

train = np.loadtxt('click.csv', delimiter=',', skiprows=1)
train_x = train[:,0]
train_y = train[:,1]

plt.plot(train_x, train_y, 'o')
plt.show()

theta0 = np.random.rand()
theta1 = np.random.rand()

def f(x):
    return theta0 + theta1 * x

def E(x, y):
    return 0.5 * np.sum((y- f(x)) ** 2)

mu = train_x.mean()
sigma = train_x.std()

def standardize(x):
    return (x - mu) / sigma

train_z = standardize(train_x)

plt.plot(train_z, train_y, 'o')
plt.show()
