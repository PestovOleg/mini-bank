# mini-bank

## Содержимое

- [ТЗ](#тз)
- [Юзкейсы](#юзкейсы)
- [Реализация](#реализация)
- [Сервисы](#сервисы)
- [Используемые сервисы](#используемые-сервисы-см-docker-compose)
- [API Routes](#api-routes)
- [Структура монорепо](#структура-монорепо)

## ТЗ
### Юзкейсы

    1. Я как пользователь хочу стать клиентом банка (регистрация)
    2. Я как клиент, хочу открыть счет
    3. Я как пользователь хочу пополнить счет под %
    4. Я как пользователь хочу видеть все свои счета

### DoD

- [x] Фичи реализуются 2 CRUD сервисами на Golang
    - Сервис 1 имеет REST API интерфейс
        - Create - создание пользователя ФИО, username, email, пароль
        - Read - поиск пользователя в базе
        - Update - обновление какой-либо информации пользователя
        - Delete - удаление пользователя из базы
    - Сервис 2 имеет REST API интерфейс
        - Create - создание счета пользователя 20 знаков
        - Read - получения счета, списка счетов пользователя
        - Update - обновления информации о счете
        - Delete - удаление счета
- [x] Каждая фича закрывается feature toggle
- [x] Информация о пользователях и счетах хранится в базе/базах PostgreSQL
- [x] Код сервисов хранится в GIT
- [x] Разработка ведется по GitFlow
- [x] Код бизнес логики сервисов покрывается Unit тестани
- [x] Код сервисов заворачивается в Docker образы
- [x] Сборка кода и деплой осуществляется автоматизированным DevOps Pipeline на выбранный хостинг
- [x] Деплой новых версий необходимо реализовать через Blue Green deployment
- [x] Обращение к сервисам осуществляется через прикладной балансировщик


## Реализация
### Структура монорепо

  - mini-bank/ 
  - ├── Makefile < *try* **'make help'** >
  - ├── README.md 
  - ├── backend 
    - ├── Makefile < *try* **'make help'** >
    - ├── config 
      - └── example.yml < *шаблон конфиг файла* >
    - ├── pkg < *local libs (to be git submodule)* >
      - ├── config 
      - ├── database 
      - ├── logger 
      - ├── middleware 
      - ├── server 
      - ├── signal 
      - └── unleash 
    - └── services < *микросервисы* >
      - ├── account 
      - ├── auth 
      - ├── mgmt 
      - └── user 
  - ├── db < *скрипты для DB* >
    - └── init-unleash.sh < *скрипт инициализации DB* >
  - ├── deploy.sh < *скрипт blue-green deployment* >
  - ├── docker-compose.yml 
  - ├── docs < *swagger docs* >
  - ├── go.work < *go workspace file* >
  - ├── nginx < *nginx конфиги* >
    - ├── conf.d < *шаблоны nginx для deploy.sh* >
      - ├── account-minibank.conf.template 
      - ├── auth-minibank.conf.template 
      - ├── mgmt-minibank.conf.template 
      - ├── user-minibank.conf.template 
      - ├── web.conf.template 
    - └── nginx.conf < *upstream конфиг* >
  - └── web < *frontend* >

### Сервисы
   - **mgmt-minibank** - сервис оркестрации работы с пользователями.
   - **user-minibank** - сервис работы с пользователями.
   - **account-minibank** - сервис работы со счетами.
   - **auth-minibank** - сервис аутентификации/авторизации.
   - **web** - frontend (to be)

   ### Используемые сервисы (см. docker-compose)
   - **Unleash** - <http://minibank.su:4242>
   - **Swagger** - <http://minibank.su:8001>


  ### API Routes
   Пример: <http://minibank.su/api/v1/mgmt-minibank-health>
| Service         | API (/api/v1)         | Method | Feature Toggle    | Basic Authorization | Description                      |
|-----------------|-----------------------|--------|-------------------|---------------------|----------------------------------|
| **mgmt-minibank**| `/mgmt-minibank-health`| GET   |                   |                     | Health Check                     |
|                 | `/mgmt`               | POST   | CreateUserToggle  |                     | Создание пользователя            |
|                 | `/mgmt/{id}`          | DELETE | DeleteUserToggle  |          +          | Удаление (деактивация) пользователя |
| **auth-minibank**| `/auth-minibank-health`| GET  |                   |                     | Health Check                     |
|                 | `/auth`               | POST   | CreateUserToggle  |                     | Создание записи аутентификации (логин/пароль) |
|                 | `/auth/login`               | POST    | AuthenticateToggle|          +          | Аутентификация                   |
|                 | `/auth/verify`          | POST    | AuthorizeToggle   |          +          | Авторизация пользователей к сервисам  |
|                 | `/auth/{id}`          | DELETE | DeleteUserToggle  |          +          | Удаление (деактивация) пользователя |
| **user-minibank**| `/user-minibank-health`| GET  |                   |                     | Health Check                     |
|                 | `/users`              | POST   | CreateUserToggle  |          +          | Данные пользователя              |
|                 | `/users/{id}`         | GET    | GetUserToggle     |          +          | Данные пользователя              |
|                 | `/users/{id}`         | PUT    | UpdateUserToggle  |          +          | Обновление данных пользователей   |
| **account-minibank**| `/account-minibank-health`| GET |              |                     | Health Check                     |
|                 | `/users/{userid}/accounts`           | POST   | CreateAccountToggle|         +          | Создание счета                   |
|                 | `/users/{userid}/accounts`           | GET    | ListAccountsToggle|         +          | Список счетов                    |
|                 | `/users/{userid}/accounts/{id}`      | PUT    | UpdateAccountToggle|        +          | Обновить данные по счету         |
|                 | `/users/{userid}/accounts/{id}`      | GET    | GetAccountToggle  |         +          | Информация о счете               |
|                 | `/users/{userid}/accounts/{id}`      | DELETE | DeleteAccountToggle|        +          | Удалить (деактивировать) счет    |
|                 | `/users/{userid}/accounts/{id}/topup`| PUT    | TopUpToggle       |         +          | Пополнить счет                   |
|                 | `/users/{userid}/accounts/{id}/withdraw`| PUT | WithdrawToggle   |         +          | Снять деньги со счета            |


### Github secrets and variables

<details>
<summary><b>Environment Variables использующиеся в Docker Compose, сервисах, и Blue-Green Deployment</b></summary>

**POSTGRES_PASSWORD** = `superSecure123`  

*адреса сервисов в сети docker (для общения между сервисами)*  
**AUTH_HOST** = `"http://nginx/api/v1/secureAuth"`
**AUTH_VERIFY_HOST** = `"http://nginx/api/v1/secureVerify"`  
**USER_HOST** = `"http://nginx/api/v1/secureUsers"`  
**ACCOUNT_HOST** = `"http://nginx/api/v1/secureAccounts"` 

*адрес для package-config.json*
**PUBLIC_URL** = `"http://minibank.su"` 

*для миграции текущего сервиса при выполнении скрипта deploy.sh*  
**MINIBANK_DB** = `orchestraDB`  
**MINIBANK_USER** = `orchestraUser`  
**MINIBANK_PASSWORD** = `orchestraPWD`  

*одноименные доступы сервисов к БД (также необходимы при инициализации БД)*  
**AUTH_MINIBANK_DB** = `authDB`  
**AUTH_MINIBANK_USER** = `authUser`  
**AUTH_MINIBANK_PASSWORD** = `authPWD`  

**USER_MINIBANK_DB** = `userDB`  
**USER_MINIBANK_USER** = `userUser`  
**USER_MINIBANK_PASSWORD** = `userPWD`  

**ACCOUNT_MINIBANK_DB** = `accountDB`  
**ACCOUNT_MINIBANK_USER** = `accountUser`  
**ACCOUNT_MINIBANK_PASSWORD** = `accountPWD`  

**DATABASE_DB** = `toggleDB`  
**DATABASE_PASSWORD** = `togglePWD`  
**DATABASE_USER** = `toggleUser`  

**UNLEASH_DB** = `unleashDB`  
**UNLEASH_PASSWORD** = `unleashPWD`  
**UNLEASH_USER** = `unleashUser`  

*версии сервисов,устанавливаемые при Blue-Green Deployment и пути к конфигам*  
**AUTH_APP_VERSION** = `latest`  
**AUTH_CONFIG_PATH** = `/etc/securePath/auth-config.yml`  
**USER_APP_VERSION** = `latest`  
**USER_CONFIG_PATH** = `/etc/securePath/user-config.yml`  
**ACCOUNT_APP_VERSION** = `latest`  
**ACCOUNT_CONFIG_PATH** = `/etc/securePath/account-config.yml`  
**MGMT_APP_VERSION** = `latest`  
**MGMT_CONFIG_PATH** = `/etc/securePath/mgmt-config.yml`
**WEB_APP_VERSION** = `latest`     
</details>