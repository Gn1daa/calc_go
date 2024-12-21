# calc_go
# Вебъ​-сервисъ для вычисленія ариѳметическихъ выраженій на языкѣ Го 
## Описаніе 
Данный проектъ реализуетъ ​вебъ​-серверъ, который вычисляетъ ​ариѳметическія​ выраженія, ​переданныя​ пользователемъ въ форматѣ JSON черезъ HTTP запросъ. Вычисленіе происходитъ путемъ перевода выраженія въ обратную польскую нотацію и послѣдующее вычисленіе отвѣта.
## Запускъ 
1. Установите [Go](https://go.dev/dl/).
2. Установите [Git](https://git-scm.com/downloads). 
3. ​Клонируйте​ ​репозиторій​ съ помощью
 ```bash
git clone github.com/Gn1daa/calc_go
```
4. Выполните команду
 ```bash
 go mod tidy
 ```
5. Установите ​всѣ​ не​обходимыя​ библіотеки 
6. Запустите серверъ съ помощью
 ```bash
 go run ./cmd/main.go
 ```
