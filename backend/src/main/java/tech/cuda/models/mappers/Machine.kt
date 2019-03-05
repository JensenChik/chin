package tech.cuda.models.mappers

import org.jetbrains.exposed.dao.EntityID
import org.jetbrains.exposed.dao.IntEntity
import org.jetbrains.exposed.dao.IntEntityClass
import org.jetbrains.exposed.dao.IntIdTable

/**
 * Created by Jensen on 18-6-18.
 */
object Machines : IntIdTable() {
    override val tableName: String
        get() = "machines"

    val name = varchar(name = "name", length = 256)
    val ip = varchar(name = "ip", length = 256)
    val mac = varchar(name = "mac", length = 256)
    val cpuLoad = integer(name = "cpu_load")
    val memLoad = integer(name = "mem_load")
    val diskLoad = integer(name = "disk_load")
    val removed = bool(name = "removed").index().default(false)
    val createTime = datetime(name = "create_time")
    val updateTime = datetime(name = "update_time")
}


class Machine(id: EntityID<Int>) : IntEntity(id) {
    companion object : IntEntityClass<Machine>(Machines)

    var name by Machines.name
    val ip by Machines.ip
    val mac by Machines.mac
    val cpuLoad by Machines.cpuLoad
    val memLoad by Machines.memLoad
    val diskLoad by Machines.diskLoad
    val removed by Machines.removed
    val createTime by Machines.createTime
    val updateTime by Machines.updateTime

    val alive: Boolean
        get() {
            return false
        }
}
