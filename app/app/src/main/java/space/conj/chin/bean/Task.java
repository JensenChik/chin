package space.conj.chin.bean;

import java.util.List;

/**
 * Created by hit-s on 2017/4/26.
 */
public class Task {
    private int id;
    private String name;
    private String createTime;
    private String command;
    private short priority;
    private List<String> machinePool;
    private List<Integer> fatherTask;
    private List<Integer> childTask;
    private boolean valid;
    private boolean rerun;
    private short rerunTimes;
    private String scheduledType;
    private short year;
    private short month;
    private short weekday;
    private short day;
    private short hour;
    private short minute;

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

    public short getPriority() {
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

    public short getRerunTimes() {
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

    public short getYear() {
        return year;
    }

    public void setYear(short year) {
        this.year = year;
    }

    public short getMonth() {
        return month;
    }

    public void setMonth(short month) {
        this.month = month;
    }

    public short getWeekday() {
        return weekday;
    }

    public void setWeekday(short weekday) {
        this.weekday = weekday;
    }

    public short getDay() {
        return day;
    }

    public void setDay(short day) {
        this.day = day;
    }

    public short getHour() {
        return hour;
    }

    public void setHour(short hour) {
        this.hour = hour;
    }

    public short getMinute() {
        return minute;
    }

    public void setMinute(short minute) {
        this.minute = minute;
    }
}
