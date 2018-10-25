package tech.cuda.models

import org.jetbrains.exposed.dao.EntityID
import org.jetbrains.exposed.dao.IntEntity
import org.jetbrains.exposed.dao.IntEntityClass
import org.jetbrains.exposed.dao.IntIdTable
import org.jetbrains.exposed.sql.Table

/**
 * Created by Jensen on 18-6-15.
 */
object Instances : IntIdTable() {
    val jobId = integer(name = "job_id").index()
    val output = blob("output")
    val removed = bool(name = "removed").index()
    val createTime = datetime(name = "create_time")
    val updateTime = datetime(name = "update_time")
}


class Instance(id: EntityID<Int>) : IntEntity(id) {
    companion object : IntEntityClass<Instance>(Instances)

    val finished: Boolean
        get() {
            return false
        }

    val success: Boolean
        get() {
            return false
        }

    fun start() {
    }

}