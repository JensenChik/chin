package space.conj.chin.bean;

import com.google.common.base.Optional;

import org.joda.time.DateTime;

import java.io.Serializable;
import java.util.List;
import java.util.Map;

/**
 * Created by hit-s on 2017/4/26.
 */
@SuppressWarnings("unchecked")
public class Task implements Serializable {
    private int id;
    private String name;
    private String createTime;
    private String command;
    private int priority;
    private List<String> machinePool;
    private List<Integer> fatherTask;
    private List<Integer> childTask;
    private boolean valid;
    private boolean rerun;
    private int rerunTimes;
    private String scheduledType;
    private String scheduledTime;
    private int year;
    private int month;
    private int weekday;
    private int day;
    private int hour;
    private int minute;


    public Task(Map<String, Object> json) {
        id = (int) json.get("id");
        name = (String) json.get("name");
        createTime = (String) json.get("create_time");
        command = (String) json.get("command");
        priority = (int) json.get("priority");
        machinePool = (List<String>) json.get("machine_pool");
        fatherTask = (List<Integer>) json.get("father_task");
        childTask = (List<Integer>) json.get("child_task");
        valid = (boolean) json.get("valid");
        rerun = (boolean) json.get("rerun");
        rerunTimes = (int) json.get("rerun_times");

        scheduledType = (String) json.get("scheduled_type");
        year = (int) Optional.fromNullable(json.get("year")).or(-1);
        month = (int) Optional.fromNullable(json.get("month")).or(-1);
        weekday = (int) Optional.fromNullable(json.get("weekday")).or(-1);
        day = (int) Optional.fromNullable(json.get("day")).or(-1);
        hour = (int) Optional.fromNullable(json.get("hour")).or(-1);
        minute = (int) Optional.fromNullable(json.get("minute")).or(-1);

        switch (scheduledType) {
            case "day":
                scheduledType = "每天";
                scheduledTime = new DateTime(1900, 1, 1, hour, minute, 0).toString("HH:mm:ss");
                break;
            case "week":
                scheduledTime = "每周";
                scheduledTime = "周" + weekday + "\t"
                        + new DateTime(1900, 1, 1, hour, minute, 0).toString("HH:mm:ss");
                break;
            case "month":
                scheduledTime = "每月";
                scheduledTime = day + "号" + "\t"
                        + new DateTime(1900, 1, 1, hour, minute, 0).toString("HH:mm:ss");
                break;
            case "once":
                scheduledTime = "一次";
                scheduledTime = new DateTime(year, month, day, hour, minute, 0)
                        .toString("yyyy-MM-dd HH:mm:ss");
                break;
        }

    }

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getCreateTime() {
        return createTime;
    }

    public void setCreateTime(String createTime) {
        this.createTime = createTime;
    }

    public String getCommand() {
        return command;
    }

    public void setCommand(String command) {
        this.command = command;
    }

    public int getPriority() {
        return priority;
    }

    public void setPriority(short priority) {
        this.priority = priority;
    }

    public List<String> getMachinePool() {
        return machinePool;
    }

    public void setMachinePool(List<String> machinePool) {
        this.machinePool = machinePool;
    }

    public List<Integer> getFatherTask() {
        return fatherTask;
    }

    public void setFatherTask(List<Integer> fatherTask) {
        this.fatherTask = fatherTask;
    }

    public List<Integer> getChildTask() {
        return childTask;
    }

    public void setChildTask(List<Integer> childTask) {
        this.childTask = childTask;
    }

    public boolean isValid() {
        return valid;
    }

    public void setValid(boolean valid) {
        this.valid = valid;
    }

    public boolean isRerun() {
        return rerun;
    }

    public void setRerun(boolean rerun) {
        this.rerun = rerun;
    }

    public int getRerunTimes() {
        return rerunTimes;
    }

    public void setRerunTimes(short rerunTimes) {
        this.rerunTimes = rerunTimes;
    }

    public String getScheduledType() {
        return scheduledType;
    }

    public void setScheduledType(String scheduledType) {
        this.scheduledType = scheduledType;
    }

    public int getYear() {
        return year;
    }

    public void setYear(short year) {
        this.year = year;
    }

    public int getMonth() {
        return month;
    }

    public void setMonth(short month) {
        this.month = month;
    }

    public int getWeekday() {
        return weekday;
    }

    public void setWeekday(short weekday) {
        this.weekday = weekday;
    }

    public int getDay() {
        return day;
    }

    public void setDay(short day) {
        this.day = day;
    }

    public int getHour() {
        return hour;
    }

    public void setHour(short hour) {
        this.hour = hour;
    }

    public int getMinute() {
        return minute;
    }

    public void setMinute(short minute) {
        this.minute = minute;
    }

    public String getScheduledTime() {
        return scheduledTime;
    }

    public void setScheduledTime(String scheduledTime) {
        this.scheduledTime = scheduledTime;
    }
}
