//
//  Go.m
//  GoReact
//
//  Created by Casey Manus on 5/25/16.
//  Copyright © 2016 Facebook. All rights reserved.
//

#import "Go.h"

@implementation Go

RCT_EXPORT_MODULE();

RCT_EXPORT_METHOD(sayHello:(RCTResponseSenderBlock)success){
  success(@[GoDemoHelloWorld()]);
}

@end
