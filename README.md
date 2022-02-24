# Coding Assignment - Picture Registration

Consider an imaginary system where users can make picture registrations using an application. Each registration consists of a *picture* accompanied by a *weight* (in grams). For these registrations the goal is to gather some statistics which are described in [description of the assignment](#description-of-the-assignment) section. We have already partially implemented this application as a set of docker containers.

These have been combined into a docker-compose application:

* **Kafka** - A Kafka message message bus that can be used for communication between components.
* **Datafeed** - The feed component. It will occasionally produce a registration message (as if coming from the user) on the the Kafka topic named *picture*
* **Consumer** - Example skeleton component that will consume and output all messages produced on the *picture* topic. This example is implemented in both Python (consumer) and GoLang (go_consumer)
* **Images** - A directory containing the dataset of 100 images. Any message produced by the *datafeed* will reference one of those images.

An example message looks like this:

```json
{
    "picture_id": "000000243626",
    "weight": 1778
}
```

## Description of the assignment

The goal of this assignment is to finalize the docker-compose application and implement the following functionality:

* Create a web API (like REST or GraphQL) which can be used to retrieve:
    - Total number of registrations so far
    - Average weight so far
* Try to come up with an architecture that makes it easy to extend. Meaning that if we come up with a new type of statistics we can easily add the core logic to the existing system.

### Notes

* Make sure the results are demo-able (either through a cURL, Postman call or a simple web interface)
* Styling is **not** important
* Persistency is **not** important. If the docker-compose application is restarted it can with a clean slate.
* The boilerplate consumers have been built using Python and Golang, but to complete the assignment you can use any technology and programming language that has your preference

## Getting started

This assignment requires *docker* and *docker-compose* to be installed on your system.

To run the skeleton application open a terminal, go to the directory containing the files for this assignment and run the following command:

```sh
$ docker-compose up --build
```

If all goes well, after some time, you should see the following output:

```sh
...
kafka_1      | [2019-09-19 14:09:18,662] INFO Replica loaded for partition picture-0 with initial high watermark 0 (kafka.cluster.Replica)
kafka_1      | [2019-09-19 14:09:18,665] INFO [Partition picture-0 broker=1001] picture-0 starts at Leader Epoch 0 from offset 0. Previous Leader Epoch was: -1 (kafka.cluster.Partition)
datafeed_1   | .. producer connected.
datafeed_1   | Sending random message...
datafeed_1   | Waiting 3 seconds...
consumer_1   | Received a message on topic picture: {'picture_id': '000000419312', 'weight': 156}
datafeed_1   | Sending random message...
datafeed_1   | Waiting 4 seconds...
consumer_1   | Received a message on topic picture: {'picture_id': '000000460347', 'weight': 1635}
...
```

If you do not have this installed yet, use the following instructions:

#### Ubuntu

```sh
$ sudo apt-get update
$ sudo apt-get install docker.io docker-compose
```

#### Windows

Install the Docker Toolbox using the instructions on https://docs.docker.com/toolbox/toolbox_install_windows/

## Running from outside of the docker container

When developing it is often more convenient to run your code from outside of the Docker container. To support this, the Kafka container exposes two ports 9092 (for internal access) and 9094 (for external access and local development). When running locally you can configure the consumer to access port 9094 by setting the `KAFKA_BROKER` environment variable like this: 

```sh
# Python (from the consumer directory)
$ KAFKA_BROKER=localhost:9094 python consumer.py

# Or when using pipenv
$ KAFKA_BROKER=localhost:9094 pipenv run python consumer.py

# Golang (from the go_consumer directory)
KAFKA_BROKER=localhost:9094 go run main.go
```

## Troubleshooting

If for some reason the Kafka data gets corrupted, and sending / receiving of messages is no longer working then clear the Kafka database by running the following from the root of the codebase:

```sh
$ docker-compose rm
```
