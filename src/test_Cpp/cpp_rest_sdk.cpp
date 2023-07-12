#include <cpprest\http_client.h>
#include <cpprest\filestream.h>
using  namespace  utility;
using  namespace  web;
using  namespace  web::http;
using  namespace  web::http::client;
using  namespace  concurrency;
 
void  TestRequest()
{
     auto  fileStream = std::make_shared<concurrency::streams::ostream>();
     pplx::task< void > requestTask = concurrency::streams::fstream::open_ostream(U( "result.html" )).then([=](concurrency::streams::ostream 
 
outFile){
         *fileStream = outFile;
 
         http_client client(U( "http://www.bing.com/" ));
         uri_builder builder(U( "/search" ));
         builder.append_query(U( "q" ), U( "Casablanca CodePlex" ));
 
         return  client.request(methods::GET, builder.to_string());
     })
     .then([=](http_response response)
     {
         return  response.body().read_to_end(fileStream->streambuf());
     }).then([=]( size_t  len){
         return  fileStream->close();
     });
 
     try
     {
         requestTask.wait();
     }
     catch  ( const  std::exception& e)
     {
         cout << e.what() << endl;
     }
}


void TestRequest()
{
    auto fileStream = std::make_shared<concurrency::streams::ostream>();
    concurrency::streams::ostream outFile = concurrency::streams::fstream::open_ostream(U("result11.html")).get();
    *fileStream = outFile;

    http_client client(L"http://www.bing.com/");
    uri_builder builder(L"/search");
    builder.append_query(L"q", L"Casablanca CodePlex");

    http_response response = client.request(methods::GET, builder.to_string()).get();
    response.body().read_to_end(fileStream->streambuf()).get();
    fileStream->close().get();
}

uri_builder builder;
builder.append_path(L"search"); //添加URL
builder.append_query(L"q", L"Casablanca CodePlex"); //添加url参数

client.request(methods::GET, builder.to_string()).get();


uri_builder builder;
builder.append_path(L"/test");

json::value obj;
obj[L"Count"] = json::value::number(6);
obj[L"Version"] = json::value::string(L"1.0");
client.request(methods::POST, builder.to_string(), obj.serialize(), L"application/json");

wchar_t buf[48] = {};
http_response response = client.request(methods::POST, builder.to_string(), buf/*L""*/, L"application/octet-stream").get();

if (response.status_code() == status_codes::OK)
{
    try
    {
        result = true;
        const json::value& jv = response.extract_json().get();
        const web::json::object& jobj = jv.as_object();
        auto result = jobj.at(L"result").as_string();
        auto access_code = result.as_object().at(L"error_code").as_string();
        wcout << result<<" "<< access_code << endl;
    }
    catch (const std::exception& e)
    {
        cout << e.what() << endl;
    }
}

wcout.imbue(locale("chs"));//本地化

http_client_config config;
config.set_timeout(utility::seconds(90)); //设置为90秒超时
http_client client(URL, config);