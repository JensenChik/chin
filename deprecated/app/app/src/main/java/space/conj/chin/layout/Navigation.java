package space.conj.chin.layout;

import android.content.Context;
import android.content.Intent;
import android.util.AttributeSet;
import android.view.LayoutInflater;
import android.view.View;
import android.widget.Button;
import android.widget.LinearLayout;

import space.conj.chin.activity.ListLog;
import space.conj.chin.activity.ListMachine;
import space.conj.chin.R;
import space.conj.chin.activity.ListTask;
import space.conj.chin.activity.ListUser;


/**
 * Created by hit-s on 2017/4/17.
 */
public class Navigation extends LinearLayout {
    private Button taskTab;
    private Button logTab;
    private Button userTab;
    private Button machineTab;

    public Navigation(Context context, AttributeSet attrs) {
        super(context, attrs);
        LayoutInflater.from(context).inflate(R.layout.navigation, this);
        taskTab = (Button) findViewById(R.id.taskTab);
        logTab = (Button) findViewById(R.id.logTab);
        userTab = (Button) findViewById(R.id.userTab);
        machineTab = (Button) findViewById(R.id.machineTab);

        taskTab.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent intent = new Intent(getContext(), ListTask.class);
                getContext().startActivity(intent);
            }
        });

        logTab.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent intent = new Intent(getContext(), ListLog.class);
                getContext().startActivity(intent);
            }
        });

        userTab.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent intent = new Intent(getContext(), ListUser.class);
                getContext().startActivity(intent);
            }
        });

        machineTab.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent intent = new Intent(getContext(), ListMachine.class);
                getContext().startActivity(intent);
            }
        });
    }


}
