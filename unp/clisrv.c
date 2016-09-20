#include <sys/wait.h>
#include <string.h>
#include <arpa/inet.h>
#include <sys/errno.h>
#include "stdio.h"
#include "unistd.h"


int cli_str(int fd);
void handle_sig_child(int sig);

/** 多进程回文服务器
 */
int main(){
    struct sockaddr_in srv_address, client_address;
    int srv_fd, client_fd;
    socklen_t len;
    int pid;

    srv_fd = socket(AF_INET, SOCK_STREAM, 0);
    // TODO[note]:设置了重用标识，程序关闭后依旧不能立刻再次绑定该端口,需要等待一会儿
    setsockopt(srv_fd, SOL_SOCKET, SO_REUSEADDR|SO_REUSEPORT, &len, sizeof(len));
    srv_address.sin_family = AF_INET;
    srv_address.sin_port = htons(9004);
    if(inet_aton("127.0.0.1", &srv_address.sin_addr) != 0){
        printf("Error aton\n");
    }
    if(0 != bind(srv_fd, (struct sockaddr*) &srv_address, sizeof(srv_address))){
        printf("Error when bind\n");
    }
    listen(srv_fd, 1000);

    // Register signal handler to handle SIG_CHILD
    signal(SIGCHLD, handle_sig_child);

    for(;;){
        printf("start listen\n");
        client_fd = accept(srv_fd, (struct sockaddr*) &client_address, &len);
        // TODO[bug]:第一次无法得到连接的地址和端口，之后的连接均可以。
        printf("Got request from %s:%d\n", inet_ntoa(client_address.sin_addr), ntohs(client_address.sin_port));
        if((pid=fork())==0){
            // Handle connection in sub process
            close(srv_fd);
            cli_str(client_fd);
            close(client_fd);
            _exit(0);
        }
        close(client_fd);
    }
}


/** Handle client request.
 */
int cli_str(int fd){
    printf("start cli_str...");
    char r_buff[1024];
    ssize_t n_read;

    // 未处理系统调用终端的重启
    reread:
        while ((n_read=read(fd, r_buff, sizeof(r_buff))) > 0){
            write(fd, r_buff, (size_t)n_read);
            printf("recv:%s\n", r_buff);
            // 刷新缓冲区
            bzero(r_buff, sizeof(r_buff));
        }
    if(n_read == -1 && errno==EINTR)
        goto reread;
}


/** 处理子进程结束的信号
 */
void handle_sig_child(int sig){
    int status;
    pid_t pid;

    // TODO[note]:为何不使用wait函数?
    while ((pid=(waitpid(-1, &status, WNOHANG))) > 0){
        printf("child process %d terminate.\n", pid);
    }
    return;
}
