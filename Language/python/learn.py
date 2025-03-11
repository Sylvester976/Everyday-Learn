print('Mustaaaaaaaarddddddd!!!!! \n turn the tv off')
print( '*'*10    )

# The `zone` variable is storing a multi-line string that contains the lyrics from the song "All of
# the Lights" by Kanye West.
zone = """  
All of the lights 
something wrong mj gone, that nigga dead
"""


print(zone.upper())
print(zone.lower())
print(zone.title())
print(zone.strip())

Loops in python
For loop

looper = [1,2,3,4]

for loopers in looper:
    print(loopers)
    
# For loop using range 
for num in range(20):
    print (str(num) +'\n')
    
for someNum in range(7,21):
    print(str(someNum) +'\n')
    
# WHILE LOOP
    incre = 1
    while incre < 20:
        print(incre)
        incre += 2
        

Prints out 0,1,2,3,4

count = 0
while True:
    print(count)
    count += 1
    if count >= 5:
        break

# Prints out only odd numbers - 1,3,5,7,9
for x in range(10):
    # Check if x is even
    if x % 2 == 0:
        continue
    print(x)

for else

successful = False
for number in range(3):
    print("Attempt") 
    if successful:
        print("successful")
        break
else:
    print("Attempted 3 times and failed")    

Nested loop
for x in range(5):
    for y in range(3):
        print(f"{x},{y}")

# iterables 
for x in  "siesmiology":
    print(x)
    
my_array = [1,2,3,4]
for y in my_array:
    print(y)
