FROM node:alpine
EXPOSE 8080
COPY front .
WORKDIR /front
RUN npm install
RUN npm start