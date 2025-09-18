    PZ2 Технологии индустриального программирования.  
    Студент: Выборнов О.А.  
    Группа: ЭФМО-02-25  

Как запустить:

1)склонировать репозиторий:

        git clone -b main --single-branch https://github.com/omnikk/PZ2.git

2.1)Через .exe файл

2.2)Через cmd командой "go run ./cmd/myapp"

Проверка Go и Git (go version, git --version);  
Инициализация проекта (go mod init);  
Создание структуры каталогов (cmd/myapp, internal/app, utils);  
Реализация логирования (utils/log.go);  
Минимальный сервер на порту 8080;  
Middleware для Request ID;  
Единый JSON-формат ошибок (/fail);  
Обработчик /ping с текущим временем RFC3339;  

![Гифка с Gifius ru (2)](https://github.com/user-attachments/assets/3d8f4bcd-ef55-4ab2-8416-4e904dec2a4a)

Пример запросов:

    http://localhost:8080/       Hello, Go project structure!  
    
<img width="661" height="134" alt="image" src="https://github.com/user-attachments/assets/2c809896-1845-4a00-af1c-034de702a7ae" />
    
    http://localhost:8080/ping   {"status": "ok","time": "2025-09-18T18:35:28Z"}
    
<img width="594" height="155" alt="image" src="https://github.com/user-attachments/assets/80c04b34-e513-44fc-a06b-c9476a790244" />

    http://localhost:8080/fail   {"error":"bad_request_example"} 
    
<img width="605" height="155" alt="image" src="https://github.com/user-attachments/assets/2631e1c1-8314-4100-8bb1-5d473888860d" />


Cтруктура проекта:

    myapp/
      cmd/
        myapp/
          main.go
      internal/
        app/
          app.go
          handlers/
            ping.go
      utils/
        log.go
        id.go
        errors.go
      go.mod
      go.sum
      README.md

Описание структуры:

    cmd/myapp - точка входа в приложение  
    internal/app - основная логика приложения и обработчики  
    utils - вспомогательные функции (логирование, генерация ID, ошибки)  
    go.mod - файл модуля Go с зависимостями  
    go.sum - контрольные суммы зависимостей  


