package space.conj.chin.activity;

import android.content.Intent;
import android.os.Bundle;
import android.support.annotation.Nullable;
import android.support.annotation.StringDef;
import android.support.v7.app.AppCompatActivity;
import android.widget.TextView;

import java.io.Serializable;

import space.conj.chin.R;
import space.conj.chin.bean.TaskInstance;

/**
 * Created by hit-s on 2017/5/15.
 */
public class InstanceDetail extends AppCompatActivity{

    private TextView id;
    private TextView taskId;
    private TextView version;
    private TextView executeMachine;
    private TextView pooledTime;
    private TextView beginTime;
    private TextView finishTime;
    private TextView runCount;
    private TextView status;
    private TextView log;

    @Override
    protected void onCreate(@Nullable Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.instance_detail);
        id = (TextView) findViewById(R.id.id);
        taskId = (TextView) findViewById(R.id.task_id);
        version = (TextView) findViewById(R.id.version);
        executeMachine = (TextView) findViewById(R.id.execute_machine);
        pooledTime = (TextView) findViewById(R.id.pooled_time);
        beginTime = (TextView) findViewById(R.id.begin_time);
        finishTime = (TextView) findViewById(R.id.finish_time);
        runCount = (TextView) findViewById(R.id.run_count);
        status = (TextView) findViewById(R.id.status);
        log = (TextView) findViewById(R.id.log);

        Intent intent = getIntent();
        TaskInstance instance = (TaskInstance) intent.getSerializableExtra("instance");

        id.setText(String.valueOf(instance.getId()));
        taskId.setText(String.valueOf(instance.getTaskId()));
        version.setText(instance.getVersion());
        executeMachine.setText(instance.getExecuteMachine());
        pooledTime.setText(instance.getPooledTime());
        beginTime.setText(instance.getBeginTime());
        finishTime.setText(instance.getFinishTime());
        runCount.setText(String.valueOf(instance.getRunCount()));
        status.setText(instance.getStatus());
        log.setText("调用网络接口获取日志");

    }
}
