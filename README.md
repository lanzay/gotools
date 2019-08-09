# gotools

Сборник небоьших библиотек/инструментов
### LeftRuneValid(s string, lenS int) string
Получает точное количество символов слева и проверяет на валидность итоговую строку 

### LeftRuneMax(str string, l int) string
Получить левую часть строки до целого количества RUNE - не ломая кодировку


### func SubByte(body []byte, startStr string, startShift int, finStr string, finShift int) []byte
Получает подстроку байт по указанному началу и окончанию строки
также задается смещение границ относительо поисковых строк
