package com.goreact;
//import react native stuff
import android.util.Log;

import com.facebook.react.bridge.NativeModule;
import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactContext;
import com.facebook.react.bridge.ReactContextBaseJavaModule;
import com.facebook.react.bridge.ReactMethod;
import com.facebook.react.bridge.Callback;
import com.facebook.react.modules.core.DeviceEventManagerModule;

import java.util.Map;

//import Go Module
import demo.Demo;

public class GoModule extends ReactContextBaseJavaModule {

  @ReactMethod
  public void sayHello(Callback success){
      final ReactApplicationContext context = this.getReactApplicationContext();
      success.invoke(Demo.helloWorld(new demo.EventBus() {
        @Override
        public void sendEvent(String channel, String message) {
          Log.i(channel, message);
          context.getJSModule(DeviceEventManagerModule.RCTDeviceEventEmitter.class)
                .emit(channel, message);
        }
      }));
  } 
  @Override
  public String getName(){
      return "Go";
  }

  public GoModule(ReactApplicationContext reactContext) {
    super(reactContext);
  }
}
