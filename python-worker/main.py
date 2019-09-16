# coding: utf-8

import time
import logging
import sys

if __name__ == '__main__':

    # logging must setting to sys.stdout
    logger = logging.getLogger()
    logger.setLevel(logging.INFO)
    handler = logging.StreamHandler(sys.stdout)
    handler.setLevel(logging.INFO)
    logger.addHandler(handler)
    
    i = 0

    while True:
        logger.info("count: %d", i)

        i += 1
        if i >= 3:
            break

        time.sleep(1)
