import collections

dict1 = {'day1': 1, 'day2': 2}
dict2 = {'day3': 3, 'day4': 4}

res = collections.ChainMap(dict1, dict2)

print(res.maps)

print('Keys = {}'.format(list(res.keys())))
print('Values = {}'.format(list(res.values())))

# Print all the elements from the result
print('elements:')
for key, val in res.items():
    print('{} = {}'.format(key, val))
print()

# Find a specific value in the result
print('day3 in res: {}'.format(('day1' in res)))
print('day4 in res: {}'.format(('day4' in res)))