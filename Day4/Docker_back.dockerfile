FROM node:alpine
EXPOSE 8080
COPY back .
WORKDIR /back
RUN npm install
RUN npm start