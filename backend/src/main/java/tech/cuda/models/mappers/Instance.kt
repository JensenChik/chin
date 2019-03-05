package tech.cuda.models.mappers

import org.jetbrains.exposed.dao.EntityID
import org.jetbrains.exposed.dao.IntEntity
import org.jetbrains.exposed.dao.IntEntityClass
import org.jetbrains.exposed.dao.IntIdTable
import tech.cuda.enums.InstanceStatus
import tech.cuda.enums.SQL

/**
 * Created by Jensen on 18-6-15.
 */
object Instances : IntIdTable() {
    override val tableName: String
        get() = "instances"

    val job = reference(name = "job_id", foreign = Jobs)
    val output = blob("output")
    val status = customEnumeration(
            name = "status", sql = SQL<InstanceStatus>(),
            fromDb = { value -> InstanceStatus.valueOf(value as String) },
            toDb = { it.name }
    )
    val removed = bool(name = "removed").index().default(false)
    val createTime = datetime(name = "create_time")
    val updateTime = datetime(name = "update_time")
}


class Instance(id: EntityID<Int>) : IntEntity(id) {
    companion object : IntEntityClass<Instance>(Instances)

    var status by Instances.status
    var job by Job referencedOn Instances.job
    var output by Instances.output
    var removed by Instances.removed
    var createTime by Instances.createTime
    var updateTime by Instances.updateTime


    val finished: Boolean
        get() {
            return this.status == InstanceStatus.Success || this.status == InstanceStatus.Failed
        }

    val success: Boolean
        get() {
            return this.status == InstanceStatus.Success
        }

    fun start() {
    }

}