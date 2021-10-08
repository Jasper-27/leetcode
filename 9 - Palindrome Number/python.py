
# def reverseString(x):


def isPalindrome(x):
    stringX = str(x)
    print (stringX)
    # revX = reverseString(stringX)

    revX = stringX[::-1]


    if (revX == stringX):
        return True 
    else: 
        return False 
    



print(isPalindrome(int(10)))
print(isPalindrome(int(101)))

print(isPalindrome(int(123123)))
print(isPalindrome(int(23)))
print(isPalindrome(int(0)))
print(isPalindrome(int(-121)))





# def greeting(name: str) -> str:
#     return 'Hello ' + name

# print(greeting("test"))
# print(reverseString("test"))