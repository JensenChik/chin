package tech.cuda.core.master

import org.joda.time.DateTime
import org.joda.time.Days

/**
 * Created by Jensen on 18-10-26.
 * 只能在 instance 里新增
 */
object JobTracker {


    fun newInstanceForReadyJob() {
        //todo
    }

    fun serve() {
        while (true) {
            newInstanceForReadyJob()

            Thread.sleep(1000)
        }
    }


}