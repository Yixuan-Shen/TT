#include <windows.h>
#include <string>
#include <stdio.h>

using std::string;

#pragma comment(lib,"ws2_32.lib")


HINSTANCE hInst;
WSADATA wsaData;
void mParseUrl(char *mUrl, string &serverName, string &filepath, string &filename);
SOCKET connectToServer(char *szServerName, WORD portNum);
int getHeaderLength(char *content);
char *readUrl2(char *szUrl, long &bytesReturnedOut, char **headerOut);


int main()
{
    const int bufLen = 1024;
    char *szUrl = "http://stackoverflow.com";
    long fileSize;
    char *memBuffer, *headerBuffer;
    FILE *fp;

    memBuffer = headerBuffer = NULL;

    if ( WSAStartup(0x101, &wsaData) != 0)
        return -1;
    memBuffer = readUrl2(szUrl, fileSize, &headerBuffer);
    printf("returned from readUrl\n");
    printf("data returned:\n%s", memBuffer);
    if (fileSize != 0)
    {
        printf("Got some data\n");
        fp = fopen("downloaded.file", "wb");
        fwrite(memBuffer, 1, fileSize, fp);
        fclose(fp);
         delete(memBuffer);
        delete(headerBuffer);
    }

    WSACleanup();
    return 0;
}

void mParseUrl(char *mUrl, string &serverName, string &filepath, string &filename)
{
    string::size_type n;
    string url = mUrl;

    if (url.substr(0,7) == "http://")
        url.erase(0,7);

    if (url.substr(0,8) == "https://")
        url.erase(0,8);

    n = url.find('/');
    if (n != string::npos)
    {
        serverName = url.substr(0,n);
        filepath = url.substr(n);
        n = filepath.rfind('/');
        filename = filepath.substr(n+1);
    }

    else
    {
        serverName = url;
        filepath = "/";
        filename = "";
    }
}