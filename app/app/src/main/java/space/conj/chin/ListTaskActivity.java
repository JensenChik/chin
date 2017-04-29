package space.conj.chin;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.util.Log;
import android.widget.ListView;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.google.common.collect.Lists;
import com.squareup.okhttp.Callback;
import com.squareup.okhttp.OkHttpClient;
import com.squareup.okhttp.Request;
import com.squareup.okhttp.Response;

import java.io.IOException;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import space.conj.chin.adapter.TaskListAdapter;
import space.conj.chin.bean.Task;
import space.conj.chin.tools.RequestClient;

@SuppressWarnings("unchecked")
public class ListTaskActivity extends AppCompatActivity {

    private TaskListAdapter adapter;
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
                final String responseJson = response.body().string();
                Map<String, Object> respondMap = new ObjectMapper().readValue(responseJson, HashMap.class);

                final List<Task> taskList = Lists.newArrayList();
                for (Map<String, Object> json : (List<Map>) respondMap.get("data")) {
                    taskList.add(new Task(json));
                }

                final String[] taskName = new String[taskList.size()];
                for (int i = 0; i < taskName.length; i++) {
                    taskName[i] = taskList.get(i).getName();
                }

                runOnUiThread(new Runnable() {
                    @Override
                    public void run() {
                        adapter = new TaskListAdapter(ListTaskActivity.this, R.layout.task_item, taskList);
                        tasksList.setAdapter(adapter);
                    }
                });
            }
        });

    }


}
