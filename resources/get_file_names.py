from os import listdir
from os.path import isfile, join

path = './cuvinte'
files = [f for f in listdir(path) if isfile(join(path, f))]

f = open('filenames','w')

for file in sorted(files):
	f.write(file+' ')