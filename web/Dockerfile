# Стадия сборки
FROM node:20.5.0 AS build
WORKDIR /app
COPY ./web/package*.json ./
COPY ./web/tsconfig.json ./
COPY ./web/src /app/src
COPY ./web/public /app/public
RUN npm install
RUN npm run build

# Стадия развертывания
FROM nginx:1.25.2-alpine
COPY --from=build /app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
