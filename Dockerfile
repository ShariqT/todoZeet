FROM public.ecr.aws/zeet/lambdahandler

WORKDIR /app

COPY . .

RUN go mod download

RUN go build

EXPOSE 3000

ENV DATABASE=localhost

CMD ./todoZeet