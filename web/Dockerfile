# Стадия сборки
FROM node:20.5.0 AS build
WORKDIR /app
COPY ./web/package*.json ./
COPY ./web/tsconfig.json ./
COPY ./web/src /app/src
COPY ./web/public /app/public
RUN npm install

ARG public_url
ARG react_app_url
ENV PUBLIC_URL=$public_url
ENV REACT_APP_URL=$react_app_url
RUN echo "$PUBLIC_URL"
RUN echo "$REACT_APP_URL"
#RUN PUBLIC_URL="http://localhost" REACT_APP_URL="http://localhost/api/v1" npm run docker-build
RUN PUBLIC_URL="$PUBLIC_URL" REACT_APP_URL="$REACT_APP_URL" npm run docker-build

# Стадия развертывания
FROM nginx:1.25.2-alpine
COPY --from=build /app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
