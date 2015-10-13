cmake_minimum_required(VERSION 3.3)
project(alg)

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++0x")

set(SOURCE_FILES tests/t_dblist.c tests/t_dblist.h)
add_executable(t_dblist ${SOURCE_FILES})