from kafka import KafkaConsumer
from kafka.errors import NoBrokersAvailable
from typing import AnyStr, Callable

import time
import json

KAFKA_CONNECT_WAIT = 5


class MessageConsumer:
    def __init__(self, topic: AnyStr, broker: AnyStr):
        """Constructs the MessageConsumer

        Args:
            topic (AnyStr): name of the Kafka topic to consume from
            broker (AnyStr): address of the broker
        """

        self._topic = topic
        self._consumer = self.connect_consumer(broker)

    def connect_consumer(self, broker: AnyStr) -> KafkaConsumer:
        """Connects to the Kafka broker. Will keep re-trying indefinitely
        until the connection is successful. 

        Args:
            broker (str): hostname and port of the Kafka broker

        Returns:
            KafkaConsumer: the connected consumer
        """

        print(f'Connecting consumer to {broker}...')
        while True:
            try:
                consumer = KafkaConsumer(
                    bootstrap_servers=broker,
                    value_deserializer=lambda v: json.loads(v.decode('utf-8'))
                )

                consumer.subscribe(self._topic)
                break
            except NoBrokersAvailable:
                print(f'Kafka broker {broker} appears not to be running yet, '
                      f'waiting {KAFKA_CONNECT_WAIT} seconds before reconnecting')

                # Add an intential delay such that we do not hammer the broker
                # when it is not fully initialized yet.
                time.sleep(KAFKA_CONNECT_WAIT)

        print('.. consumer connected.')
        return consumer

    def start_blocking(self, handler: Callable):
        """Start consuming message from the Kafka topic. For each message the
        specified handler will be called, passing the message value as the first
        parameter.

        Args:
            handler (Callable): the handler to be called for each message
        """
        
        for msg in self._consumer:
            handler(msg.topic, msg.value)
