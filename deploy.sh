#!/usr/bin/env bash
# проверка переменной имени сервиса
if [[ -z $SERVICE ]]; then
  echo "var SERVICE is empty. Exiting."
  exit 1
fi

echo "Starting deploy ${SERVICE}..."
echo "Checking nginx template for service ${SERVICE}"
[ -f "./nginx/conf.d/${SERVICE}.conf.template" ] && 
echo "Nginx template for ${SERVICE} is found" || 
echo "Nginx template for ${SERVICE} not found"

# Список конфигов на проверку-TODO:удалить после dynamic upstream
files=(
    "./nginx/conf.d/mgmt-minibank.nginx.conf"
    "./nginx/conf.d/account-minibank.nginx.conf"
    "./nginx/conf.d/auth-minibank.nginx.conf"
    "./nginx/conf.d/user-minibank.nginx.conf"
    "./nginx/conf.d/uproxy-minibank.nginx.conf"
    "./nginx/conf.d/web-minibank.nginx.conf"
)

# Проверяем наличие конфигов для всех сервисов
for file in "${files[@]}"; do
    if [ ! -e "$file" ]; then
        # Генерируем имя файла *.conf.template на основе имени файла
        template_file="${file%%.nginx.conf}.conf.template"
        service_name=$(basename "$file" | cut -d. -f1)
        # Копируем файл *.conf.template на место отсутствующего файла
        cp "$template_file" "$file"
        sed -i "s|proxy_pass http://.*;|proxy_pass http://$service_name-blue;|g" "$file"
        echo "ALARM!!! Скопирован файл ${SERVICE} $template_file в $file ,необходимо изменение содержимого после деплоя файла"
    fi
done

echo "Checking current nginx.conf"
if grep -q "proxy_pass http://${SERVICE}-blue" "./nginx/conf.d/${SERVICE}.nginx.conf"
then
    export CURRENT_SERVICE="${SERVICE}-blue"
    export NEXT_SERVICE="${SERVICE}-green"
    echo "$CURRENT_SERVICE is running"
else 
    export CURRENT_SERVICE="${SERVICE}-green"
    export NEXT_SERVICE="${SERVICE}-blue"
    echo "$CURRENT_SERVICE is running"
fi

echo "Removing old container of ${SERVICE} if it hasn't been removed..."
docker compose rm -f -s -v $NEXT_SERVICE 
echo "Waiting 3 sec"
sleep 3

if [ "$MIGRATE" == "YES" ]; then
  echo "Migrating database ..."
  docker compose run -d $MIGRATE_SERVICE 
  printf "%s\n" "Done...waiting 3 sec"
  sleep 3
else 
  echo "Skipping database migration."
fi

echo "Starting $NEXT_SERVICE"
docker compose up -d $NEXT_SERVICE 
rv=$?
if [ $rv -eq 0 ]; then
    echo "New \"$NEXT_SERVICE\" container started"
else
    echo "Docker compose failed with exit code: $rv"
    echo "Aborting..."
    exit 1
fi
echo "Waiting for starting "
sleep 10

# Извлекаем динамический порт из запущенного контейнера
PORT=$(docker port $NEXT_SERVICE | cut -d':' -f2)

# Проверяем, что удалось получить порт
if [ -z "$PORT" ]; then
  echo "Failed to get port for $NEXT_SERVICE. Exiting."
  exit 1
fi

# Создаем URL с динамическим портом
URL="http://0.0.0.0:$PORT/api/v1/${SERVICE}-health"

echo "Checking..."

# Выполняем вызов через curl
response=$(curl -s $URL)

# Проверяем ответ
if echo "$response" | grep -q "Service is healthy"; then
  echo "Service is healthy"
else
  echo "Service is not healthy"
fi
#---------------------------------------------------------------

echo "Copying current nginx.conf to nginx.conf.back"
cp ./nginx/conf.d/${SERVICE}.nginx.conf ./nginx/conf.d/${SERVICE}.conf.back 2>/dev/null
cp ./nginx/conf.d/${SERVICE}.conf.template ./nginx/conf.d/${SERVICE}.nginx.conf 2>/dev/null

echo "Creating and checking nginx config for next service"
sed -i "s|proxy_pass http://.*;|proxy_pass http://$NEXT_SERVICE:3333;|g" ./nginx/conf.d/${SERVICE}.nginx.conf
docker compose exec -T nginx nginx -g 'daemon off; master_process on;' -t
rv=$?
if [ $rv != 0 ]; then
    cp ./nginx/conf.d/${SERVICE}.nginx.conf.back ./nginx/conf.d/${SERVICE}.nginx.conf 2>/dev/null
    echo "Checking failed with exit code: $rv"
    echo "Aborting..."
    exit 1
fi

echo "Reloading nginx with $NEXT_SERVICE"
docker compose exec -T nginx nginx -g 'daemon off; master_process on;' -s reload 
rv=$?
if [ $rv != 0 ]; then
    echo "Reloading is failed with exit code: $rv"
    echo "Aborting..."
    cp ./nginx/conf.d/${SERVICE}.nginx.conf.back ./nginx/conf.d/${SERVICE}.nginx.conf 2>/dev/null
    exit 1
else    
    echo "Nginx reloaded"
fi

echo "Testing minibank..."
curl -s http://0.0.0.0/api/v1/${SERVICE}-health | grep "Service is healthy"
rv=$?
if [ $rv != 0 ]; then
    echo "Testing is failed with exit code: $rv"
    echo "Aborting..."
    cp ./nginx/conf.d/${SERVICE}.nginx.conf ./nginx/conf.d/${SERVICE}.error.conf.back 2>/dev/null
    cp ./nginx/conf.d/${SERVICE}.nginx.conf.back ./nginx/conf.d/${SERVICE}.nginx.conf 2>/dev/null
    echo "Reloading nginx with $CURRENT_SERVICE"
    docker compose exec -T nginx nginx -g 'daemon off; master_process on;' -s reload 
    echo "Nothing has happend,do not forget about migration..."
    exit 1
else    
    echo "Testing is OK"
fi

echo "Deleting old container: $CURRENT_SERVICE"
docker compose rm -f -s -v $CURRENT_SERVICE 

echo "Blue-Green deployment is done!"