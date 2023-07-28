#ifndef FOO_HPP
#include <string>


class cxxFoo {
public:
  int a;
  cxxFoo(int _a):a(_a){};
  ~cxxFoo(){};
  void Print(std::string str);
};

#endif
