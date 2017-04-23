package space.conj.chin;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.util.Log;
import android.widget.ArrayAdapter;
import android.widget.ListView;
import android.widget.TextView;

import com.squareup.okhttp.Callback;
import com.squareup.okhttp.OkHttpClient;
import com.squareup.okhttp.Request;
import com.squareup.okhttp.Response;

import java.io.IOException;

import space.conj.chin.tools.RequestClient;

public class TaskActivity extends AppCompatActivity {

    private String[] task = {"【爬虫】京东显卡数据", "ip池维护", "【爬虫】京东硬盘数据", "【爬虫】京东手机数据",
            "【爬虫】京东主板数据", "【爬虫】京东显示器数据", "【爬虫】京东相机数据", "【爬虫】京东电视数据", "【爬虫】京东空调数据",
            "【爬虫】京东洗衣机数据", "【爬虫】京东冰箱数据", "【爬虫】京东笔记本数据", "【爬虫】京东平板数据",
            "【爬虫】京东路由器数据", "【数据仓库source】京东数据sqoop同步"};
    private TextView httpResponse;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.task);
        ArrayAdapter<String> adapter = new ArrayAdapter<>(TaskActivity.this,
                R.layout.support_simple_spinner_dropdown_item, task);
        ListView tasks = (ListView) findViewById(R.id.list_task);
        httpResponse = (TextView) findViewById(R.id.http_response);

        OkHttpClient client = RequestClient.getInstance();

        Request request = new Request.Builder()
                .url("http://chin.nazgrim.com/get_log_by_page?order=asc&offset=10&limit=10")
                .build();
        client.newCall(request).enqueue(new Callback() {
            @Override
            public void onFailure(Request request, IOException e) {

            }

            @Override
            public void onResponse(final Response response) throws IOException {
                Log.i("COOKIE", response.header("Set-Cookie"));
                final String content = response.body().string();
                Log.i("HTML", content);
                runOnUiThread(new Runnable() {
                    @Override
                    public void run() {
                        httpResponse.setText(content);
                    }
                });
            }
        });

        assert tasks != null;
        tasks.setAdapter(adapter);
    }


}
