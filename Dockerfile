FROM node AS ui

WORKDIR ./app

COPY ./ui .

RUN rm -r ./dist

RUN npm ci
RUN npm run build

RUN mv dist/index.html dist/main.html

FROM golang:1.20.2-alpine

WORKDIR ./app

COPY . .

RUN go mod download

COPY --from=ui /app/dist ./ui/dist

RUN ls ./ui/dist/
RUN go build -o mongoly

EXPOSE 8080:8080/tcp

CMD [ "./mongoly" ]

