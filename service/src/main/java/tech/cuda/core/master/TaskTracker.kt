package tech.cuda.core.master

import org.jetbrains.exposed.sql.select
import org.joda.time.DateTime
import org.joda.time.Days
import tech.cuda.models.Task
import tech.cuda.models.Tasks

/**
 * Created by Jensen on 18-10-25.
 * 只能在 job 里 insert
 */
object TaskTracker {
    var lastDateTime: DateTime = DateTime()

    fun newJobForRoutineTask() {
        Task.find { Tasks.removed eq false }.forEach {
            if (it.shouldScheduledToday) {
                it.createJob()
            }
        }
    }

    val newDayComing: Boolean
        get() {
            return Days.daysBetween(lastDateTime, DateTime()).days > 0
        }

    fun serve() {
        while (true) {
            if (newDayComing) {
                lastDateTime = DateTime()
                newJobForRoutineTask()
            } else {
                Thread.sleep(1000)
            }
        }
    }

}