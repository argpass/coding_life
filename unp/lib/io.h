#ifndef UNP_IO_H
#define UNP_IO_H

#include "unistd.h"

/** Read "n" types form a descriptor.
 */
size_t Read(int fd, void *v_ptr, size_t n);


/** Write "n" types to a descriptor.
 */
size_t Write(int fd, void *v_ptr, size_t n);


#endif //UNP_IO_H
