package tech.cuda.models

import org.jetbrains.exposed.dao.EntityID
import org.jetbrains.exposed.dao.IntEntity
import org.jetbrains.exposed.dao.IntEntityClass
import org.jetbrains.exposed.dao.IntIdTable
import javax.sql.rowset.serial.SerialBlob

/**
 * Created by Jensen on 18-6-15.
 */

enum class InstanceStatus {
    Waiting, Running, Success, Failed, Killing
}

fun String.toBlob(): SerialBlob = SerialBlob(this.toByteArray())

object InstanceTable : IntIdTable() {
    override val tableName: String
        get() = "instances"

    val job = reference(name = "job_id", foreign = JobTable)
    val machine = reference(name = "machine_id", foreign = MachineTable)
    val output = blob("output")
    val status = customEnumeration(
            name = "status",
            sql = InstanceStatus.values().joinToString(",", "ENUM(", ")") { "'${it.name}'" },
            fromDb = { value -> InstanceStatus.valueOf(value as String) },
            toDb = { it.name }
    )
    val removed = bool(name = "removed").index().default(false)
    val createTime = datetime(name = "create_time")
    val updateTime = datetime(name = "update_time")
}


class Instance(id: EntityID<Int>) : IntEntity(id) {
    companion object : IntEntityClass<Instance>(InstanceTable)

    var status by InstanceTable.status
    var job by Job referencedOn InstanceTable.job
    var machine by Machine referencedOn InstanceTable.machine
    var output by InstanceTable.output
    var removed by InstanceTable.removed
    var createTime by InstanceTable.createTime
    var updateTime by InstanceTable.updateTime


    val finished: Boolean
        get() {
            return this.status == InstanceStatus.Success || this.status == InstanceStatus.Failed
        }

    val success: Boolean
        get() {
            return this.status == InstanceStatus.Success
        }


}