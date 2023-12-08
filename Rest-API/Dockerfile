FROM node:18.8.0

ENV PORT=80

WORKDIR /Rest-API

COPY ["package.json", "package-lock.json", "./"]

RUN npm install

RUN npm install -g nodemon

COPY . .

CMD ["nodemon", "app.js"]
