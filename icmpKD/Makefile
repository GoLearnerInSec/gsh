TARGET=ikd
SOURCE=./crypto/aes.c ./crypto/sha1.c kd.c 
CC=gcc
CFLAGS+=-g -O0
LD_LIBRARY_PATH=

default:
	$(CC) $(CFLAGS) -I./include  $(SOURCE) -o $(TARGET)

clean:
	rm -rf *.o 

del:
	rm -rf *.o $(TARGET)
