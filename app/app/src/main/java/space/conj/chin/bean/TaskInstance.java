package space.conj.chin.bean;

import com.google.common.base.Optional;

import java.io.Serializable;
import java.util.Map;

/**
 * Created by hit-s on 2017/4/26.
 */
public class TaskInstance implements Serializable {
    private int id;
    private int taskId;
    private String version;
    private String executeMachine;
    private String pooledTime;
    private String beginTime;
    private String finishTime;
    private int runCount;
    private String status;
    private String log;
    private boolean notify;

    public TaskInstance(Map<String, Object> json) {
        id = (int) json.get("id");
        taskId = (int) json.get("task_id");
        version = (String) json.get("version");
        executeMachine = (String) json.get("execute_machine");
        pooledTime = (String) Optional.fromNullable(json.get("pooled_time")).or("");
        beginTime = (String) Optional.fromNullable(json.get("begin_time")).or("");
        finishTime = (String) Optional.fromNullable(json.get("finish_time")).or("");
        runCount = (int) json.get("run_count");
        status = (String) json.get("status");
        log = (String) json.get("log");
        notify = (boolean) json.get("notify");
    }

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public int getTaskId() {
        return taskId;
    }

    public void setTaskId(int taskId) {
        this.taskId = taskId;
    }

    public String getVersion() {
        return version;
    }

    public void setVersion(String version) {
        this.version = version;
    }

    public String getExecuteMachine() {
        return executeMachine;
    }

    public void setExecuteMachine(String executeMachine) {
        this.executeMachine = executeMachine;
    }

    public String getPooledTime() {
        return pooledTime;
    }

    public void setPooledTime(String pooledTime) {
        this.pooledTime = pooledTime;
    }

    public String getBeginTime() {
        return beginTime;
    }

    public void setBeginTime(String beginTime) {
        this.beginTime = beginTime;
    }

    public String getFinishTime() {
        return finishTime;
    }

    public void setFinishTime(String finishTime) {
        this.finishTime = finishTime;
    }

    public int getRunCount() {
        return runCount;
    }

    public void setRunCount(short runCount) {
        this.runCount = runCount;
    }

    public String getStatus() {
        return status;
    }

    public void setStatus(String status) {
        this.status = status;
    }

    public String getLog() {
        return log;
    }

    public void setLog(String log) {
        this.log = log;
    }

    public boolean isNotify() {
        return notify;
    }

    public void setNotify(boolean notify) {
        this.notify = notify;
    }
}
