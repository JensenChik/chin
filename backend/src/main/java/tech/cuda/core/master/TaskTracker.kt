package tech.cuda.core.master

import org.joda.time.DateTime
import org.joda.time.Days
import tech.cuda.services.JobService
import tech.cuda.services.TaskService

/**
 * Created by Jensen on 18-10-25.
 * 只能在 job 里 insert
 */
object TaskTracker {
    var lastDateTime: DateTime = DateTime()
    private val newDayComing = Days.daysBetween(lastDateTime, DateTime()).days > 0

    fun serve() {
        while (true) {
            if (newDayComing) {
                lastDateTime = DateTime()
                TaskService.getMany().filter {
                    it.shouldScheduledToday
                }.forEach {
                    JobService.createOneForTask(it)
                }
            } else {
                Thread.sleep(1000)
            }
        }
    }

}