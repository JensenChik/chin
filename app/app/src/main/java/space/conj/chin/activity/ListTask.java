package space.conj.chin.activity;

import android.content.Intent;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.widget.AdapterView;
import android.widget.ListView;

import java.io.IOException;
import java.util.List;

import space.conj.chin.R;
import space.conj.chin.adapter.TaskListAdapter;
import space.conj.chin.bean.Task;
import space.conj.chin.tools.RequestClient;

@SuppressWarnings("unchecked")
public class ListTask extends AppCompatActivity {

    private TaskListAdapter adapter;
    private ListView taskListView;
    private List<Task> taskList;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.task);
        taskListView = (ListView) findViewById(R.id.list_task);

        initTaskListView();

        taskListView.setOnItemClickListener(new AdapterView.OnItemClickListener() {
            @Override
            public void onItemClick(AdapterView<?> parent, View view, int position, long id) {
                Task task = taskList.get(position);
                Intent intent = new Intent(ListTask.this, TaskDetail.class);
                Bundle bundle = new Bundle();
                bundle.putSerializable("task", task);
                intent.putExtras(bundle);
                startActivity(intent);
            }
        });

    }

    private void initTaskListView() {
        new Thread(new Runnable() {
            @Override
            public void run() {
                try {
                    taskList = RequestClient.getTaskList();
                    adapter = new TaskListAdapter(ListTask.this, R.layout.task_item, taskList);
                    runOnUiThread(new Runnable() {
                        @Override
                        public void run() {
                            taskListView.setAdapter(adapter);
                        }
                    });
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }
        }).start();
    }

}
