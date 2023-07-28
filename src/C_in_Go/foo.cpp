#include <iostream>
#include <cpprest/http_listener.h>
#include <cpprest/json.h>


using namespace web;
using namespace web::http;
using namespace web::http::experimental::listener;

void handle_get(http_request request) {
    json::value response;
    response[U("message")] = json::value::string(U("Hello from C++ REST API!"));

    request.reply(status_codes::OK, response);
}

int main() {
    http_listener listener(U("http://localhost:8888"));
    listener.support(methods::GET, handle_get);

    try {
        listener.open().wait();
        ucout << U("C++ REST API listening on port 8888...") << std::endl;
        std::getchar(); // Wait for a keypress to stop the server
        listener.close().wait();
    } catch (const std::exception &e) {
        ucout << U("Error: ") << e.what() << std::endl;
    }

    return 0;
}