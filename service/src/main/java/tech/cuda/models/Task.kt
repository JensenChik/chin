package tech.cuda.models

import org.jetbrains.exposed.dao.EntityID
import org.jetbrains.exposed.dao.IntEntity
import org.jetbrains.exposed.dao.IntEntityClass
import org.jetbrains.exposed.dao.IntIdTable
import tech.cuda.enums.ScheduleType
import tech.cuda.enums.SQL


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
    val command = text(name = "command")
    val removed = bool(name = "removed").index().default(false)
    val createTime = datetime(name = "create_time")
    val updateTime = datetime(name = "update_time")
}


class Task(id: EntityID<Int>) : IntEntity(id) {
    companion object : IntEntityClass<Task>(Tasks)

    var user by User referencedOn Tasks.user

    var scheduleType by Tasks.scheduleType
    var name by Tasks.name
    var command by Tasks.command
    var removed by Tasks.removed
    var createTime by Tasks.createTime
    var updateTime by Tasks.updateTime

    val shouldScheduledToday: Boolean
        get() {
            return false
        }

    fun createJob() {
        //todo
    }

}

