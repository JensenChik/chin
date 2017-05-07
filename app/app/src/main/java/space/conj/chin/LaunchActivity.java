package space.conj.chin;

import android.content.Intent;
import android.os.Bundle;
import android.os.Handler;
import android.support.annotation.Nullable;
import android.support.v7.app.ActionBar;
import android.support.v7.app.AppCompatActivity;
import android.util.Log;

import com.squareup.okhttp.OkHttpClient;

import java.net.CookieManager;
import java.net.HttpCookie;
import java.util.List;

import space.conj.chin.tools.RequestClient;

/**
 * Created by hit-s on 2017/4/22.
 */
public class LaunchActivity extends AppCompatActivity {

    private OkHttpClient client;

    @Override
    protected void onCreate(@Nullable Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        ActionBar bar = getSupportActionBar();
        if (bar != null) {
            bar.hide();
        }
        setContentView(R.layout.launch);
        client = RequestClient.getInstance();
        Log.i("Home", "进入主页面");
        new Handler().postDelayed(new Runnable() {
            public void run() {
                if (RequestClient.hasCookieOf("chin.conj.space")) {
                    Log.i("LaunchActivity", "cookie已存在，直接跳转任务页");
                    startActivity(new Intent(LaunchActivity.this, ListTaskActivity.class));
                } else {
                    Log.i("LaunchActivity", "cookie未存在，跳转登陆页");
                    startActivity(new Intent(LaunchActivity.this, LoginActivity.class));
                }
                finish();
            }
        }, 500);
    }

}
