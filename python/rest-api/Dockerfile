# syntax=docker/dockerfile:1

FROM python:alpine

WORKDIR /python-docker

COPY requirements.txt requirements.txt
RUN pip3 install -r requirements.txt

COPY . .
RUN ls
USER 10016
CMD [ "flask", "run", "--host=0.0.0.0"]
