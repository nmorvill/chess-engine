import struct

def float_to_hex(f):
    return hex(struct.unpack('<I', struct.pack('<f', f))[0])


def getDistanceToEdge(pos, direction):
    match direction:
        case 0 :
            return int((63 - pos) / 8)
        case 1 :
            return int(pos / 8)
        case 2 :
            return pos % 8
        case 3 :
            return 7 - pos % 8
        case 4 :
            return min(7-(pos%8), int((63-pos)/8))
        case 5 :
            return min(pos%8, int((63-pos)/8))
        case 6 :
            return min(7-(pos%8), int(pos/8))
        case 7 :
            return min(pos%8, int(pos/8))

def generate_knight():
    end = []
    for i in range(64):
            
        squares = 0
        leftMax = getDistanceToEdge(i, 2)
        upMax = getDistanceToEdge(i,0)
        rightMax = getDistanceToEdge(i, 3)
        downMax = getDistanceToEdge(i, 1)

        if leftMax >= 2 and upMax >= 1 :
            squares += 2 ** (i + 6)
        if leftMax >= 1 and upMax >= 2 :
            squares += 2 ** (i + 15)
        if leftMax >=1 and downMax >=2 :
            squares += 2 ** (i-17)
        if leftMax >=2 and downMax >=1 :
            squares += 2 ** (i-10)

        if rightMax >= 1 and upMax >= 2:
            squares += 2 ** (i + 17)    
        if rightMax >= 2 and upMax >= 1:
            squares += 2 ** (i + 10)
        if rightMax >= 2 and downMax >= 1:
            squares += 2 ** (i - 6)
        if rightMax >= 1 and downMax >=2:
            squares += 2 ** (i-15)

        end.append(hex(int(squares)))

    return end

def generate_king():
    end = []
    for i in range(64):
        squares = 0
        upMax = getDistanceToEdge(i,0)
        downMax = getDistanceToEdge(i, 1)
        leftMax = getDistanceToEdge(i, 2)
        rightMax = getDistanceToEdge(i, 3)
        upRightMax = getDistanceToEdge(i,4)
        upLeftMax = getDistanceToEdge(i,5)
        downRightMax = getDistanceToEdge(i,6)
        downLeftMax = getDistanceToEdge(i,7)

        if upMax >=1 :
            squares += 2 ** (i + 8)
        if downMax >= 1:
            squares += 2 ** (i-8)
        if leftMax >= 1 :
            squares += 2 ** (i-1)
        if rightMax >= 1 :
            squares += 2 ** (i+1)
        if upLeftMax >= 1 :
            squares += 2 ** (i+7)
        if upRightMax >= 1 :
            squares += 2 ** (i+9)
        if downLeftMax >= 1 :
            squares += 2 ** (i-9)
        if downRightMax >= 1 :
            squares += 2 ** (i-7)

        end.append(hex(int(squares)))

    return end
        
def generate_bishops() :
    end = []
    for i in range(64) :
        squares = 0
        for j in range(1, getDistanceToEdge(i,4) + 1) :
            squares += 2 ** (i + (j * 9))
        for j in range(1, getDistanceToEdge(i,5) + 1) :
            squares += 2 ** (i + (j * 7))
        for j in range(1, getDistanceToEdge(i,6) + 1) :
            squares += 2 ** (i + (j * -9))
        for j in range(1, getDistanceToEdge(i,7) + 1) :
            squares += 2 ** (i + (j * -7))

        end.append(hex(int(squares) & 35604928818740736))

    arr = list(divide_chunks(end, 8))
    result = map(lambda l : list(reversed(l)), arr)
    ret = [item for sublist in list(result) for item in sublist]
    return ret

def generate_rooks() :
    end = []
    for i in range(64) :
        squares = 0
        for j in range(1, getDistanceToEdge(i,0) + 1) :
            squares += 2 ** (i + (j * 8))
        for j in range(1, getDistanceToEdge(i,1) + 1) :
            squares += 2 ** (i + (j * -8))
        for j in range(1, getDistanceToEdge(i,2) + 1) :
            squares += 2 ** (i - j)
        for j in range(1, getDistanceToEdge(i,3) + 1) :
            squares += 2 ** (i + j)
        end.append(hex(int(squares) & 35604928818740736))

    arr = list(divide_chunks(end, 8))
    result = map(lambda l : list(reversed(l)), arr)
    ret = [item for sublist in list(result) for item in sublist]
    return ret

def divide_chunks(l, n):
    for i in range(0, len(l), n):
        yield l[i:i + n]

def generate_pawns_attacks(): 
    whites = []
    blacks = []
    for i in range(64):
        white_squares = 0
        black_squares = 0
            
        upRightMax = getDistanceToEdge(i,4)
        upLeftMax = getDistanceToEdge(i,5)
        downRightMax = getDistanceToEdge(i,6)
        downLeftMax = getDistanceToEdge(i,7)

        if upLeftMax >= 1 :
            white_squares += 2 ** (i+7)
        if upRightMax >= 1 :
            white_squares += 2 ** (i+9)
        if downLeftMax >= 1 :
            black_squares += 2 ** (i-9)
        if downRightMax >= 1 :
            black_squares += 2 ** (i-7)
        
        whites.append(hex(int(white_squares)))
        blacks.append(hex(int(black_squares)))
    return whites, blacks

#print(generate_knight())
#print(generate_king())
#print(generate_bishops())
#print(generate_rooks())
print(generate_pawns_attacks()[0])
print(generate_pawns_attacks()[1])