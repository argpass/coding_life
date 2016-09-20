#include <string.h>
#include "unistd.h"
#include "time.h"
#include "netinet/in.h"
#include "sys/socket.h"
#include "lib/io.h"
#include "arpa/inet.h"
#include "stdio.h"


/** 开启tcp服务，向每一个client发送一句话和时间戳
 */
int main(int argc, char **argv){
    int srv_fd, client_fd;
    struct sockaddr_in srv_address, client_address;
    socklen_t len;
    int backlog = 1024;
    size_t n_read;
    char w_buff[1024];
    char r_buff[1024];

    // Initialize a socket
    srv_fd = socket(AF_INET, SOCK_STREAM, 0);
    srv_address.sin_family = AF_INET;
    inet_aton("127.0.0.1", &srv_address.sin_addr);
    srv_address.sin_port = htons(9000);
    // Bind address to socket
    bind(srv_fd, (struct sockaddr *) &srv_address, sizeof(srv_address));
    // Listening forever
    listen(srv_fd, backlog);
    printf("start accept at addr=> %s:%d\n", inet_ntoa(srv_address.sin_addr), ntohs(srv_address.sin_port));

    for (;;){
        // Accept connection
        printf("start wait...\n");
        client_fd = accept(srv_fd, (struct sockaddr *) &client_address, &len);
        printf("Got connection form %s:%d\n", inet_ntoa(client_address.sin_addr), ntohs(client_address.sin_port));
        // Read data from client
        n_read = Read(client_fd, r_buff, 1024);
        printf("Got message(%d):%s\n", (int) n_read, r_buff);
        sprintf(w_buff, "hello from server and time is:%d\n", time(NULL));
        // Write message to client
        Write(client_fd, w_buff, strlen(w_buff));
        // Close client connection
        shutdown(client_fd, SHUT_RDWR);
        printf("done\n");
    }
}
