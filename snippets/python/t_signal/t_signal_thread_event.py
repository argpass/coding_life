#!coding: utf-8
"""
对比使用signal.pause()等待信号以阻塞主线程，
使用event来阻塞主线程会遇到无法正常接收信号得问题，
以下就是错误用法，主线程等待event被set而阻塞，无法正常应答子线程发来得信号
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


def sub_thread(event):
    for i in xrange(5):
        os.kill(os.getpid(), signal.SIGUSR1)
        time.sleep(1)
    os.kill(os.getpid(), signal.SIGUSR2)
    event.set()
    print "sub_thread done"


if __name__ == "__main__":
    print os.getpid()
    signal.signal(signal.SIGUSR1, handle_sig)
    signal.signal(signal.SIGUSR2, handle_sig)

    main_event = threading.Event()
    threading._start_new_thread(sub_thread, (main_event, ), {})
    main_event.wait(None)
