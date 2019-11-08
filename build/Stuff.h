///-----------------------------------------------------------------------------
/// @file Stuff.h
///
/// Copyright (C) Circle Cardiovascular Imaging 2019/11/08
///
/// Author: Evan Loughlin
///
/// Description: 
///-----------------------------------------------------------------------------


#ifndef 
#define 

#include "I_Stuff.h"

#include <QObject>


class Stuff : public QObject, public I_Stuff
{
Q_OBJECT

 public:
    explicit Stuff();
    virtual ~Stuff();

 public:

 
 signals:


 private:

};

#endif //