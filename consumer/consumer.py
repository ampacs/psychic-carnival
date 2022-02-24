from KafkaConsumer import MessageConsumer

import os
import time
import random
import glob
from typing import AnyStr, Callable

# The Kafka broker to connect to
# By default this is set to kafka:9092. If you are running this script from
# outside of the docker container, then please set the KAFKA_BROKER environment
# variable to "localhost:9094"
KAFKA_BROKER = os.environ.get('KAFKA_BROKER', 'kafka:9092')

# The topic to consume messages from
KAFKA_TOPIC = 'picture'


def message_handler(topic: AnyStr, msg: dict):
    """Called for every message that is consumed from the Kafka topic

    Args:
        topic (AnyStr): the topic the message is consumed from
        msg (dict): the message as published on the Kafka topic
    """    

    print(f'[Python] Received a message on topic {topic}: {msg}')


if __name__ == '__main__':
    # Connect to the Kafka bus and start consuming messages
    consumer = MessageConsumer(KAFKA_TOPIC, KAFKA_BROKER)
    consumer.start_blocking(message_handler)
