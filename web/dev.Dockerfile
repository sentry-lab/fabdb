FROM node:lts-alpine

WORKDIR /srv

COPY package.json package-lock.json .

RUN npm install

COPY . .

CMD ["npm", "run", "dev"]

EXPOSE 3000
