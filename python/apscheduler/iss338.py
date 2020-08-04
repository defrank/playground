#!/usr/bin/env python3
"""
GitHub Issue:
* https://github.com/agronholm/apscheduler/issues/338

"""
import time

from apscheduler.events import EVENT_JOB_EXECUTED
from apscheduler.schedulers.blocking import BlockingScheduler


def hello():
    print('Hello, World!')
    time.sleep(.1)


def shutdown(event):
    if scheduler.running:
        scheduler.shutdown(wait=False)


if __name__ == '__main__':
    scheduler = BlockingScheduler()
    scheduler.add_listener(shutdown, EVENT_JOB_EXECUTED)
    scheduler.add_job(hello)

    try:
        scheduler.start()
    except KeyboardInterrupt:
        scheduler.shutdown()
