package tech.cuda.models

import org.jetbrains.exposed.dao.EntityID
import org.jetbrains.exposed.dao.IntEntity
import org.jetbrains.exposed.dao.IntEntityClass
import org.jetbrains.exposed.dao.IntIdTable

/**
 * Created by Jensen on 18-6-18.
 */
object Jobs : IntIdTable() {
    val taskId = integer(name = "task_id").index()
    val removed = bool(name = "removed").index()
    val createTime = datetime(name = "create_time")
    val updateTime = datetime(name = "update_time")
}


class Job(id: EntityID<Int>) : IntEntity(id) {
    companion object : IntEntityClass<Job>(Jobs)

    val shouldRunNow: Boolean
        get() {
            //todo
            return false
        }

    fun createInstance() {
        //todo

    }


}