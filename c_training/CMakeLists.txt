cmake_minimum_required(VERSION 3.3)
project(c_training)

set(SOURCE_FILES main.c)
add_executable(c_training ${SOURCE_FILES})


set(idea_py.src idea_py/main.c idea_py/libs/object.c idea_py/libs/int.c)
add_executable(idea_py ${idea_py.src})