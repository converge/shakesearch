# todo: create a production container with alpine / less privileges
FROM node:18 AS build

WORKDIR /opt/frontend

ARG API_URL

COPY ./package.json ./yarn.lock ./
COPY ./public ./public
COPY ./src ./src

RUN yarn install

RUN yarn build

#CMD sh -c "sleep 1d"

FROM nginx:1.23-alpine

WORKDIR /usr/share/nginx/html

EXPOSE 80

COPY ./deployments/nginx/default.conf /etc/nginx/conf.d/default.conf
COPY --from=build /opt/frontend/dist .

# run nginx with global directives and daemon off
ENTRYPOINT ["nginx", "-g", "daemon off;"]
