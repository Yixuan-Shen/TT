#include <cpprest/http_listener.h>
#include <cpprest/json.h>

using namespace web;
using namespace web::http;
using namespace web::http::experimental::listener;

void handle_get(http_request request)
{
    // Respond with a JSON object
    json::value response;
    response[U("message")] = json::value::string(U("Hello, World!"));

    request.reply(status_codes::OK, response);
}

int main()
{
    // Create HTTP listener at http://localhost:8080/todos
    http_listener listener(U("http://localhost:8080/todos"));
    listener.support(methods::GET, handle_get);

    try
    {
        listener.open().then([]() { std::wcout << L"Listening on http://localhost:8080/todos" << std::endl; }).wait();

        // Keep the program running to listen for incoming requests
        while (true)
        {
            std::this_thread::sleep_for(std::chrono::seconds(1));
        }
    }
    catch (const std::exception& ex)
    {
        std::wcout << ex.what() << std::endl;
    }

    return 0;
}
