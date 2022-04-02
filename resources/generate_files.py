lines = open('cuvinte_small.txt').readlines()

#output = open('output.txt', 'w')

letters = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z']
starting = []
for i in range(len(letters)):
    for j in range(len(letters)):
        starting.append(letters[i] + letters[j])

leng=len(lines)
for i in range(leng):
    line=lines[i]
    output = open('./cuvinte/'+line[0]+line[1], 'a')
    output.write(line)
    output.close()

    # txt=open(line[0]+line[1],'r').read()

    # txt=txt[:-1]
    
    # output = open(line[0]+line[1], 'w')
    # output.write(txt)
    # output.close()



    # for word in starting:
    #     if line[0] == word[0] and line[1] == word[1]:
    #         output.write(word + '\n')