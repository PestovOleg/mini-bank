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

echo "Removing old container if it hasn't been removed..."
docker compose rm -f -s -v $NEXT_SERVICE 
echo "Waiting 3 sec"
sleep 3
printf "%s\n" "Done"

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
echo "Waiting 3 sec"
sleep 3

echo "Copying current nginx.conf to nginx.conf.back"
cp ./nginx/conf.d/${SERVICE}.nginx.conf ./nginx/conf.d/${SERVICE}.conf.back 2>/dev/null

echo "Checking nginx config for next service"
docker exec -e NEXT_SERVICE=$NEXT_SERVICE \
-i nginx envsubst '$NEXT_SERVICE' < ./nginx/conf.d/${SERVICE}.conf.template > ./nginx/conf.d/${SERVICE}.nginx.conf \
&& docker compose exec -T nginx nginx -g 'daemon off; master_process on;' -t
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