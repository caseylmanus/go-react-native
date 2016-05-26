package com.goreact;

import android.widget.Toast;
import android.content.ContextWrapper;

import com.facebook.react.ReactActivity;
import com.facebook.react.ReactPackage;
import com.facebook.react.shell.MainReactPackage;

import java.io.File;
import java.util.Arrays;
import java.util.List;


import go.demo.Demo;

public class MainActivity extends ReactActivity {

    @Override
    protected void onResume() {
        super.onResume();
        File dir = new ContextWrapper(this).getFilesDir();
        Demo.WriteFile(dir.getAbsolutePath());
        //new Thread(new Runnable() {
        //    @Override
        //    public void run() {
                Demo.StartListening();
        //    }
        //}).start();
    }

    /**
     * Returns the name of the main component registered from JavaScript.
     * This is used to schedule rendering of the component.
     */
    @Override
    protected String getMainComponentName() {
        return "GoReact";
    }

    /**
     * Returns whether dev mode should be enabled.
     * This enables e.g. the dev menu.
     */
    @Override
    protected boolean getUseDeveloperSupport() {
        return BuildConfig.DEBUG;
    }

    /**
     * A list of packages used by the app. If the app uses additional views
     * or modules besides the default ones, add more packages here.
     */
    @Override
    protected List<ReactPackage> getPackages() {
        return Arrays.<ReactPackage>asList(
            new MainReactPackage(),
            new GoPackage()
        );
    }
}
