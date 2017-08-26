package space.conj.chin.adapter;

import android.content.Context;
import android.graphics.Color;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ArrayAdapter;
import android.widget.TextView;

import java.util.List;

import space.conj.chin.R;
import space.conj.chin.bean.TaskInstance;

/**
 * Created by hit-s on 2017/5/10.
 */
public class InstanceListAdapter extends ArrayAdapter<TaskInstance> {

    private int resourceId;

    public InstanceListAdapter(Context context, int resource, List<TaskInstance> taskInstance) {
        super(context, resource, taskInstance);
        resourceId = resource;
    }

    @Override
    public View getView(int position, View convertView, ViewGroup parent) {
        TaskInstance item = getItem(position);
        View view = LayoutInflater.from(getContext()).inflate(resourceId, null);

        TextView version = (TextView) view.findViewById(R.id.version);
        TextView finishTime = (TextView) view.findViewById(R.id.finish_time);

        version.setText(item.getVersion());
        finishTime.setText(item.getFinishTime());

        switch (item.getStatus()) {
            case "success":
                version.setTextColor(Color.rgb(92, 184, 92));
                finishTime.setTextColor(Color.rgb(92, 184, 92));
                break;
            case "running":
                version.setTextColor(Color.rgb(240, 173, 78));
                finishTime.setTextColor(Color.rgb(240, 173, 78));
                break;
            case "failed":
                version.setTextColor(Color.rgb(217, 83, 79));
                finishTime.setTextColor(Color.rgb(217, 83, 79));
                break;
        }

        return view;
    }
}
