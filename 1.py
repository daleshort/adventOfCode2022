import heapq

file_path = "foodlist.txt"
f = open(file_path, "r")
calSum = 0
maxCalList = []

for line in f:
    if(line == '\n'):
        if(len(maxCalList) <3):  
            heapq.heappush(maxCalList,calSum)
        elif(calSum > maxCalList[0]):
            heapq.heappushpop(maxCalList,calSum)
        calSum = 0

    else:
        calSum += int(line)

print(sum(maxCalList))
