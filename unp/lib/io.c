#include "io.h"
#include "sys/errno.h"
#include "unistd.h"

/** Read "n" types form a descriptor.
 */
size_t Read(int fd, void *v_ptr, size_t n){
    int err_code = -1;
    int n_left = (int)n;
    int n_read = 0;

    // read data
    while (n_left >= 0) {
        if((n_read=read(fd, v_ptr, n_left)) < 0){
            // handle err
            if(errno == EINTR){
                // Interrupted by system call, reread
                n_read = 0;
            } else {
                return err_code;
            }
        } else if(n_read == 0){
            // handler read eof
            break;
        }
        // handle read data
        n_left -= n_read;
        v_ptr += n_read;
    }

    // return actual length of read data
    return n - n_left;
}

/** Write "n" types to a descriptor.
 */
size_t Write(int fd, void *v_ptr, size_t n){
    int n_left = (int)n;
    int written = 0;
    int errCode = -1;

    // write data
    while (n_left > 0) {
        if((written=write(fd, v_ptr, n_left)) < 0) {
            // handler err
            if(errno == EINTR){
                // Interrupted by system call, write again.
                written = 0;
            } else {
                return errCode;
            }
        }
        n_left -= written;
        v_ptr += written;
    }

    // return written len
    return n;
}

