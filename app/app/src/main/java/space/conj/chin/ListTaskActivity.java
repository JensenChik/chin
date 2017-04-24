package space.conj.chin;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.util.Log;
import android.widget.ArrayAdapter;
import android.widget.ListView;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.squareup.okhttp.Callback;
import com.squareup.okhttp.OkHttpClient;
import com.squareup.okhttp.Request;
import com.squareup.okhttp.Response;

import java.io.IOException;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import space.conj.chin.tools.RequestClient;

@SuppressWarnings("unchecked")
public class ListTaskActivity extends AppCompatActivity {

    private ArrayAdapter<String> adapter;
    private ListView tasksList;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.task);
        tasksList = (ListView) findViewById(R.id.list_task);

        OkHttpClient client = RequestClient.getInstance();
        Request request = new Request.Builder()
                .url("http://chin.nazgrim.com/api/list_task")
                .build();
        client.newCall(request).enqueue(new Callback() {
            @Override
            public void onFailure(Request request, IOException e) {

            }

            @Override
            public void onResponse(final Response response) throws IOException {
                Log.i("COOKIE", response.header("Set-Cookie"));
                final String json = response.body().string();
                Log.i("json", json);
                List<Map<String, Object>> tasks = (List<Map<String, Object>>) new ObjectMapper()
                        .readValue(json, HashMap.class).get("data");
                final String[] taskName = new String[tasks.size()];
                for (int i = 0; i < taskName.length; i++) {
                    taskName[i] = tasks.get(i).get("name").toString();
                }

                runOnUiThread(new Runnable() {
                    @Override
                    public void run() {
                        adapter = new ArrayAdapter<>(ListTaskActivity.this,
                                R.layout.support_simple_spinner_dropdown_item, taskName);
                        tasksList.setAdapter(adapter);
                    }
                });
            }
        });

    }


}
