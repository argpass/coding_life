cmake_minimum_required(VERSION 3.3)
project(alg)

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++0x")
include_directories(src)

set(SOURCE_FILES tests/t_dblist.c src/dblist.c src/dblist.h)
add_executable(t_dblist ${SOURCE_FILES})

set(src_dstring tests/t_dstring.c src/dstring.c)
add_executable(t_dstring ${src_dstring})