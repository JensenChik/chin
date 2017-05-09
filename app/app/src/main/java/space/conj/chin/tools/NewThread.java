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

    public static void run(final Object object, String methodName, final Object[] args) {
        try {
            Class<?>[] argsClassType = new Class[args.length];
            for (int i = 0; i < args.length; i++) {
                argsClassType[i] = args[i].getClass();
            }

            final Method method = object.getClass().getDeclaredMethod(methodName, argsClassType);
            method.setAccessible(true);
            new Thread(new Runnable() {
                @Override
                public void run() {
                    try {
                        method.invoke(object, args);
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
