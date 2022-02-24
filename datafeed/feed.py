from KafkaProducer import MessageProducer

import time
import random
import glob
import os

KAFKA_BROKER = 'kafka:9092'
KAFKA_TOPIC = 'picture'
PICTURE_PATH = '/images'

WEIGHT_RANGE = (20, 5000)
RANDOM_WAIT_RANGE_SECONDS = (2, 5)


class Feed:
    def __init__(self):
        self._kafka_producer = MessageProducer(KAFKA_TOPIC, KAFKA_BROKER)

    def start(self):
        while True:
            self._send_message()
            self._random_wait()

    def _get_random_picture_id(self):
        pictures = list(glob.glob(f'{PICTURE_PATH}/*.jpg'))

        picture_id = None
        if pictures:
            random.shuffle(pictures)
            picture_id, _ = os.path.splitext(os.path.basename(pictures[0]))

        return picture_id

    def _get_random_weight(self):
        return random.randrange(*WEIGHT_RANGE)

    def _send_message(self):
        print('Sending random message...')
        msg = {
            'picture_id': self._get_random_picture_id(),
            'weight': self._get_random_weight()
        }

        self._kafka_producer.send(msg)

    def _random_wait(self):
        wait_seconds = random.randrange(*RANDOM_WAIT_RANGE_SECONDS)
        print(f'Waiting {wait_seconds} seconds...')
        time.sleep(wait_seconds)


if __name__ == '__main__':
    feed = Feed()
    feed.start()
