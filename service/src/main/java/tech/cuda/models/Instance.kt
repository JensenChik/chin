package tech.cuda.models

import org.jetbrains.exposed.dao.EntityID
import org.jetbrains.exposed.dao.IntEntity
import org.jetbrains.exposed.dao.IntEntityClass
import org.jetbrains.exposed.dao.IntIdTable

/**
 * Created by Jensen on 18-6-15.
 */
object Instances : IntIdTable() {
    val job = reference(name = "job", foreign = Jobs)
    val output = blob("output")
    val removed = bool(name = "removed").index()
    val createTime = datetime(name = "create_time")
    val updateTime = datetime(name = "update_time")
}


class Instance(id: EntityID<Int>) : IntEntity(id) {
    companion object : IntEntityClass<Instance>(Instances)

    var job by Job referencedOn Instances.job
    var output by Instances.output
    var removed by Instances.removed
    var createTime by Instances.createTime
    var updateTime by Instances.updateTime


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