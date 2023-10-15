# mini-bank
  [minibank.su](http://minibank.su)

## Содержимое

- [ТЗ](#тз)
- [Юзкейсы](#юзкейсы)
- [Реализация](#реализация)
- [Сервисы](#сервисы)
- [Используемые сервисы](#используемые-сервисы-см-docker-compose)
- [API Routes](#api-routes)
- [Структура монорепо](#структура-монорепо)
- [Сервисы](#сервисы)
- [Сторонние сервисы](#используемые-сторонние-сервисы-см-docker-compose)
- [Routes](#api-routes)
- [Окружение](#github-secrets-and-variables)

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
      - ├── uproxy 
      - └── user 
  - ├── db < *скрипты для DB* >
    - └── init-*unleash*.sh < *скриптs инициализации DB* >
  - ├── compose < *доп скрипты для docker compose* >
    - └── compose-*.sh < *скрипт инициализации DB* >  
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
      - ├── uproxy-minibank.conf.template 
      - ├── web.conf.template 
    - └── nginx.conf < *upstream конфиг* >
  - └── web < *frontend* >

### Сервисы
   - **mgmt-minibank** - сервис оркестрации работы с пользователями.
   - **user-minibank** - сервис работы с пользователями.
   - **account-minibank** - сервис работы со счетами.
   - **auth-minibank** - сервис аутентификации/авторизации.
   - **uproxy-minibank** - сервис прокси для unleash (необходим для client-side сервисов).
   - **web** - react-frontend 

   ### Используемые сторонние сервисы (см. docker-compose)
   - **Unleash** - <http://minibank.su:4242>
   - **Swagger** - <http://minibank.su:8001>


  ### API Routes
   Пример: <http://minibank.su/api/v1/mgmt-minibank-health>
| Service           | API (/api/v1)             | Method | Feature Toggle        | Basic Authorization | Description                              |
| ----------------- | ------------------------- | ------ | --------------------- | ------------------- | ---------------------------------------- |
| **mgmt-minibank** | `/mgmt-minibank-health`   | GET    |                       |                     | Health Check                             |
|                   | `/mgmt`                   | POST   | CreateUserToggle      |                     | Create User                              |
|                   | `/mgmt/{id}`              | DELETE | DeleteUserToggle      | +                   | Delete (Deactivate) User                 |
| **auth-minibank** | `/auth-minibank-health`   | GET    |                       |                     | Health Check                             |
|                   | `/auth`                   | POST   | CreateUserToggle      |                     | Create Authentication Entry (Login/Password) |
|                   | `/auth/login`             | POST   | AuthenticateToggle    | +                   | Authenticate                             |
|                   | `/auth/verify`            | POST   | AuthorizeToggle       | +                   | Authorize Users to Services              |
|                   | `/auth/{id}`              | DELETE | DeleteUserToggle      | +                   | Delete (Deactivate) User                 |
| **user-minibank** | `/user-minibank-health`   | GET    |                       |                     | Health Check                             |
|                   | `/users`                  | POST   | CreateUserToggle      | +                   | User Data                                |
|                   | `/users/{id}`             | GET    | GetUserToggle         | +                   | User Data                                |
|                   | `/users/{id}`             | PUT    | UpdateUserToggle      | +                   | Update User Data                         |
| **account-minibank** | `/account-minibank-health` | GET  |                       |                     | Health Check                             |
|                   | `/users/{userid}/accounts` | POST   | CreateAccountToggle   | +                   | Create Account                           |
|                   | `/users/{userid}/accounts` | GET    | ListAccountsToggle    | +                   | List Accounts                            |
|                   | `/users/{userid}/accounts/{id}` | PUT | UpdateAccountToggle | +                   | Update Account Data                      |
|                   | `/users/{userid}/accounts/{id}` | GET | GetAccountToggle   | +                   | Account Information                      |
|                   | `/users/{userid}/accounts/{id}` | DELETE | DeleteAccountToggle | +                | Delete (Deactivate) Account              |
|                   | `/users/{userid}/accounts/{id}/topup` | PUT | TopUpToggle     | +                   | Top Up Account                           |
|                   | `/users/{userid}/accounts/{id}/withdraw` | PUT | WithdrawToggle | +                | Withdraw Money from Account              |
| **uproxy-minibank** | `/uproxy-minibank-health` | GET  |                      |                     | Health Check                             |
|                   | `/uproxy`                 | GET    |                       |                     | List Toggles                             |
| **web**            | `/web-minibank-health`    | GET    |                       |                     | Health Check                             |
|                   | `/`                       |       |                       |                     | Main Page                                |



### Github secrets and variables

<details>
<summary><b>Environment Variables использующиеся в Docker Compose, сервисах, и Blue-Green Deployment</b></summary>

**POSTGRES_PASSWORD** = `superSecure123`  

*адреса сервисов в сети docker (для общения между сервисами)*  
**AUTH_HOST** = `"http://nginx/api/v1/secureAuth"`
**AUTH_VERIFY_HOST** = `"http://nginx/api/v1/secureVerify"`  
**USER_HOST** = `"http://nginx/api/v1/secureUsers"`  
**ACCOUNT_HOST** = `"http://nginx/api/v1/secureAccounts"` 

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
**UPROXY_APP_VERSION**=latest
**UPROXY_CONFIG_PATH**= `/etc/securePath/uproxy-config.yml`

*для сборки web* 
**REACT_APP_URL**= `http://localhost/api/v1`
**PUBLIC_URL**= `http://localhost`
</details>