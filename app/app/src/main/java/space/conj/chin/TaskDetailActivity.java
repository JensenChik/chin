package space.conj.chin;

import android.content.Intent;
import android.os.Bundle;
import android.support.annotation.Nullable;
import android.support.v7.app.AppCompatActivity;
import android.util.Log;
import android.widget.TextView;


import space.conj.chin.bean.Task;

/**
 * Created by hit-s on 2017/4/30.
 */
public class TaskDetailActivity extends AppCompatActivity {

    private TextView taskId;
    private TextView taskName;
    private TextView createTime;
    private TextView command;
    private TextView priority;
    private TextView machinePool;
    private TextView fatherTask;
    private TextView childTask;
    private TextView valid;
    private TextView rerun;
    private TextView rerunTimes;
    private TextView scheduledType;
    private TextView scheduledTime;

    @Override
    protected void onCreate(@Nullable Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.task_detail);
        taskId = (TextView) findViewById(R.id.task_id);
        taskName = (TextView) findViewById(R.id.task_name);
        createTime = (TextView) findViewById(R.id.create_time);
        command = (TextView) findViewById(R.id.command);
        priority = (TextView) findViewById(R.id.priority);
        machinePool = (TextView) findViewById(R.id.machine_pool);
        fatherTask = (TextView) findViewById(R.id.father_task);
        childTask = (TextView) findViewById(R.id.child_task);
        valid = (TextView) findViewById(R.id.valid);
        rerun = (TextView) findViewById(R.id.rerun);
        rerunTimes = (TextView) findViewById(R.id.rerun_times);
        scheduledType = (TextView) findViewById(R.id.scheduled_type);
        scheduledTime = (TextView) findViewById(R.id.scheduled_time);

        Intent intent = getIntent();
        Task task = (Task) intent.getSerializableExtra("task");

        taskId.setText(String.valueOf(task.getId()));
        taskName.setText(task.getName());
        createTime.setText(task.getCreateTime());
        command.setText(task.getCommand());
        priority.setText(String.valueOf(task.getPriority()));

    }
}
