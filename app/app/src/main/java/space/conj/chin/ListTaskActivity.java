package space.conj.chin;

import android.content.Intent;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.util.Log;
import android.view.View;
import android.widget.AdapterView;
import android.widget.ListView;
import android.widget.Toast;

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
    private ListView taskListView;
    private List<Task> taskList = Lists.newArrayList();

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.task);
        taskListView = (ListView) findViewById(R.id.list_task);

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
                Map<String, Object> responseJson = new ObjectMapper()
                        .readValue(response.body().string(), HashMap.class);

                for (Map<String, Object> metaJson : (List<Map>) responseJson.get("data")) {
                    taskList.add(new Task(metaJson));
                }

                runOnUiThread(new Runnable() {
                    @Override
                    public void run() {
                        adapter = new TaskListAdapter(ListTaskActivity.this, R.layout.task_item, taskList);
                        taskListView.setAdapter(adapter);
                    }
                });
            }
        });

        taskListView.setOnItemClickListener(new AdapterView.OnItemClickListener() {
            @Override
            public void onItemClick(AdapterView<?> parent, View view, int position, long id) {
                Task task = taskList.get(position);
                Intent intent = new Intent(ListTaskActivity.this, TaskDetailActivity.class);
                Bundle bundle = new Bundle();
                bundle.putSerializable("task", task);
                intent.putExtras(bundle);
                startActivity(intent);
            }
        });

    }


}
