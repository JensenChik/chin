package tech.cuda.models.mappers

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
object Tasks : IntIdTable() {
    override val tableName: String
        get() = "tasks"

    val user = reference(name = "user_id", foreign = Users)
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
    companion object : IntEntityClass<Task>(Tasks)

    var user by User referencedOn Tasks.user

    var scheduleType by Tasks.scheduleType
    var scheduleFormat by Tasks.scheduleFormat
    var name by Tasks.name
    var command by Tasks.command
    var latestJobId by Tasks.latestJobId
    var removed by Tasks.removed
    var createTime by Tasks.createTime
    var updateTime by Tasks.updateTime

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


    fun createJob() {
        val now = DateTime.now()
        val job = Job.new {
            task = this@Task
            createTime = now
            updateTime = now
        }
        this.latestJobId = job.id.value
    }

}

