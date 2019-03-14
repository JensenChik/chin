package tech.cuda.core.master

import org.joda.time.DateTime
import org.joda.time.Days
import tech.cuda.services.InstanceService
import tech.cuda.services.JobService

/**
 * Created by Jensen on 18-10-26.
 * 只能在 instance 里新增
 */
object JobTracker {

    fun serve() {
        while (true) {
            val today = DateTime.now().toDate()
            JobService.getManyByDate(today).filter {
                it.shouldRunNow
            }.forEach {
                InstanceService.createOneForJob(it)
            }
            Thread.sleep(1000)
        }
    }


}