import numpy as np
import pandas as pd


# From ndarray
print(np.random.randn(5))

s = pd.Series(np.random.randn(5), index=['a','b','c','d','e'])

print(s)
print(s.index)

print(pd.Series(np.random.randn(5)))



# from dict
d = {'b' : 1, 'a' : 0, 'c' : 2}
print(pd.Series(d))



# from scalar value
print(pd.Series(5, index=['a','b','c','d','e']))
