# Тестовое задание Рязанова Сергея
---

### [Ссылка на страницу вакансии](https://spb.hh.ru/vacancy/72784262?hhtmFrom=employer_vacancies)
---

### [Описание технического задания](https://github.com/MajotraderLucky/TestTask1/blob/master/backupsDb/%D0%91%D0%B0%D0%B7%D0%BE%D0%B2%D1%8B%D0%B9_DevOps.docx)

## Первая часть - задача по анализу алгоритма
![](https://github.com/MajotraderLucky/TestTask1/blob/master/backupsDb/arrBlock.png)

---
*Для решения данной задачи использовал следующий алгоритм:*
```go
package main

import "fmt"

func main() {
 numArr := [4]int{3, 2, 1, 5}
 var key, i, arrInI int
 fmt.Println(numArr)
 for j := 1; j <= len(numArr)-1; j++ {
  key = numArr[j]
  i = j - 1
  arrInI = numArr[i]
  fmt.Println(j, "-", key, i, "-", arrInI)
  // Let`s find out if iteration happens when
  // i > 0 and arrInI > key
  if i > 0 && arrInI > key {
   fmt.Println("In iteration", j, "condition i > 0 and arrInI > key is met")
   fmt.Println("This iteration have: j - 1 =", j-1, "key =", key, "i =", i)
   fmt.Println("And iteration", j, "have numArr[i] =", numArr[i])
   fmt.Println("Condition where A[i+1] = A[i] has the form:")
   fmt.Print("A[", i+1, "] = ", numArr[i+1], ":= A[", i, "] = ", numArr[i])
   fmt.Println("")
   numArr[i+1] = numArr[i]
   fmt.Println(numArr)
   i--
   fmt.Println(i)
  }
  fmt.Println("In iteration", j, "A[i+1] =", numArr[i+1], "key =", key)
  numArr[i+1] = key
  fmt.Println(numArr)
 }
}
```

[Ссылка на файл с кодом на языке *GO*](https://github.com/MajotraderLucky/TestTask1/blob/master/backupsDb/counterAlgo.go)

---
## Вторая часть - Bash-скрипты

---
### Файл настройки дампа БД PostgreSQL
```bash
#!/bin/bash

createDump="/PathTo/createDump.sh"
recoveryDB="/PathTo/recoveryDb.sh"
host="localhost"
db_name=""
db_user=""
db_schema=""
db_streams="5"
db_path=""
db_dump_path=""
db_dump_catalog_path=""
```

#### [Ссылка на файл](https://github.com/MajotraderLucky/TestTask1/blob/master/backupsDb/settings.sh)

---

### Файл создающий дамп БД PostgreSQL
```bash
#!/bin/bash

. /PathTo/settings.sh
. /PathTo/recoveryDb.sh
echo "$createDump"
echo "$recoveryDB"
echo "$info"

echo "pg_dump -U $db_user -F d -f $db_dump_catalog_path $db_name -j $db_streams $db_schema"
pg_dump -U $db_user -F d -f $db_dump_catalog_path $db_name -j $db_streams $db_schema
```
  #### [Ссылка на файл](https://github.com/MajotraderLucky/TestTask1/blob/master/backupsDb/createDump.sh)

---
### Файл восстанавливает БД
```bash
#!/bin/bash

. /PathTo/settings.sh

info="Recovery is connected"
pg_restore  -U $db_user  -d $db_name  -F d $db_dump_catalog_path
```

#### [Ссылка на файл](https://github.com/MajotraderLucky/TestTask1/blob/master/backupsDb/recoveryDb.sh)
---
