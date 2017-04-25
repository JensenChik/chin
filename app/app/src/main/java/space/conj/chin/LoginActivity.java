package space.conj.chin;

import android.app.ProgressDialog;
import android.content.Intent;
import android.content.SharedPreferences;
import android.os.Bundle;
import android.os.Looper;
import android.preference.PreferenceManager;
import android.support.annotation.Nullable;
import android.support.v7.app.ActionBar;
import android.support.v7.app.AppCompatActivity;
import android.util.Log;
import android.view.View;
import android.widget.Button;
import android.widget.CheckBox;
import android.widget.EditText;
import android.widget.Toast;

import com.squareup.okhttp.Callback;
import com.squareup.okhttp.FormEncodingBuilder;
import com.squareup.okhttp.OkHttpClient;
import com.squareup.okhttp.Request;
import com.squareup.okhttp.Response;

import java.io.IOException;

import space.conj.chin.tools.RequestClient;

/**
 * Created by hit-s on 2017/4/22.
 */
public class LoginActivity extends AppCompatActivity {

    private Button login;
    private EditText userName;
    private EditText password;
    private CheckBox rememberPassword;
    private SharedPreferences pref;
    private SharedPreferences.Editor editor;
    private boolean isRemember;

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
        rememberPassword = (CheckBox) findViewById(R.id.remember_password);
        pref = PreferenceManager.getDefaultSharedPreferences(this);
        editor = pref.edit();
        isRemember = pref.getBoolean("remember_password", false);


        if (isRemember) {
            userName.setText(pref.getString("user_name", ""));
            password.setText(pref.getString("password", ""));
            rememberPassword.setChecked(true);
        }

        login.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                final ProgressDialog progressDialog = new ProgressDialog(LoginActivity.this);
                progressDialog.setTitle("正在登陆中");
                progressDialog.setMessage("耐心等待...");
                progressDialog.show();

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
                            if (rememberPassword.isChecked()) {
                                editor.putBoolean("remember_password", true);
                                editor.putString("user_name", userName.getText().toString());
                                editor.putString("password", password.getText().toString());
                            } else {
                                editor.clear();
                            }
                            editor.apply();
                            startActivity(new Intent(LoginActivity.this, ListTaskActivity.class));
                            progressDialog.dismiss();
                            finish();
                        } else {
                            Looper.prepare();
                            editor.clear();
                            editor.apply();
                            Toast.makeText(LoginActivity.this, "登陆账号密码有误", Toast.LENGTH_SHORT).show();
                            Looper.loop();
                        }
                    }
                });
                progressDialog.hide();
            }
        });

        if (isRemember) {
            login.performClick();
        }


    }
}
