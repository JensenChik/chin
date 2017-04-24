package space.conj.chin.tools;

import com.squareup.okhttp.OkHttpClient;

import java.net.CookieManager;
import java.net.HttpCookie;
import java.util.List;

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

    public static boolean hasCookieOf(String domain){
        boolean hasCookie = false;
        List<HttpCookie> cookies = ((CookieManager) client.getCookieHandler()).getCookieStore().getCookies();
        for (HttpCookie cookie : cookies) {
            if(cookie.getDomain().equals(domain)) {
                hasCookie = true;
                break;
            }
        }
        return hasCookie;
    }

    ;
}
