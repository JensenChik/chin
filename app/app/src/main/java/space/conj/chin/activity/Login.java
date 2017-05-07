package space.conj.chin.activity;

import android.content.Intent;
import android.content.SharedPreferences;
import android.os.Bundle;
import android.os.Looper;
import android.preference.PreferenceManager;
import android.support.annotation.Nullable;
import android.support.v7.app.ActionBar;
import android.support.v7.app.AppCompatActivity;
import android.view.View;
import android.widget.Button;
import android.widget.CheckBox;
import android.widget.EditText;
import android.widget.Toast;


import space.conj.chin.R;
import space.conj.chin.tools.RequestClient;

/**
 * Created by hit-s on 2017/4/22.
 */
public class Login extends AppCompatActivity {

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
                login();
            }
        });

        if (isRemember) {
            login.performClick();
        }
    }

    private void login() {
        new Thread(new Runnable() {
            @Override
            public void run() {
                String userNameString = userName.getText().toString();
                String passwordString = password.getText().toString();

                boolean loginSuccess = RequestClient.login(userNameString, passwordString);
                if (loginSuccess) {
                    if (rememberPassword.isChecked()) {
                        editor.putBoolean("remember_password", true);
                        editor.putString("user_name", userNameString);
                        editor.putString("password", passwordString);
                    } else {
                        editor.clear();
                    }
                    editor.apply();
                    startActivity(new Intent(Login.this, ListTask.class));
                    finish();
                } else {
                    Looper.prepare();
                    editor.clear();
                    editor.apply();
                    Toast.makeText(Login.this, "登陆账号密码有误", Toast.LENGTH_SHORT).show();
                    Looper.loop();
                }
            }
        }).start();
    }
}
