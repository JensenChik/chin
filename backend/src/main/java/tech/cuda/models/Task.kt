package tech.cuda.models

import org.jetbrains.exposed.dao.EntityID
import org.jetbrains.exposed.dao.IntEntity
import org.jetbrains.exposed.dao.IntEntityClass
import org.jetbrains.exposed.dao.IntIdTable
import org.joda.time.DateTime
import org.joda.time.Days
import tech.cuda.enums.SQL
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
    val hour: Int?
    val minute: Int?
    val second: Int?

    constructor(weekday: Int?,
                year: Int?, month: Int?, day: Int?,
                hour: Int?, minute: Int?, second: Int?) {
        when {
            weekday != null && weekday !in 0..6 -> throw InvalidParameterException("weekday must in 0..6")
            year != null && year !in 2019..2099 -> throw InvalidParameterException("year must in 2019..2099")
            month != null && month !in 1..12 -> throw InvalidParameterException("month must in 1..12")
            day != null && day !in 1..31 -> throw InvalidParameterException("day must in 1..31")
            hour != null && hour !in 0..23 -> throw InvalidParameterException("hour must in 0..23")
            minute != null && minute !in 0..59 -> throw InvalidParameterException("minute must in 0..59")
            second != null && second !in 0..59 -> throw InvalidParameterException("sedond must in 0..59")
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
        this.hour = hour.toIntOrNull()
        this.minute = minute.toIntOrNull()
        this.second = second.toIntOrNull()
    }

    override fun toString(): String {
        val weekday = this.weekday?.toString() ?: "*"
        val year = this.year?.toString() ?: "****"
        val month = this.month?.toString() ?: "**"
        val day = this.day?.toString() ?: "**"
        val hour = this.hour?.toString() ?: "**"
        val minute = this.minute?.toString() ?: "**"
        val second = this.second?.toString() ?: "**"
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

