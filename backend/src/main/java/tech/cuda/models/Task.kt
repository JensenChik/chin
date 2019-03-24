package tech.cuda.models

import org.jetbrains.exposed.dao.EntityID
import org.jetbrains.exposed.dao.IntEntity
import org.jetbrains.exposed.dao.IntEntityClass
import org.jetbrains.exposed.dao.IntIdTable
import org.joda.time.DateTime
import org.joda.time.Days
import tech.cuda.exceptions.IllegalScheduleFormatException
import tech.cuda.exceptions.InvalidParameterException


/**
 * Created by Jensen on 18-6-18.
 */

enum class ScheduleType {
    Day, Week, Month, Year, Once;
}

class ScheduleFormat {
    val weekday: Int?
    val year: Int?
    val month: Int?
    val day: Int?
    val hour: Int
    val minute: Int
    val second: Int

    constructor(
            weekday: Int? = null,
            year: Int? = null, month: Int? = null, day: Int? = null,
            hour: Int, minute: Int, second: Int = 0
    ) {
        when {
            weekday != null && weekday !in 0..6 -> throw InvalidParameterException("weekday must in 0..6")
            year != null && year !in 2019..2099 -> throw InvalidParameterException("year must in 2019..2099")
            month != null && month !in 1..12 -> throw InvalidParameterException("month must in 1..12")
            day != null && day !in 1..31 -> throw InvalidParameterException("day must in 1..31")
            hour !in 0..23 -> throw InvalidParameterException("hour must in 0..23")
            minute !in 0..59 -> throw InvalidParameterException("minute must in 0..59")
            second !in 0..59 -> throw InvalidParameterException("second must in 0..59")
            else -> {
                this.weekday = weekday
                this.year = year
                this.month = month
                this.day = day
                this.hour = hour
                this.minute = minute
                this.second = second
            }
        }
    }

    constructor(format: String) {
        val regex = """(.+) (.+)-(.+)-(.+) (.+):(.+):(.+)""".toRegex()
        val matchResult = regex.matchEntire(format)
        val (weekday, year, month, day, hour, minute, second) = matchResult!!.destructured
        this.weekday = weekday.toIntOrNull()
        this.year = year.toIntOrNull()
        this.month = month.toIntOrNull()
        this.day = day.toIntOrNull()
        this.hour = hour.toInt()
        this.minute = minute.toInt()
        this.second = second.toInt()
    }

    override fun toString(): String {
        val weekday = this.weekday?.toString() ?: "*"
        val year = this.year?.toString() ?: "****"
        val month = this.month?.toString() ?: "**"
        val day = this.day?.toString() ?: "**"
        val hour = this.hour
        val minute = this.minute
        val second = this.second
        return "$weekday $year-$month-$day $hour:$minute:$second"
    }

}

object TaskTable : IntIdTable() {
    override val tableName: String
        get() = "tasks"

    val user = reference(name = "user_id", foreign = UserTable)
    const val NAME_MAX_LEN = 256
    val name = varchar(name = "name", length = NAME_MAX_LEN)
    val scheduleType = customEnumeration(
            name = "schedule_type",
            sql = ScheduleType.values().joinToString(",", "ENUM(", ")") { "'${it.name}'" },
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
    var scheduleFormat by TaskTable.scheduleFormat.transform(
            toColumn = { format: ScheduleFormat -> format.toString() },
            toReal = { format: String -> ScheduleFormat(format) }
    )
    var name by TaskTable.name
    var command by TaskTable.command
    var latestJobId by TaskTable.latestJobId
    var removed by TaskTable.removed
    var createTime by TaskTable.createTime
    var updateTime by TaskTable.updateTime

    val shouldScheduledToday: Boolean
        get() {
            val now = DateTime.now()
            return when (this.scheduleType) {
                ScheduleType.Week -> now.dayOfWeek == this.scheduleFormat.weekday
                ScheduleType.Once -> now.year == this.scheduleFormat.year
                        && now.monthOfYear == this.scheduleFormat.month
                        && now.dayOfMonth == this.scheduleFormat.day
                ScheduleType.Year -> now.monthOfYear == this.scheduleFormat.month
                        && now.dayOfMonth == this.scheduleFormat.day
                ScheduleType.Month -> now.dayOfMonth == this.scheduleFormat.day
                ScheduleType.Day -> true
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

    val todayStatus: InstanceStatus
        get() {
            // todo
            return InstanceStatus.Success
        }

    val lastThreeDayStatus: List<InstanceStatus>
        get() {
            // todo
            return listOf(InstanceStatus.Success, InstanceStatus.Success, InstanceStatus.Success)
        }
}

