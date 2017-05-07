package space.conj.chin.tools;

import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;

/**
 * Created by hit-s on 2017/5/7.
 */
public class NewThread {
    public static void run(final Object object, String methodName) {
        try {
            final Method method = object.getClass().getDeclaredMethod(methodName);
            method.setAccessible(true);
            new Thread(new Runnable() {
                @Override
                public void run() {
                    try {
                        method.invoke(object);
                    } catch (IllegalAccessException | InvocationTargetException e) {
                        e.printStackTrace();
                    }
                }
            }).start();
        } catch (NoSuchMethodException e) {
            e.printStackTrace();
        }


    }
}
