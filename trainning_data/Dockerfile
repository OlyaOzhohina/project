FROM python:3.9-alpine

WORKDIR /app

RUN apk add --update make

RUN pip3 install pipenv

COPY ./Pipfile .
COPY ./Pipfile.lock .

RUN pipenv install --system --deploy

COPY . .

EXPOSE 3000

CMD ["make", "start"]