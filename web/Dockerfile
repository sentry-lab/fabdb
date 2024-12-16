FROM node:lts-alpine AS build

WORKDIR /app

COPY package.json package-lock.json .

RUN npm install

COPY --from=env .env .
COPY . .

RUN npm run build

FROM node:lts-alpine

RUN npm install -g http-server

WORKDIR /srv

COPY --from=build /app/dist /app/public .

CMD http-server ./ -p 3000 --corse
EXPOSE 3000
