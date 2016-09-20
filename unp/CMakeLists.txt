cmake_minimum_required(VERSION 3.3)
project(unp)

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++0x")

set(byteOrder_src byte_order.c)
add_executable(typeOrder ${byteOrder_src})

set(daytime_tcp_srv_src daytime_tcp_srv.c lib/io.c)
add_executable(daytime_tcp_srv ${daytime_tcp_srv_src})

set(clisrv.src clisrv.c)
add_executable(clisrv ${clisrv.src})