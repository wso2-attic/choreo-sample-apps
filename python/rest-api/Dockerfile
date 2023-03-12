# syntax=docker/dockerfile:1

FROM python:alpine

WORKDIR /python-docker

COPY requirements.txt requirements.txt
RUN pip3 install -r requirements.txt

COPY . .
RUN ls
# Create a new user with UID 10016
RUN useradd -u 10016 -m choreouser
USER 10016
EXPOSE 5000
CMD [ "flask", "run", "--host=0.0.0.0"]
