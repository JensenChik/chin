package tech.cuda.models

import org.jetbrains.exposed.dao.EntityID
import org.jetbrains.exposed.dao.IntEntity
import org.jetbrains.exposed.dao.IntEntityClass
import org.jetbrains.exposed.dao.IntIdTable

/**
 * Created by Jensen on 18-6-18.
 */
object Tasks : IntIdTable() {
    val groupId = integer(name = "group_id")
    val userId = integer(name = "user_id")
    val name = varchar(name = "name", length = 256)
    val command = text(name = "command")

    val removed = bool(name = "removed").index()
    val createTime = datetime(name = "create_time")
    val updateTime = datetime(name = "update_time")


}


class Task(id: EntityID<Int>) : IntEntity(id) {
    companion object : IntEntityClass<Task>(Tasks)

    var name by Tasks.name

    val shouldScheduledToday: Boolean
        get() {
            return false
        }

    fun createJob() {
        //todo
    }

}

