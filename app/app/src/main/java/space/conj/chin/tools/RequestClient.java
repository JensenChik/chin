package space.conj.chin.tools;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.google.common.collect.Lists;
import com.squareup.okhttp.FormEncodingBuilder;
import com.squareup.okhttp.OkHttpClient;
import com.squareup.okhttp.Request;

import java.io.IOException;
import java.net.CookieManager;
import java.net.HttpCookie;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import space.conj.chin.bean.Task;
import space.conj.chin.bean.TaskInstance;

/**
 * Created by hit-s on 2017/4/24.
 */
@SuppressWarnings("unchecked")
public class RequestClient {

    private static OkHttpClient client = new OkHttpClient().setCookieHandler(new CookieManager());
    private static final String domain = "chin.conj.space";
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

    public static boolean login(String userName, String password) {
        FormEncodingBuilder builder = new FormEncodingBuilder();
        builder.add("user_name", userName);
        builder.add("password", password);
        Request request = new Request.Builder().url(host + "login").post(builder.build()).build();
        try {
            client.newCall(request).execute();
        } catch (IOException e) {
            return false;
        }
        return hasCookieOf(domain);
    }

    private static Map<String, Object> getJson(String url) {
        try {
            Request request = new Request.Builder().url(url).build();
            String response = client.newCall(request).execute().body().string();
            return new ObjectMapper().readValue(response, HashMap.class);
        } catch (IOException e) {
            return null;
        }
    }

    public static List<Task> getTaskList() {
        Map<String, Object> json = getJson(host + "api/list_task");
        assert json != null;

        List<Task> taskList = Lists.newArrayList();
        for (Map<String, Object> metaJson : (List<Map>) json.get("data")) {
            taskList.add(new Task(metaJson));
        }
        return taskList;
    }

    public static List<TaskInstance> getTaskInstanceOf(int taskId) {
        Map<String, Object> json = getJson(host + "api/list_instance_of/" + taskId);
        assert json != null;

        List<TaskInstance> taskInstance = Lists.newArrayList();
        for (Map<String, Object> instance : (List<Map>) json.get("data")) {
            taskInstance.add(new TaskInstance(instance));
        }
        return taskInstance;
    }

}
