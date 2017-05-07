package space.conj.chin.tools;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.google.common.collect.Lists;
import com.squareup.okhttp.Callback;
import com.squareup.okhttp.OkHttpClient;
import com.squareup.okhttp.Request;
import com.squareup.okhttp.Response;

import java.io.IOException;
import java.net.CookieManager;
import java.net.HttpCookie;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import space.conj.chin.R;
import space.conj.chin.adapter.TaskListAdapter;
import space.conj.chin.bean.Task;

/**
 * Created by hit-s on 2017/4/24.
 */
@SuppressWarnings("unchecked")
public class RequestClient {

    private static OkHttpClient client = new OkHttpClient().setCookieHandler(new CookieManager());
    private static final String host = "http://chin.conj.space/";

    private RequestClient() {
    }

    public static OkHttpClient getInstance() {
        return client;
    }

    public static boolean hasCookieOf(String domain) {
        boolean hasCookie = false;
        List<HttpCookie> cookies = ((CookieManager) client.getCookieHandler()).getCookieStore().getCookies();
        for (HttpCookie cookie : cookies) {
            if (cookie.getDomain().equals(domain)) {
                hasCookie = true;
                break;
            }
        }
        return hasCookie;
    }

    public static List<Task> getTaskList() throws IOException {
        Request request = new Request.Builder().url(host + "api/list_task").build();
        String response = client.newCall(request).execute().body().string();
        Map<String, Object> responseJson = new ObjectMapper().readValue(response, HashMap.class);

        List<Task> taskList = Lists.newArrayList();
        for (Map<String, Object> metaJson : (List<Map>) responseJson.get("data")) {
            taskList.add(new Task(metaJson));
        }
        return taskList;
    }


}
