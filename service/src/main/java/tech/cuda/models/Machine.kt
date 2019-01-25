package tech.cuda.models

import org.jetbrains.exposed.dao.IntIdTable

/**
 * Created by Jensen on 18-6-18.
 */
object Machines : IntIdTable() {
    val name = varchar(name = "name", length = 256)

    val ip = varchar(name = "ip", length = 256)
    val mac = varchar(name = "mac", length = 256)

    val cpuLoad = integer(name = "cpu_load")
    val memLoad = integer(name = "mem_load")
    val diskLoad = integer(name = "disk_load")

    val removed = bool(name = "removed").index()
    val createTime = datetime(name = "create_time")
    val updateTime = datetime(name = "update_time")

    val alive: Boolean
        get() {
            return false
        }
}