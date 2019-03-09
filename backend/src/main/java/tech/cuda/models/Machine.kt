package tech.cuda.models

import org.jetbrains.exposed.dao.EntityID
import org.jetbrains.exposed.dao.IntEntity
import org.jetbrains.exposed.dao.IntEntityClass
import org.jetbrains.exposed.dao.IntIdTable

/**
 * Created by Jensen on 18-6-18.
 */
object MachineTable : IntIdTable() {
    override val tableName: String
        get() = "machines"
    const val NAME_MAX_LEN = 256
    val name = varchar(name = "name", length = NAME_MAX_LEN)
    const val IP_MAX_LEN = 256
    val ip = varchar(name = "ip", length = IP_MAX_LEN)
    const val MAC_MAX_LEN = 256
    val mac = varchar(name = "mac", length = MAC_MAX_LEN)
    val cpuLoad = integer(name = "cpu_load")
    val memLoad = integer(name = "mem_load")
    val diskLoad = integer(name = "disk_load")
    val removed = bool(name = "removed").index().default(false)
    val createTime = datetime(name = "create_time")
    val updateTime = datetime(name = "update_time")
}


class Machine(id: EntityID<Int>) : IntEntity(id) {
    companion object : IntEntityClass<Machine>(MachineTable)

    var name by MachineTable.name
    var ip by MachineTable.ip
    var mac by MachineTable.mac
    var cpuLoad by MachineTable.cpuLoad
    var memLoad by MachineTable.memLoad
    var diskLoad by MachineTable.diskLoad
    var removed by MachineTable.removed
    var createTime by MachineTable.createTime
    var updateTime by MachineTable.updateTime

    val alive: Boolean
        get() {
            return false
        }
}
