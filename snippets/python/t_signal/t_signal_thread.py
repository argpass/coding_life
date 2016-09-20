#!coding: utf-8
"""
以下演示使用信号进行子线程与主线程得交互
"""
__author__ = 'zkchen'
import time
import os
import signal
import threading


def handle_sig(sig_id, frame):
    print "get sig ", sig_id
    if sig_id == signal.SIGUSR2:
        print "exit ..."
        os._exit(0)


def sub_thread():
    for i in xrange(5):
        os.kill(os.getpid(), signal.SIGUSR1)
        time.sleep(1)
    os.kill(os.getpid(), signal.SIGUSR2)
    print "sub_thread done"


if __name__ == "__main__":
    signal.signal(signal.SIGUSR1, handle_sig)
    signal.signal(signal.SIGUSR2, handle_sig)

    threading._start_new_thread(sub_thread, (), {})
    while True:
        signal.pause()
