Для решения задачи был использован суффиксный массив, хранящийся на диске.
Int64Array инкапсулирует байтовый файл, представляя интерфейс а-ля массива.
SuffixArray - суффиксный массив, с int64Array вместо обычных массивов.

Чтобы запустить, надо установить go (я использовал 1.15.6). Затем склонировать репозиторий в go/src, зайти в него и запусть `go run . <имя файла>`.
Возможно, компилятор потребует сделать `go get -u github.com/stretchr/testify/assert`
Затем по одной строчке с клавиатуры вводятся строки для поиска. Если строка находится, выводится ее индекс,
 если нет, то -1. Чтобы закончить, надо ввести EOF.