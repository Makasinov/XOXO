### Ссылки
https://docs.google.com/spreadsheets/d/1OBlOwu-R8ieAGCyKz0-7pLMdB72VxjxBxhP1reyyqz8/edit#gid=0

### Взаимодействие происходит через Makefile
```shell
# запуск через бинарник
$ ./bin -latitude 55.754376 -longitude 37.559364

# запуск без бинарника
$ make run ARGS="-latitude 55.754376 -longitude 37.559364"
```


Пример использования
```shell
$ ./bin -latitude 55.754376 -longitude 37.559364
Таганско-Краснопресненская Улица 1905 года
Разница: 0.009999863649072056
55.754376 37.559364
55.763944 37.562271
```
