#!/usr/bin/env bash
echo "Starting deploy ..."

if grep -q "proxy_pass http://blue-backend" './nginx/conf.d/nginx.conf'
then
    export CURRENT_BACKEND="blue-backend"
    export NEXT_BACKEND="green-backend"
    echo "$CURRENT_BACKEND in proccess"
else 
    export CURRENT_BACKEND="green-backend"
    export NEXT_BACKEND="blue-backend"
    echo "$CURRENT_BACKEND in proccess"
fi

echo "Removing old container if it hasn't been removed..."
docker compose rm -f -s -v $NEXT_BACKEND
printf "%s\n" "Done"

echo "Migrating database ..."
docker compose up -d migrate
printf "%s\n" "Done...waiting 5 sec"
sleep 5

echo "Starting $NEXT_BACKEND"
docker compose up -d $NEXT_BACKEND
echo "Waiting 5 sec"
sleep 5

echo "Copying current nginx.conf to nginx.conf.back"
cp ./nginx/conf.d/nginx.conf ./nginx/conf.d/nginx.conf.back 2>/dev/null

echo "Checking nginx config for next backend"
docker exec -e NEXT_BACKEND=$NEXT_BACKEND \
-i nginx envsubst '$NEXT_BACKEND' < ./nginx/conf.d/nginx.conf.template > ./nginx/conf.d/nginx.conf \
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
curl -s http://localhost/api/v1/health | grep "Service is healthy"
rv=$?
if [ $rv != 0 ]; then
    echo "Testing is failed with exit code: $rv"
    echo "Aborting..."
    cp ./nginx/conf.d/nginx.conf ./nginx/conf.d/error.conf.back 2>/dev/null
    cp ./nginx/conf.d/nginx.conf.back ./nginx/conf.d/nginx.conf 2>/dev/null
    echo "Reloading nginx with $CURRENT_BACKEND"
    docker compose exec -T nginx nginx -g 'daemon off; master_process on;' -s reload
    echo "Nothing has happend,do not forget about migration..."
    exit 1
else    
    echo "Testing is OK"
fi

echo "Deleting old container: $CURRENT_BACKEND"
docker compose rm -f -s -v $CURRENT_BACKEND

echo "Blue-Green deployment is done!"