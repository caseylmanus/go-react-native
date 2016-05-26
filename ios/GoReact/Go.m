//
//  Go.m
//  GoReact
//
//  Created by Casey Manus on 5/25/16.
//  Copyright © 2016 Facebook. All rights reserved.
//

#import "Go.h"
#import "RCTBridge.h"
#import "RCTEventDispatcher.h"

@implementation Go

@synthesize bridge = _bridge;

RCT_EXPORT_MODULE();

RCT_EXPORT_METHOD(sayHello:(RCTResponseSenderBlock)success){
  success(@[GoDemoHelloWorld(self)]);
}
-(void)sendEvent:(NSString*)channel message:(NSString*)message{
  [self.bridge.eventDispatcher sendDeviceEventWithName:channel body:message];
}

@end
