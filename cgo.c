#include <stdio.h>
#include <unistd.h>
#include "goclient.h"
#include <string.h>
#include <pthread.h>
#include <sys/time.h>

/**
 * allocate pointer, pass it to go routine and use it
 */
int main()
{
	GoChan ptr = NULL;
	int i;

	printf("1. ptr %p\n", ptr);
	GetChanPointer(&ptr);
	printf("2. ptr %p\n", ptr);
	GetChanPointer(&ptr);
	printf("3. ptr %p\n", ptr);
	GetChanPointer(&ptr);
	printf("4. ptr %p\n", ptr);
	GetChanPointer(&ptr);

	return 1;
}
