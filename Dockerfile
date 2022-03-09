FROM node:14-alpine

VOLUME [ "/app/data" ]

WORKDIR /app
COPY package.json package-lock.json ./
RUN npm install

COPY . ./
CMD [ "node", "./app.js" ]
