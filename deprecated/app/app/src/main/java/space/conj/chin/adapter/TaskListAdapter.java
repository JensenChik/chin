package space.conj.chin.adapter;

import android.content.Context;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ArrayAdapter;
import android.widget.TextView;

import java.util.List;

import space.conj.chin.R;
import space.conj.chin.bean.Task;

/**
 * Created by hit-s on 2017/4/29.
 */
public class TaskListAdapter extends ArrayAdapter<Task> {

    private int resourceId;

    public TaskListAdapter(Context context, int resource, List<Task> taskList) {
        super(context, resource, taskList);
        resourceId = resource;
    }

    @Override
    public View getView(int position, View convertView, ViewGroup parent) {
        Task task = getItem(position);
        View view = LayoutInflater.from(getContext()).inflate(resourceId, null);
        TextView taskName = (TextView) view.findViewById(R.id.task_name);
        taskName.setText(task.getName());
        return view;
    }
}
