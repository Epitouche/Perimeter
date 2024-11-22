ARG GLEAM_VERSION=v1.5.1

FROM ghcr.io/gleam-lang/gleam:${GLEAM_VERSION}-erlang-alpine

WORKDIR /app

COPY ./project .

# Install dependencies
RUN gleam deps download

RUN gleam build

EXPOSE 1234

CMD [ "gleam", "run", "-m", "lustre/dev", "start" ]