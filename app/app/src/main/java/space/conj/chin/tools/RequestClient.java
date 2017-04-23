package space.conj.chin.tools;

import com.squareup.okhttp.OkHttpClient;

import java.net.CookieManager;

/**
 * Created by hit-s on 2017/4/24.
 */
public class RequestClient {

    private static OkHttpClient client = new OkHttpClient().setCookieHandler(new CookieManager());

    private RequestClient() {
    }

    public static OkHttpClient getInstance() {
        return client;
    }

    ;
}
