package space.conj.chin;

import android.app.Activity;
import android.content.Intent;
import android.os.Bundle;
import android.support.annotation.Nullable;
import android.support.v7.app.AppCompatActivity;
import android.util.Log;
import android.view.View;
import android.widget.Button;
import android.widget.Toast;

/**
 * Created by hit-s on 2017/4/15.
 */
public class LogActivity extends AppCompatActivity{
    private Button logView;
    @Override
    protected void onCreate(@Nullable Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.log);
        Intent intent = getIntent();
        String data = intent.getStringExtra("data");
        Toast.makeText(LogActivity.this, data, Toast.LENGTH_SHORT).show();
        logView = (Button) findViewById(R.id.logView);
        logView.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent returnIntent = new Intent();
                returnIntent.putExtra("return", "返回主页面");
                setResult(RESULT_OK, returnIntent);
                finish();
            }
        });
    }

    @Override
    public void onBackPressed() {
        super.onBackPressed();
        Intent returnIntent = new Intent();
        returnIntent.putExtra("return", "通过返回键返回主页面");
        setResult(RESULT_OK, returnIntent);

    }
}
