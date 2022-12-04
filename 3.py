import string

file_path = "day3.txt"
f = open(file_path, "r")
score = 0
# part 1
for line in f:
    if(len(line)%2 ==0 ):
        l = line
    else:
        l = line[:len(line)-1]
    left= l[:int(len(l)/2)]
    right = l[int(len(l)/2):]
    print("line", l, "left", left, "right", right)
    print(len(left), len(right),len(l))
    keepLooping = True
    for c in left:
        if keepLooping:
            if c in right:
                print(c)
                keepLooping = False
                if c in string.ascii_lowercase:
                    score += string.ascii_lowercase.find(c)+1
                else:
                    score += 27+ string.ascii_uppercase.find(c)
            else:
                pass
print(score)
f = open(file_path, "r")
score = 0
lineBuffer = [None,None,None]
lineBufferCount = 0

def processBuffer(buffer):
    common1 = getCommonChars(buffer[0],buffer[1])
    print(common1)
    common2 = getCommonChars(buffer[1],buffer[2])
    print(common2)
    finalCommon = getCommonChars(common1,common2)
    print(finalCommon)
    return getPriority(finalCommon[0])


def getCommonChars(list1, list2):
    common = []
    for c in list1:
        if c in list2:
            common.append(c)
    return common

def getPriority(c):
    if c in string.ascii_lowercase:
        return string.ascii_lowercase.find(c)+1
    else:
        return 27+ string.ascii_uppercase.find(c)

for line in f:
    lineBuffer[lineBufferCount] = line.strip('\n')
    lineBufferCount +=1
    if lineBufferCount>2:
        score += processBuffer(lineBuffer)
        lineBufferCount = 0
print(score)

