#!/usr/bin/env bash
echo "Starting deploy ..."

BACKEND_PORT=$(grep 'addr' $CONFIG_PATH | awk -F'addr: ' '{gsub(/[:"]/,"",$2); print $2}')
if [ -z "$BACKEND_PORT" ]; then
    BACKEND_PORT="3333"
fi
echo "BACKEND_PORT is: $BACKEND_PORT"

if grep -q "proxy_pass http://blue-backend" './nginx/conf.d/nginx.conf'
then
    CURRENT_BACKEND="blue-backend"
    NEXT_BACKEND="green-backend"
    echo "$CURRENT_BACKEND backend in proccess"
else 
    CURRENT_BACKEND="green-backend"
    NEXT_BACKEND="blue-backend"
    echo "$CURRENT_BACKEND in proccess"
fi

echo "Removing old container ..."
docker compose rm -f -s -v $NEXT_BACKEND
printf "%s\n" "Done"

echo "Starting $NEXT_BACKEND"
docker compose up -d $NEXT_BACKEND
echo "Waiting 5 sec"
sleep 5

echo "Renaming current nginx.conf in nginx.conf.back"
cp ./nginx/conf.d/nginx.conf ./nginx/conf.d/nginx.conf.back 2>/dev/null

echo "Checking nginx config for next backend"
docker exec -e NEXT_BACKEND=$NEXT_BACKEND -e BACKEND_PORT=$BACKEND_PORT \
-i nginx envsubst '$NEXT_BACKEND,$BACKEND_PORT' < ./nginx/conf.d/nginx.conf.template > ./nginx/conf.d/nginx.conf \
&& docker compose exec -T nginx nginx -g 'daemon off; master_process on;' -t
rv=$?
if [ $rv != 0 ]; then
    cp ./nginx/conf.d/nginx.conf.back ./nginx/conf.d/nginx.conf 2>/dev/null
    echo "Checking failed with exit code: $rv"
    echo "Aborting..."
    exit 1
fi

echo "Reloading nginx with $NEXT_BACKEND"
docker compose exec -T nginx nginx -g 'daemon off; master_process on;' -s reload
rv=$?
if [ $rv != 0 ]; then
    echo "Reloading is failed with exit code: $rv"
    echo "Aborting..."
    cp ./nginx/conf.d/nginx.conf.back ./nginx/conf.d/nginx.conf 2>/dev/null
    exit 1
else    
    echo "Nginx reloaded"
fi

echo "Waiting 5 sec..."
sleep 5

echo "Testing minibank..."
curl -s localhost/v1/health | grep "Service is healthy"
rv=$?
if [ $rv != 0 ]; then
    echo "Testing is failed with exit code: $rv"
    echo "Aborting..."
    cp ./nginx/conf.d/nginx.conf.back ./nginx/conf.d/nginx.conf 2>/dev/null
    exit 1
else    
    echo "Testing is OK"
fi

echo "Deleting old container: $CURRENT_BACKEND"
docker compose rm -f -s -v $CURRENT_BACKEND

echo "Blue-Green deployment is done!"