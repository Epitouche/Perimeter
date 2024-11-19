FROM alpine

WORKDIR /app

RUN apk add erlang rebar3 gleam

RUN gleam --version

COPY ./project .

CMD [ "gleam", "run", "-m", "lustre/dev", "start" ]