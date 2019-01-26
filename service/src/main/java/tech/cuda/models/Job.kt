package tech.cuda.models

import org.jetbrains.exposed.dao.EntityID
import org.jetbrains.exposed.dao.IntEntity
import org.jetbrains.exposed.dao.IntEntityClass
import org.jetbrains.exposed.dao.IntIdTable
import org.joda.time.DateTime
import tech.cuda.enums.ScheduleType
import tech.cuda.exceptions.IllegalScheduleFormatException

/**
 * Created by Jensen on 18-6-18.
 */
object Jobs : IntIdTable() {
    override val tableName: String
        get() = "jobs"

    val task = reference(name = "task_id", foreign = Tasks)
    val removed = bool(name = "removed").index().default(false)
    val createTime = datetime(name = "create_time")
    val updateTime = datetime(name = "update_time")
}


class Job(id: EntityID<Int>) : IntEntity(id) {
    companion object : IntEntityClass<Job>(Jobs)

    var task by Task referencedOn Jobs.task
    var removed by Jobs.removed
    var createTime by Jobs.createTime
    var updateTime by Jobs.updateTime

    val shouldRunNow: Boolean
        get() {
            val regex = """(.+) (.+)-(.+)-(.+) (.+):(.+):(.+)""".toRegex()
            val matchResult = regex.matchEntire(this.task.scheduleFormat)
            if (matchResult != null) {
                val (_, _, _, _, hour, minute, second) = matchResult.destructured
                val now = DateTime.now()
                return now.hourOfDay > hour.toInt()
                        && now.minuteOfHour > minute.toInt()
                        && now.secondOfMinute > second.toInt()
            } else {
                throw IllegalScheduleFormatException("schedule format should be `weekday yyyy-mm-dd HH:MM:SS`")
            }
        }


    fun createInstance() {
        //todo

    }


}