# mini-bank

## Содержимое

- [ТЗ](#тз)
- [Юзкейсы](#юзкейсы)
- [Реализация](#реализация)
- [Сервисы](#сервисы)
- [API Routes](#api-routes)
- [Структура монорепо](#структура-проекта)

## ТЗ
### Юзкейсы

    1. Я как пользователь хочу стать клиентом банка (регистрация)
    2. Я как клиент, хочу открыть счет
    3. Я как пользователь хочу пополнить счет под %
    4. Я как пользователь хочу видеть все свои счета

### DoD

    1. [x] Фичи реализуются 2 CRUD сервисами на Golang
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
    2. [x] Каждая фича закрывается feature toggle
    3. [x] Информация о пользователях и счетах хранится в базе/базах PostgreSQL
    4. [x] Код сервисов хранится в GIT
    5. [x] Разработка ведется по GitFlow
    6. [x] Код бизнес логики сервисов покрывается Unit тестани
    7. [x] Код сервисов заворачивается в Dockeг образы
    8. [x] Сборка кода и деплой осуществляется автоматизированным DevOps Pipeline на выбранный хостинг
    9. [x] Деплой новых версий необходимо реализовать через Blue Green deployment
    10. [x] Обращение к сервисам осуществляется через прикладной балансировщик

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
   
  ### API Routes
   Пример: <http://minibank.su/api/v1/mgmt-minibank-health>
| Service         | API (/api/v1)         | Method | Feature Toggle    | Basic Authorization | Description                      |
|-----------------|-----------------------|--------|-------------------|---------------------|----------------------------------|
| **mgmt-minibank**| `/mgmt-minibank-health`| GET   |                   |                     | Health Check                     |
|                 | `/mgmt`               | POST   | CreateUserToggle  |                     | Создание пользователя            |
|                 | `/mgmt/{id}`          | DELETE | DeleteUserToggle  |          +          | Удаление (деактивация) пользователя |
| **auth-minibank**| `/auth-minibank-health`| GET  |                   |                     | Health Check                     |
|                 | `/auth`               | POST   | CreateUserToggle  |                     | Создание записи аутентификации (логин/пароль) |
|                 | `/auth`               | GET    | AuthenticateToggle|          +          | Аутентификация                   |
|                 | `/auth/{id}`          | GET    | AuthorizeToggle   |          +          | Авторизация пользователей к сервисам  |
|                 | `/auth/{id}`          | DELETE | DeleteUserToggle  |          +          | Удаление (деактивация) пользователя |
| **user-minibank**| `/user-minibank-health`| GET  |                   |                     | Health Check                     |
|                 | `/users`              | POST   | CreateUserToggle  |          +          | Данные пользователя              |
|                 | `/users/{id}`         | GET    | GetUserToggle     |          +          | Данные пользователя              |
|                 | `/users/{id}`         | PUT    | UpdateUserToggle  |          +          | Обновление данных пользователей   |
| **account-minibank**| `/account-minibank-health`| GET |              |                     | Health Check                     |
|                 | `/accounts`           | POST   | CreateAccountToggle|         +          | Создание счета                   |
|                 | `/accounts`           | GET    | ListAccountsToggle|         +          | Список счетов                    |
|                 | `/accounts/{id}`      | PUT    | UpdateAccountToggle|        +          | Обновить данные по счету         |
|                 | `/accounts/{id}`      | GET    | GetAccountToggle  |         +          | Информация о счете               |
|                 | `/accounts/{id}`      | DELETE | DeleteAccountToggle|        +          | Удалить (деактивировать) счет    |
|                 | `/accounts/{id}/topup`| PUT    | TopUpToggle       |         +          | Пополнить счет                   |
|                 | `/accounts/{id}/withdraw`| PUT | WithdrawToggle   |         +          | Снять деньги со счета            |


## Пререквизиты
### Github variables

    | VARIABLE    |                                         |
    | ----------- | --------------------------------------- |
    | CONFIG_PATH | Путь к конфигу на сервере               |
    | SECRET_PATH | Путь к секретам на сервере (БД и т.д.)  |

### Github secrets

    | VARIABLE              |                               |
    | --------------------- | ----------------------------- |
    | DOCKERHUB_TOKEN       | токен docker registry         |
    | DOCKERHUB_USERNAME    | user dockerhub registry       |
    | SSH_PRIVATE_KEY       |                               |
    | SSH_SERVER_IP         | ip где разворачиваем          |
    | SSH_SERVER_USER       |                               |


### На стороне сервера

    | CONFIG      |                                             |
    | ----------- | ------------------------------------------- |
    | CONFIG_PATH | Путь к конфигу (-см. Variables)             |
    | SECRET_PATH | Путь к секретам на сервере (-см. Variables) |

## Endpoints (v1):

    - /v1/health-check  //проверка состояния сервера