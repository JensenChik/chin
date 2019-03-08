package tech.cuda.models

import org.jetbrains.exposed.dao.EntityID
import org.jetbrains.exposed.dao.IntEntity
import org.jetbrains.exposed.dao.IntEntityClass
import org.jetbrains.exposed.dao.IntIdTable
import org.joda.time.DateTime
import org.joda.time.Days
import tech.cuda.enums.ScheduleType
import tech.cuda.enums.SQL
import tech.cuda.exceptions.IllegalScheduleFormatException


/**
 * Created by Jensen on 18-6-18.
 */
object TaskTable : IntIdTable() {
    override val tableName: String
        get() = "tasks"

    val user = reference(name = "user_id", foreign = UserTable)
    val name = varchar(name = "name", length = 256)
    val scheduleType = customEnumeration(
            name = "schedule_type", sql = SQL<ScheduleType>(),
            fromDb = { value -> ScheduleType.valueOf(value as String) },
            toDb = { it.name }
    )
    val scheduleFormat = varchar(name = "schedule_format", length = 256)
    val command = text(name = "command")
    val latestJobId = integer(name = "latest_job_id").nullable()
    val removed = bool(name = "removed").index().default(false)
    val createTime = datetime(name = "create_time")
    val updateTime = datetime(name = "update_time")
}


class Task(id: EntityID<Int>) : IntEntity(id) {
    companion object : IntEntityClass<Task>(TaskTable)

    var user by User referencedOn TaskTable.user

    var scheduleType by TaskTable.scheduleType
    var scheduleFormat by TaskTable.scheduleFormat
    var name by TaskTable.name
    var command by TaskTable.command
    var latestJobId by TaskTable.latestJobId
    var removed by TaskTable.removed
    var createTime by TaskTable.createTime
    var updateTime by TaskTable.updateTime

    val shouldScheduledToday: Boolean
        get() {
            val regex = """(.+) (.+)-(.+)-(.+) (.+):(.+):(.+)""".toRegex()
            val matchResult = regex.matchEntire(this.scheduleFormat)
            if (matchResult != null) {
                val (weekDay, year, month, day, _, _, _) = matchResult.destructured
                val now = DateTime.now()
                return when (this.scheduleType) {
                    ScheduleType.Week -> now.dayOfWeek == weekDay.toInt()
                    ScheduleType.Once -> now.year == year.toInt()
                            && now.monthOfYear == month.toInt()
                            && now.dayOfMonth == day.toInt()
                    ScheduleType.Year -> now.monthOfYear == month.toInt()
                            && now.dayOfMonth == day.toInt()
                    ScheduleType.Month -> now.dayOfMonth == day.toInt()
                    ScheduleType.Day -> true
                }
            } else {
                throw IllegalScheduleFormatException("schedule format should be `weekday yyyy-mm-dd HH:MM:SS`")
            }
        }


    val jobCreated: Boolean
        get() {
            return if (this.latestJobId != null) {
                val latestJob = Job.findById(this.latestJobId!!)
                latestJob != null && Days.daysBetween(latestJob.createTime, DateTime()).days == 0
            } else {
                false
            }
        }
}

