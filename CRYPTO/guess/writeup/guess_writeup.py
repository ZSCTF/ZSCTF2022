from pwn import *

# context.log_level = "debug"
p = remote("127.0.0.1", 1234)

allChar = string.digits + string.ascii_letters + string.punctuation
flag = ""

for i in range(1, 27):
    abnormalTime = float()
    abnormalChar = ""
    for chunk in allChar:
        send = flag + chunk + "a" * (27 - i)

        startTime = time.time()
        for times in range(10):
            p.sendafter('flag:\n', str.encode(send))
            p.recvuntil(str.encode("error"))
        endTime = time.time()
        useTime = endTime - startTime

        if abnormalTime < useTime:
            abnormalTime = useTime
            abnormalChar = chunk

    flag += abnormalChar
    print(flag)
