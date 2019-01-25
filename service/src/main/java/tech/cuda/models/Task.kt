package tech.cuda.models

import org.jetbrains.exposed.dao.EntityID
import org.jetbrains.exposed.dao.IntEntity
import org.jetbrains.exposed.dao.IntEntityClass
import org.jetbrains.exposed.dao.IntIdTable


/**
 * Created by Jensen on 18-6-18.
 */
object Tasks : IntIdTable() {
    val group = reference(name = "group", foreign = Groups)
    val user = reference(name = "user", foreign = Users)

    val name = varchar(name = "name", length = 256)
    val command = text(name = "command")

    val removed = bool(name = "removed").index()
    val createTime = datetime(name = "create_time")
    val updateTime = datetime(name = "update_time")


}


class Task(id: EntityID<Int>) : IntEntity(id) {
    companion object : IntEntityClass<Task>(Tasks)

    var group by Group referencedOn Tasks.group
    var user by User referencedOn Tasks.user

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

