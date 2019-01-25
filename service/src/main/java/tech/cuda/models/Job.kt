package tech.cuda.models

import org.jetbrains.exposed.dao.EntityID
import org.jetbrains.exposed.dao.IntEntity
import org.jetbrains.exposed.dao.IntEntityClass
import org.jetbrains.exposed.dao.IntIdTable

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
            //todo
            return false
        }

    fun createInstance() {
        //todo

    }


}