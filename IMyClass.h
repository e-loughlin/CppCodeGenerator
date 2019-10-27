/* 
    Copyright 2019 Evan Loughlin 
*/

#ifndef IMYCLASS_H
#define IMYCLASS_H

class IMyClass 
{
 public:
    virtual ~IMyClass(){}

 public:
    virtual void doSomething() = 0;
    virtual QString doSomethingElse(std::string name) = 0;

 public signals:

};

#endif
