package space.conj.chin;

import android.content.Intent;
import android.net.Uri;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.widget.Button;
import android.widget.TextView;
import android.widget.Toast;

public class MainActivity extends AppCompatActivity {
    private TextView textView;
    private Button taskTab, logTab, userTab, machineTab;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        textView = (TextView) findViewById(R.id.textView);

        taskTab = (Button) findViewById(R.id.taskTab);
        logTab = (Button) findViewById(R.id.logTab);
        userTab = (Button) findViewById(R.id.userTab);
        machineTab = (Button) findViewById(R.id.machineTab);

        taskTab.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Toast.makeText(MainActivity.this, "查看任务", Toast.LENGTH_SHORT).show();
            }
        });

        logTab.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent intent = new Intent(MainActivity.this, LogActivity.class);
                String data = "查看任务执行情况";
                intent.putExtra("data", data);
                startActivity(intent);
            }
        });

        userTab.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent intent = new Intent(Intent.ACTION_VIEW);
                intent.setData(Uri.parse("http://www.baidu.com"));
                startActivity(intent);
            }
        });
    }

    @Override
    protected void onActivityResult(int requestCode, int resultCode, Intent data) {
        super.onActivityResult(requestCode, resultCode, data);
        switch (requestCode){
            case 1:
                String returnData = data.getStringExtra("return");
                Toast.makeText(MainActivity.this, returnData, Toast.LENGTH_SHORT).show();
                break;
            default:
                break;
        }
    }
}
