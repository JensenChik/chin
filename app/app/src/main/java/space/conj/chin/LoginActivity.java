package space.conj.chin;

import android.content.Intent;
import android.os.Bundle;
import android.os.Looper;
import android.support.annotation.Nullable;
import android.support.v7.app.ActionBar;
import android.support.v7.app.AppCompatActivity;
import android.util.Log;
import android.view.View;
import android.widget.Button;
import android.widget.EditText;
import android.widget.Toast;

import com.squareup.okhttp.Callback;
import com.squareup.okhttp.FormEncodingBuilder;
import com.squareup.okhttp.OkHttpClient;
import com.squareup.okhttp.Request;
import com.squareup.okhttp.Response;

import java.io.IOException;
import java.net.CookieManager;

import space.conj.chin.tools.RequestClient;

/**
 * Created by hit-s on 2017/4/22.
 */
public class LoginActivity extends AppCompatActivity {

    Button login;
    EditText userName;
    EditText password;

    @Override
    protected void onCreate(@Nullable Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        ActionBar bar = getSupportActionBar();
        if (bar != null) {
            bar.hide();
        }
        setContentView(R.layout.login);
        login = (Button) findViewById(R.id.login);
        userName = (EditText) findViewById(R.id.user_name);
        password = (EditText) findViewById(R.id.password);

        login.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {

                OkHttpClient client = RequestClient.getInstance();

                FormEncodingBuilder builder = new FormEncodingBuilder();
                builder.add("user_name", userName.getText().toString());
                builder.add("password", password.getText().toString());
                final Request request = new Request.Builder()
                        .url("http://chin.nazgrim.com/login")
                        .post(builder.build())
                        .build();
                client.newCall(request).enqueue(new Callback() {
                    public void onFailure(Request request, IOException e) {

                    }

                    @Override
                    public void onResponse(Response response) throws IOException {
                        if (response.headers().names().contains("Set-Cookie")) {
                            Log.i("LOGIN COOKIE", response.header("Set-Cookie"));
                            startActivity(new Intent(LoginActivity.this, TaskActivity.class));
                            finish();
                        } else {
                            Looper.prepare();
                            Toast.makeText(LoginActivity.this, "登陆账号密码有误", Toast.LENGTH_SHORT).show();
                            Looper.loop();
                        }

                    }
                });

            }
        });


    }
}
