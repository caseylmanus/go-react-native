package com.goreact;
//import react native stuff
import com.facebook.react.bridge.NativeModule;
import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactContext;
import com.facebook.react.bridge.ReactContextBaseJavaModule;
import com.facebook.react.bridge.ReactMethod;
import com.facebook.react.bridge.Callback;

import java.util.Map;

//import Go Module
import go.demo.Demo;

public class GoModule extends ReactContextBaseJavaModule {

  @ReactMethod
  public void sayHello(Callback success){
      success.invoke(Demo.HelloWorld());
  } 
  @Override
  public String getName(){
      return "Go";
  }

  public GoModule(ReactApplicationContext reactContext) {
    super(reactContext);
  }
}